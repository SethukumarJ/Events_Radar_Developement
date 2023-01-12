package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return UserHandler{
		userUseCase: usecase,
	}
}

// @Summary update Profileabout
// @ID Update userprofile
// @Tags User
// @Produce json
// @Security BearerAuth
// @param UpdateProfile body domain.Bios{} true "update profile with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/update/profile [patch]
func (cr *UserHandler) UpdateProfile(c *gin.Context) {

	var updatedProfile domain.Bios
	fmt.Println("Updating event")
	//fetching data
	c.Bind(&updatedProfile)

	username := c.Writer.Header().Get("userName")
	fmt.Println("username ", username)

	//check event exit or not

	err := cr.userUseCase.UpdateProfile(updatedProfile, username)
	fmt.Println("error on updaed profile", err)

	log.Println(updatedProfile)

	if err != nil && err != sql.ErrNoRows {
		response := response.ErrorResponse("Failed to Update Profile", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	event, _ := cr.userUseCase.FindUser(updatedProfile.UserName)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary update password
// @ID Update password
// @Tags User
// @Produce json
// @Param  email   query  string  true  "Email: "
// @param Updateevent body domain.Users{} true "update password with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/password/update [patch]
func (cr *UserHandler) UpdatePassword(c *gin.Context) {

	var updatedPassword domain.Users
	fmt.Println("Updating event")
	//fetching data
	c.Bind(&updatedPassword)
	fmt.Println("userPassword", updatedPassword.UserName)
	email := c.Query("email")

	//check event exit or not

	err := cr.userUseCase.UpdatePassword(updatedPassword, email)

	fmt.Println(updatedPassword.Password)

	if err != nil {
		response := response.ErrorResponse("Failed to Update Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.userUseCase.FindUser(email)
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}


// SendVerificationEmail sends the verification email

// @Summary Send verification
// @ID Send verifiation code via email
// @Tags User
// @Produce json
// @Param  email   query  string  true  "Email: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/send/verification [post]
func (cr *UserHandler) SendVerificationMail(c *gin.Context) {

	email := c.Query("email")
	var code int
	fmt.Println(code)
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



// @Summary list all Public faqas
// @ID list all public faqas
// @Tags User
// @Produce json
// @Param  title   query  string  true  "Event title: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router user/event/faqas/get [get]
func (cr *EventHandler) GetPublicFaqas(c *gin.Context) {

	title := c.Query("title")
	faqas,err := cr.userUserCase.GetPublicFaqas(title)


	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Events", result)
	utils.ResponseJSON(*c, response)

}

// @Summary Post Question function
// @ID User Post Question
// @Tags User
// @Produce json
// @Security BearerAuth
// @param title query string true "Getting the title of the event"
// @param organizername query string true "Getting the title of the event"
// @param PostQuestion body domain.Faqas{} true "Post question"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/event/post/question [post]
// PostQuesition handles Posting events
func (cr *UserHandler) PostQuestion(c *gin.Context) {

	var question domain.Faqas
	title := c.Query("title")
	organizerName := c.Query("organizername")
	username := c.Writer.Header().Get("userName")
	c.Bind(&question)

	question.Title = title
	question.UserName = username
	question.OrganizerName = organizerName

	err := cr.userUseCase.PostQuestion(question)

	log.Println(question)

	if err != nil {
		response := response.ErrorResponse("Failed to Post question", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	
	response := response.SuccessResponse(true, "SUCCESS", question)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)



}
