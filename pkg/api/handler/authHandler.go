package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type AuthHandler struct {
	jwtUserUsecase usecase.JWTUsecase
	authUsecase    usecase.AuthUsecase
	userUsecase    usecase.UserUseCase
}

func NewAuthHandler(
	jwtUserUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	authUsecase usecase.AuthUsecase,

) *AuthHandler {
	return &AuthHandler{
		jwtUserUsecase: jwtUserUsecase,
		authUsecase:    authUsecase,
		userUsecase:    userUsecase,
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
