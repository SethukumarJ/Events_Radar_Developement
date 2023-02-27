package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	jwtUsecase   usecase.JWTUsecase
	authUsecase  usecase.AuthUsecase
	adminUsecase usecase.AdminUsecase
	userUsecase  usecase.UserUseCase
	cfg          config.Config
}

func NewAuthHandler(
	jwtUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	adminUsecase usecase.AdminUsecase,
	authUsecase usecase.AuthUsecase,
	cfg config.Config,

) AuthHandler {
	return AuthHandler{
		jwtUsecase:   jwtUsecase,
		authUsecase:  authUsecase,
		userUsecase:  userUsecase,
		adminUsecase: adminUsecase,
		cfg:          cfg,
	}
}

func (cr *AuthHandler) VerifyAccount(c *gin.Context) {
	tokenString := c.Query("token")
	fmt.Println("varify account from authhandler called , ", tokenString)
	var email string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid verification token")
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// get the username from the claims
		email = claims["username"].(string)

	} else {
		c.String(http.StatusBadRequest, "Invalid verification token")
		return
	}

	err = cr.authUsecase.VerifyAccount(email, tokenString)

	if err != nil {
		response := response.ErrorResponse("Verification failed, Jwt is not valid", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Account verified successfully", email)
	utils.ResponseJSON(*c, response)

}

// @Summary SignUp for users
// @ID User SignUp
// @Tags User
// @Produce json
// @param RegisterUser body domain.Users{} true "user signup with username, phonenumber email ,password"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/signup [post]
// UserSignup handles the user signup
func (cr *AuthHandler) UserSignup(c *gin.Context) {

	var newUser domain.Users
	fmt.Println("user signup")
	//fetching data
	c.Bind(&newUser)

	//check username exit or not

	err := cr.userUsecase.CreateUser(newUser)

	log.Println(newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.userUsecase.FindUser(newUser.Email)
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:3000/user/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

func (cr *AuthHandler) InitializeOAuthGoogle() {
	oauthConfGl.ClientID = cr.cfg.CLIENT_ID
	oauthConfGl.ClientSecret = cr.cfg.CLIENT_SECRET
	oauthStateStringGl = cr.cfg.OauthStateString
	fmt.Printf("\n\n%v\n\n", oauthConfGl)
}

// @Summary Authenticate With Google
// @ID Authenticate With Google
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/refresh-tocken [get]
func (cr *AuthHandler) GoogleAuth(c *gin.Context) {
	HandileLogin(c, oauthConfGl, oauthStateStringGl)
}

func HandileLogin(c *gin.Context, oauthConf *oauth2.Config, oauthStateString string) error {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		fmt.Printf("\n\n\nerror in handile login :%v\n\n", err)
		return err
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	fmt.Printf("\n\nurl : %v\n\n", oauthConf.RedirectURL)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil

}
func (cr *AuthHandler) CallBackFromGoogle(c *gin.Context) {
	fmt.Print("\n\nfuck\n\n")
	c.Request.ParseForm()
	state := c.Request.FormValue("state")

	if state != oauthStateStringGl {
		fmt.Println("//////////////////////hellooooo1////////////////////////////////////")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Println("//////////////////////hellooooo2////////////////////////////////////")
	code := c.Request.FormValue("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, "Code Not Found to provide AccessToken..\n")

		reason := c.Request.FormValue("error_reason")
		if reason == "user_denied" {
			fmt.Println("//////////////////////hai////////////////////////////////////")
			c.JSON(http.StatusBadRequest, "User has denied Permission..")
		}
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			return
		}
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		fmt.Println("//////////////////////hai2No err////////////////////////////////////")
		if err != nil {
			fmt.Println("//////////////////////hai2////////////////////////////////////")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		defer resp.Body.Close()

		response1, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("//////////////////////hai3////////////////////////////////////")
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		fmt.Println("//////////////////////hai4////////////////////////////////////")
		type data struct {
			Id             int
			Email          string
			Verified_email bool
			Picture        string
			// data           string
		}
		var any data
		json.Unmarshal(response1, &any)
		
		fmt.Printf("\n\ndata :%v\n\n", string(response1))
		fmt.Printf("\n\ndata :%v\n\n", any)
		fmt.Println("email",any.Email)

		newUser := domain.Users{}
		newUser.UserName,newUser.Email,newUser.Profile = any.Email, any.Email,any.Picture
	

		user, err := cr.userUsecase.FindUser(any.Email)
		if err != nil {
			fmt.Println(err)
		}
		if user == nil {
			err = cr.userUsecase.CreateUser(newUser)
			log.Println(newUser)
			

			if err != nil {
				response := response.ErrorResponse("Failed to create user", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnprocessableEntity)
				utils.ResponseJSON(*c, response)
				return
			}
			newUser, err := cr.userUsecase.FindUser(any.Email)
			if err != nil {
				fmt.Println(err)
			}
			accesstoken, err := cr.jwtUsecase.GenerateAccessToken(newUser.UserId, newUser.UserName, "user")
			if err != nil {
				response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			newUser.AccessToken = accesstoken
			refreshtoken, err := cr.jwtUsecase.GenerateRefreshToken(newUser.UserId, newUser.UserName, "user")
			if err != nil {
				response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			newUser.RefreshToken = refreshtoken

			Tokens := map[string]string{"AccessToken": newUser.AccessToken, "RefreshToken": newUser.RefreshToken}
			response := response.SuccessResponse(true, "SUCCESSfully created the user and signed in", Tokens)
			utils.ResponseJSON(*c, response)

			fmt.Println("google login function returned successfully")
		} else {
			
			accesstoken, err := cr.jwtUsecase.GenerateAccessToken(user.UserId, user.UserName, "user")
			if err != nil {
				response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			user.AccessToken = accesstoken
			refreshtoken, err := cr.jwtUsecase.GenerateRefreshToken(user.UserId, user.UserName, "user")
			if err != nil {
				response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				return
			}
			user.RefreshToken = refreshtoken

			Tokens := map[string]string{"AccessToken": user.AccessToken, "RefreshToken": user.RefreshToken}
			response := response.SuccessResponse(true, "SUCCESS, logged in", Tokens)
			utils.ResponseJSON(*c, response)

		}

		c.JSON(http.StatusOK, "Hello, I'm protected\n")
		c.JSON(http.StatusOK, string(response1))
		return
	}
}





// UserLogin handles the user login

// @Summary Login for users
// @ID User Login
// @Tags User
// @Produce json
// @Tags User
// @Param  UserLogin   body  domain.Login{}  true  "userlogin: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
// UserLogin handles the user login
func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var userLogin domain.Login

	c.Bind(&userLogin)
	fmt.Println("userLOgin", userLogin.Password)
	//verify User details
	err := cr.authUsecase.VerifyUser(userLogin.Email, userLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	user, _ := cr.userUsecase.FindUser(userLogin.Email)
	accesstoken, err := cr.jwtUsecase.GenerateAccessToken(user.UserId, user.UserName, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.AccessToken = accesstoken
	refreshtoken, err := cr.jwtUsecase.GenerateRefreshToken(user.UserId, user.UserName, "user")
	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	user.RefreshToken = refreshtoken

	Tokens := map[string]string{"AccessToken": user.AccessToken, "RefreshToken": user.RefreshToken}
	response := response.SuccessResponse(true, "SUCCESS", Tokens)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// user refresh token

// @Summary Refresh token for users
// @ID User RefreshToken
// @Tags User
// @Produce json
// @Tags User
// @Security BearerAuth
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/token-refresh [post]
// UserLogin handles the user login
func (cr *AuthHandler) UserRefreshToken(c *gin.Context) {

	autheader := c.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	fmt.Printf("\n\ntocen : %v\n\n", autheader)
	token := bearerToken[1]
	ok, claims := cr.jwtUsecase.VerifyToken(token)
	if !ok {
		log.Fatal("referesh token not valid")
	}

	fmt.Println("//////////////////////////////////", claims.UserName)
	accesstoken, err := cr.jwtUsecase.GenerateAccessToken(claims.UserId, claims.UserName, claims.Role)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", accesstoken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// UserSignup handles the admin signup
// @Summary SignUp for Admin
// @ID SignUp authentication
// @Tags Admin
// @Produce json
// @Tags Admin
// @param RegisterAdmin body domain.Admins{} true "admin signup with username, phonenumber email ,password"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/signup [post]
// AdminSignup handles the user signup
func (cr *AuthHandler) AdminSignup(c *gin.Context) {

	var newAdmin domain.Admins
	fmt.Println("admin signup")
	//fetching data
	c.Bind(&newAdmin)

	//check username exit or not

	err := cr.adminUsecase.CreateAdmin(newAdmin)

	log.Println(newAdmin)

	if err != nil {
		response := response.ErrorResponse("Failed to create admin", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	admin, _ := cr.adminUsecase.FindAdmin(newAdmin.Email)
	response := response.SuccessResponse(true, "SUCCESS", admin)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// UserLogin handles the user login

// @Summary Login for Admin
// @ID Admin Login
// @Tags Admin
// @Produce json
// @Tags Admin
// @Param  AdminLogin   body  domain.Login{}  true  "adminlogin: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/login [post]
// AdminLogin handles the user login
func (cr *AuthHandler) AdminLogin(c *gin.Context) {

	var adminLogin domain.Admins

	c.Bind(&adminLogin)
	fmt.Println("adminLoginpassword", adminLogin.Password)
	//verify User details
	err := cr.authUsecase.VerifyAdmin(adminLogin.Email, adminLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	admin, _ := cr.adminUsecase.FindAdmin(adminLogin.Email)
	accesstoken, err := cr.jwtUsecase.GenerateAccessToken(admin.AdminId, admin.AdminName, "admin")
	if err != nil {
		response := response.ErrorResponse("Failed to generate access token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	admin.AccessToken = accesstoken
	refreshtoken, err := cr.jwtUsecase.GenerateRefreshToken(admin.AdminId, admin.AdminName, "admin")
	if err != nil {
		response := response.ErrorResponse("Failed to generate refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}
	admin.RefreshToken = refreshtoken

	Tokens := map[string]string{"AccessToken": admin.AccessToken, "RefreshToken": admin.RefreshToken}
	response := response.SuccessResponse(true, "SUCCESS", Tokens)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// user refresh token
// @Summary Refresh token for admin
// @ID Admin RefreshToken
// @Tags Admin
// @Produce json
// @Tags Admin
// @Security BearerAuth
// @Param  Authorization   header  string  true  "token string: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/token-refresh [post]
func (cr *AuthHandler) AdminRefreshToken(c *gin.Context) {

	autheader := c.Request.Header["Authorization"]
	auth := strings.Join(autheader, " ")
	bearerToken := strings.Split(auth, " ")
	fmt.Printf("\n\ntocen : %v\n\n", autheader)
	token := bearerToken[1]
	ok, claims := cr.jwtUsecase.VerifyToken(token)
	if !ok {
		log.Fatal("referesh token not valid")
	}

	fmt.Println("//////////////////////////////////", claims.UserName)
	accesstoken, err := cr.jwtUsecase.GenerateAccessToken(claims.UserId, claims.UserName, claims.Role)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", accesstoken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}
