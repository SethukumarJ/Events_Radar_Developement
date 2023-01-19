package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

// @Summary Add Admins
// @ID Add admins for the organizaition
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param addMembers body []string true "addMembers:"
// @Param pathrole query string true "Your role:"
// @Param role query striing true "member role"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /Organization/add-memebers [Post]
func (cr *UserHandler) AddMembers(c *gin.Context) {


	var newMembers []string
	username := c.Writer.Header().Get("userName")
	fmt.Println("username ", username)
	organizationName := c.Writer.Header().Get("organizationName")
	fmt.Println("organizationName ", organizationName)
	role := c.Writer.Header().Get("role")
	memberRole := c.Query("role")
	fmt.Println("role ", role)
	c.Bind(&newMembers)

	if role > "4" {
		response := response.ErrorResponse("Your role is not eligible for this action","no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err := cr.userUseCase.AddMembers(username, memberRole)
	if err != nil {
		response := response.ErrorResponse("error while adding memebers to the database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Showing the newly added members", newMembers)
	utils.ResponseJSON(*c, response)



}

// @Summary Get Organization
// @ID Get Organizaition by name
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  organizationName   query  string  true  "OrganizationName: "
// @Param  pathRole query string true "role:"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /get-organization [get]
func (cr *UserHandler) GetOrganization(c *gin.Context) {
	username := c.Writer.Header().Get("userName")
	fmt.Println("username ", username)
	organizationName := c.Writer.Header().Get("organizationName")
	fmt.Println("organizationName ", organizationName)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)

	if role > "4" {
		response := response.ErrorResponse("Your role is not eligible for this action","no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	organization, err := cr.userUseCase.FindOrganization(organizationName)

	fmt.Println("organization:", organization)

	if err != nil {
		response := response.ErrorResponse("error while getting event from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Showing the event", organization)
	utils.ResponseJSON(*c, response)

}

// @Summary Joining organization
// @ID Join organization
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param  organizationName   query  string  true  "organization name: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/organization/join [patch]
func (cr *UserHandler) JoinOrganization(c *gin.Context) {

	username := c.Writer.Header().Get("userName")
	fmt.Println("username ", username)

	organizationName := (c.Query("organizationName"))

	err := cr.userUseCase.JoinOrganization(organizationName, username)

	if err != nil {
		response := response.ErrorResponse("Joining organization failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Requested to join organization ", organizationName)
	utils.ResponseJSON(*c, response)

}

// @Summary list all registered organizations for user
// @ID list all registered organizations
// @Tags User
// @Produce json
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-organizations [get]
func (cr *UserHandler) ListOrganizations(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	log.Println(page, "   ", pageSize)

	fmt.Println("page :", page)
	fmt.Println("pagesize", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	fmt.Println("pagenation", pagenation)

	organizations, metadata, err := cr.userUseCase.ListOrganizations(pagenation)

	fmt.Println("events:", organizations)

	result := struct {
		Organizations *[]domain.OrganizationsResponse
		Meta          *utils.Metadata
	}{
		Organizations: organizations,
		Meta:          metadata,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting organization applications from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Organization applications", result)
	utils.ResponseJSON(*c, response)

}

// @Summary Create Organization
// @ID Create Organizatioin from user
// @Tags User
// @Produce json
// @Security BearerAuth
// @param CreateOrganization body domain.Organizations{} true "Create organization"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/organization/create [post]
// Create Organization
func (cr *UserHandler) CreateOrganization(c *gin.Context) {

	var newOrganization domain.Organizations

	fmt.Println("Creating Organizations")
	//fetching data
	c.Bind(&newOrganization)

	fmt.Println("event", newOrganization)
	newOrganization.CreatedBy = c.Writer.Header().Get("userName")
	newOrganization.CreatedAt = time.Now()

	err := cr.userUseCase.CreateOrganization(newOrganization)

	log.Println(newOrganization)

	if err != nil {
		response := response.ErrorResponse("Failed to create Organization", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	organization, _ := cr.userUseCase.FindOrganization(newOrganization.OrganizationName)
	response := response.SuccessResponse(true, "SUCCESS", organization)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
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
// @Router /user/list/faqas [get]
func (cr *UserHandler) GetPublicFaqas(c *gin.Context) {

	title := c.Query("title")
	faqas, err := cr.userUseCase.GetPublicFaqas(title)
	fmt.Println("faqas from handler", faqas)
	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All faqas", faqas)
	fmt.Println("response", response)
	utils.ResponseJSON(*c, response)

}

// @Summary list all Asked questions
// @ID list all asked questions
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param  title   query  string  true  "Event title: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list/questions [get]
func (cr *UserHandler) GetQuestions(c *gin.Context) {

	title := c.Query("title")
	questions, err := cr.userUseCase.GetQuestions(title)
	fmt.Println("Questions from handler", questions)
	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All faqas", questions)
	fmt.Println("response", response)
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

// @Summary Post Answer function
// @ID User Post Answer
// @Tags User
// @Produce json
// @Security BearerAuth
// @param faqaid query string true "Getting the id of the question"
// @param PostAnswer body domain.Answers{} true "Post Answer"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/event/post/answer [post]
// PostQuesition handles Posting events
func (cr *UserHandler) PostAnswer(c *gin.Context) {

	var answer domain.Answers
	question_id, _ := strconv.Atoi(c.Query("faqaid"))

	c.Bind(&answer)

	err := cr.userUseCase.PostAnswer(answer, question_id)

	log.Println(question_id)

	if err != nil {
		response := response.ErrorResponse("Failed to Post question", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", answer)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}
