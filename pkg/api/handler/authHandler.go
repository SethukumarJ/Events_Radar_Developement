package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type AuthHandler struct {
	jwtUsecase   usecase.JWTUsecase
	authUsecase  usecase.AuthUsecase
	adminUsecase usecase.AdminUsecase
	userUsecase  usecase.UserUseCase
}

func NewAuthHandler(
	jwtUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	adminUsecase usecase.AdminUsecase,
	authUsecase usecase.AuthUsecase,

) AuthHandler {
	return AuthHandler{
		jwtUsecase:   jwtUsecase,
		authUsecase:  authUsecase,
		userUsecase:  userUsecase,
		adminUsecase: adminUsecase,
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
