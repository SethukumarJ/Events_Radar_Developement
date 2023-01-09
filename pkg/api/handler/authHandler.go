package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type AuthHandler struct {
	jwtUsecase usecase.JWTUsecase
	authUsecase    usecase.AuthUsecase
	adminUsecase usecase.AdminUsecase
	userUsecase    usecase.UserUseCase
}

func NewAuthHandler(
	jwtUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	adminUsecase usecase.AdminUsecase,
	authUsecase usecase.AuthUsecase,

) AuthHandler {
	return AuthHandler{
		jwtUsecase: jwtUsecase,
		authUsecase:    authUsecase,
		userUsecase:    userUsecase,
		adminUsecase:   adminUsecase,
	}
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
	fmt.Println("userid",newUser.UserId)

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
// @Param  UserLogin   body  domain.Users{}  true  "userlogin: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/login [post]
// UserLogin handles the user login
func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var userLogin domain.Users

	c.Bind(&userLogin)
	fmt.Println("userLOgin",userLogin.Password)
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
	token := cr.jwtUsecase.GenerateToken(user.UserId, user.UserName, "user")
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user.Token)
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
// @Param  Authorization   header  string  true  "token string: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/token/refresh [post]
// UserLogin handles the user login
func (cr *AuthHandler) UserRefreshToken(c *gin.Context) {

	autheader := c.Request.Header["Authorization"]
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntocen : %v\n\n", autheader)
		token := bearerToken[1]

	refreshToken, err := cr.jwtUsecase.GenerateRefreshToken(token)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", refreshToken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

//Adminsiginup handles the user signup



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
// @Param  email   path  string  true  "admin email: "
// @Param  password   path  string  true  "admin password: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/login [post]
// AdminLogin handles the user login
func (cr *AuthHandler) AdminLogin(c *gin.Context) {

	var adminLogin domain.Admins

	c.Bind(&adminLogin)
	fmt.Println("adminLoginpassword",adminLogin.Password)
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
	token := cr.jwtUsecase.GenerateToken(admin.AdminId, admin.Email, "admin")
	admin.Token = token
	response := response.SuccessResponse(true, "SUCCESS", admin.Token)
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
// @Router /admin/token/refresh [post]
func (cr *AuthHandler) AdminRefreshToken(c *gin.Context) {

	autheader := ("Authorization")
	bearerToken := strings.Split(autheader, " ")
	token := bearerToken[1]

	refreshToken, err := cr.jwtUsecase.GenerateRefreshToken(token)

	if err != nil {
		response := response.ErrorResponse("error generating refresh token", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", refreshToken)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}
