package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// SendVerificationEmail sends the verification email

func (cr *UserHandler) SendVerificationMail(c *gin.Context) {

	email := c.Query("Email")

	_, err := cr.userUseCase.FindUser(email)
	fmt.Println("email: ", email)
	fmt.Println("err: ", err)

	if err == nil {
		err = cr.userUseCase.SendVerificationEmail(email)
	}

	fmt.Println(err)

	if err != nil {
		response := response.ErrorResponse("Error while sending verification mail", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Verification mail sent successfully", email)
	utils.ResponseJSON(*c, response)

}

// verifyAccount verifies the account

func (cr *UserHandler) VerifyAccount(c *gin.Context) {

	email := c.Query("Email")
	code, _ := strconv.Atoi(c.Query("Code"))

	err := cr.userUseCase.VerifyAccount(email, code)

	if err != nil {
		response := response.ErrorResponse("Verification failed, Invalid OTP", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Account verified successfully", email)
	utils.ResponseJSON(*c, response)

}
