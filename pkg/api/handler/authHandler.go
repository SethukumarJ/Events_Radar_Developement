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
	jwtUserUsecase usecase.JWTUsecase
	jwtAdminUsecase usecase.JWTUsecase
	authUsecase    usecase.AuthUsecase
	adminUsecase usecase.AdminUsecase
	userUsecase    usecase.UserUseCase
}

func NewAuthHandler(
	jwtUserUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	adminUsecase usecase.AdminUsecase,
	authUsecase usecase.AuthUsecase,

) AuthHandler {
	return AuthHandler{
		jwtUserUsecase: jwtUserUsecase,
		authUsecase:    authUsecase,
		userUsecase:    userUsecase,
		adminUsecase:   adminUsecase,
	}
}

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

func (cr *AuthHandler) UserLogin(c *gin.Context) {

	var userLogin domain.Users

	c.Bind(&userLogin)

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
	token := cr.jwtUserUsecase.GenerateToken(user.UserId, user.Email, "user")
	user.Token = token
	response := response.SuccessResponse(true, "SUCCESS", user.Token)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}

// user refresh token
func (cr *AuthHandler) UserRefreshToken(c *gin.Context) {

	autheader := ("Authorization")
	bearerToken := strings.Split(autheader, " ")
	token := bearerToken[1]

	refreshToken, err := cr.jwtUserUsecase.GenerateRefreshToken(token)

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

func (cr *AuthHandler) AdminSignup(c *gin.Context) {

	var newAdmin domain.Admins
	fmt.Println("user signup")
	//fetching data
	c.Bind(&newAdmin)

	//check username exit or not

	err := cr.adminUsecase.CreateAdmin(newAdmin)

	log.Println(newAdmin)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.adminUsecase.FindAdmin(newAdmin.Email)
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}


// AdminLogin handles the admin login
func (cr *AuthHandler) AdminLogin(c *gin.Context) {

	var adminLogin domain.Admins

	c.Bind(&adminLogin)

	fmt.Println("adminLogin.passwrodk", adminLogin.Password)
	fmt.Println("adminLogin.username", adminLogin.AdminName)
	//verify User details
	err := cr.authUsecase.VerifyAdmin(adminLogin.AdminName, adminLogin.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to login", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnauthorized)
		utils.ResponseJSON(*c, response)
		return
	}

	//fetching user details
	admin, _ := cr.adminUsecase.FindAdmin(adminLogin.AdminName)
	token := cr.jwtAdminUsecase.GenerateToken(admin.AdminId, admin.AdminName, "admin")
	admin.Password = ""
	admin.Token = token
	response := response.SuccessResponse(true, "SUCCESS", admin.Token)
	utils.ResponseJSON(*c, response)

	fmt.Println("login function returned successfully")

}