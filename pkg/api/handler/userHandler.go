package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	razorpay "github.com/razorpay/razorpay-go"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) UserHandler {
	return UserHandler{
		userUseCase: usecase,
	}
}

// Initialize a map with key/value pairs
var packages = map[string]int{"basic": 10000, "stadard": 25000, "premium": 50000}

// @Summary Promote
// @ID promote event
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param Event_id query int true "Event_id :"
// @Param Organization_id query int true "Organization_id :"
// @param plan query string true "plan"
// @param email query string true "email"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/promote [Get]
func (cr *UserHandler) Pay(c *gin.Context) {

	// email := (c.Query("email"))
	// plan := (c.Query("plan"))
	// event_id, _ := strconv.Atoi(c.Query("Event_id"))

	// Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	// fmt.Println("Organization Id ", Organization_id)
	// role := c.Writer.Header().Get("role")
	// fmt.Println("role ", role)

	email := "sethukumarj.76@gmail.com"
	plan := "basic"
	event_id := 1

	Organization_id  := 1
	fmt.Println("Organization Id ", Organization_id)
	role := "1"
	fmt.Println("role ", role)


	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	promotion := domain.Promotion{}
	promotion.PromotedBy = email
	promotion.OrganizationId = Organization_id
	promotion.EventId = event_id
	promotion.Amount = fmt.Sprint(packages[plan])
	promotion.Plan = plan
	page := &domain.PageVariables{}
	page.Amount = promotion.Amount
	page.Email = email
	page.Name = ""
	page.Contact = ""
	//Create order_id from the server
	client := razorpay.NewClient("rzp_test_kEtg65WKqGTpKd", "gPURxG4gzTmeNJKqqz61YCHV")

	data := map[string]interface{}{
		"amount":   promotion.Amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	fmt.Println("////////////////reciept", body)
	if err != nil {
		fmt.Println("Problem getting the repository information", err)
		os.Exit(1)
	}

	value := body["id"]

	str := value.(string)
	promotion.OrderId = str
	fmt.Println("str////////////////", str)
	HomePageVars := domain.PageVariables{ //store the order_id in a struct
		OrderId: str,
		Amount:  page.Amount,
		Email:   page.Email,
		Name:    page.Name,
		Contact: page.Contact,
	}

	err = cr.userUseCase.PromoteEvent(promotion)
	if err != nil {
		response := response.ErrorResponse("Failed promote event event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}
	err = cr.userUseCase.FeaturizeEvent(str)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		response := response.ErrorResponse("Failed featurizing event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	c.HTML(http.StatusOK, "index.html", HomePageVars)

}

func (cr *UserHandler) Template(c *gin.Context) {

	// role := c.Writer.Header().Get("role")
	// fmt.Println("role ", role)


	c.HTML(http.StatusOK, "pay.html", nil)

}

func (cr *UserHandler) PaymentSuccess(c *gin.Context) {

	paymentid := c.Query("paymentid")
	orderid := c.Query("orderid")
	signature := c.Query("signature")
	err := cr.userUseCase.Prmotion_Success(orderid, paymentid)
	if err != nil {
		fmt.Println(err)
	}

	response := response.SuccessResponse(true, "SUCCESSFULLLY promoted event", orderid)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
	fmt.Println(paymentid, "paymentid")
	fmt.Println(orderid, "orderid")
	fmt.Println(signature, "signature")

}

func (cr *UserHandler) PaymentFaliure(c *gin.Context) {

	orderid := c.Query("orderid")
	errmsg := c.Query("errmsg")
	paymentid := c.Query("paymentid")
	fmt.Println(orderid, "orderid")
	fmt.Println(errmsg, "errmsg")
	res := []string{orderid, errmsg, paymentid}
	cr.userUseCase.Prmotion_Faliure(orderid, paymentid)
	response := response.ErrorResponse("Failed to make payments and cannot promote event", "", res)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary ApplyEvent
// @ID Apply event
// @Tags User-Event Management
// @Produce json
// @Security BearerAuth
// @Param Event_id query int true "Event_id"
// @param ApplyEvent body domain.ApplicationForm{} true "Apply event"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/apply-event [post]
// apply event
func (cr *UserHandler) ApplyEvent(c *gin.Context) {

	var newApplication domain.ApplicationForm

	fmt.Println("Applying event")
	//fetching data
	c.Bind(&newApplication)

	fmt.Println("organization", newApplication)
	newApplication.UserId, _ = strconv.Atoi(c.Writer.Header().Get("user_id"))
	newApplication.EventId, _ = strconv.Atoi(c.Query("Event_id"))
	newApplication.AppliedAt = time.Now()

	err := cr.userUseCase.ApplyEvent(newApplication)
	fmt.Println(err)
	log.Println(newApplication)

	if err != nil {
		response := response.ErrorResponse("Failed to apply event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	application, _ := cr.userUseCase.FindApplication(newApplication.UserId, newApplication.EventId)
	response := response.SuccessResponse(true, "SUCCESS", application)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Admit member
// @ID Admit member
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param  joinstatusid   query  int  true  "JoinStatusId: "
// @Param Organization_id query int true "Organization_id :"
// @Param role query string true "member role"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/admin/admit-member [patch]
func (cr *UserHandler) AdmitMember(c *gin.Context) {

	JoinStatusId, _ := strconv.Atoi(c.Query("joinstatusid"))
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)
	memberRole := c.Query("role")

	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	err := cr.userUseCase.AdmitMember(JoinStatusId, memberRole)

	if err != nil {
		response := response.ErrorResponse("admit member failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Member admitted", JoinStatusId)
	utils.ResponseJSON(*c, response)

}

// @Summary List Join Requests
// @ID Join requests to organization
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id :"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/join-requests [get]
func (cr *UserHandler) ListJoinRequests(c *gin.Context) {

	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)

	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	requests, err := cr.userUseCase.ListJoinRequests(user_id, Organization_id)
	if err != nil {
		response := response.ErrorResponse("error while getting requests applications from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed Join requests", requests)
	utils.ResponseJSON(*c, response)
}

// @Summary List Members
// @ID List Members of the organization
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id :"
// @Param memberRole query string true "Member role :"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/admin/list-members [get]
func (cr *UserHandler) ListMembers(c *gin.Context) {

	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)
	memberRole := c.Query("memberRole")
	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	members, err := cr.userUseCase.ListMembers(memberRole, Organization_id)
	if err != nil {
		response := response.ErrorResponse("error while getting members from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed Members", members)
	utils.ResponseJSON(*c, response)

}

// @Summary remove member
// @ID remvoe member
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id : "
// @Param user_id query int true "user_id to remove :"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/admin/delete-member [delete]
func (cr *UserHandler) RemoveMember(c *gin.Context) {

	role := c.Writer.Header().Get("role")

	user_id, _ := strconv.Atoi(c.Query("user_id"))

	Organization_id, _ := strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	fmt.Println(user_id, "////////////", Organization_id, ".................", role, "///////////////////////")
	fmt.Println("Organization Id ", Organization_id)
	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err := cr.userUseCase.DeleteMember(user_id, Organization_id)

	if err != nil {
		response := response.ErrorResponse("Could not remove member", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "removed member successfully!", user_id)
	utils.ResponseJSON(*c, response)

}

// @Summary Update member role
// @ID update  member role
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id : "
// @Param user_id query int true "user_id to update :"
// @Param updatedRole query string true "Role to update :"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/admin/update-role [patch]
func (cr *UserHandler) UpdateRole(c *gin.Context) {

	role := c.Writer.Header().Get("role")
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	updatedRole := c.Query("updatedRole")
	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err := cr.userUseCase.UpdateRole(user_id, Organization_id, updatedRole)

	if err != nil {
		response := response.ErrorResponse("Could not updae role of the member member", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Role updated successfully successfully!", user_id)
	utils.ResponseJSON(*c, response)

}

// @Summary Accept invitation to join an organization
// @ID Accept invitation to join organization
// @Tags User-Organization Management
// @Produce json
// @Param  token   query  string  true  "token: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/accept-invitation [get]
func (cr *UserHandler) AcceptJoinInvitation(c *gin.Context) {

	tokenString := c.Query("token")
	fmt.Println("varify account from authhandler called , ", tokenString)
	var email, role string
	var organization_id int
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("token , ", token)
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Println("error/////", err)
		c.String(http.StatusBadRequest, "Invalid invitation token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// get the username from the claims
		email = claims["username"].(string)
		organizationName := claims["organizationName"].(string)
		fmt.Println("//1//////////////////////////////////////", organizationName)
		organization_id, _ = strconv.Atoi((claims["organizationId"].(string)))

		fmt.Println("//2//////////////////////////////////////")
		role = claims["memberRole"].(string)

	} else {
		c.String(http.StatusBadRequest, "Invalid verification token")

		return
	}

	user, err := cr.userUseCase.FindUserByName(email)
	if err != nil {
		response := response.ErrorResponse("Joining failed, User is not signed up. signup to join the organization!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err = cr.userUseCase.AcceptJoinInvitation(int(user.UserId), organization_id, role)

	if err != nil {
		response := response.ErrorResponse("Verification failed, Jwt is not valid", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Joined organization successfully", email)
	utils.ResponseJSON(*c, response)

}

// @Summary Add Admins
// @ID Add admins for the organizaition
// @Tags Organizaton-Admin Role
// @Produce json
// @Security BearerAuth
// @Param addMembers body []domain.AddMembers{} true "addMembers:"
// @Param  Organization_id   query  string  true  "Organization_id: "
// @Param memberrole query string true "member role"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/admin/add-members [post]
func (cr *UserHandler) AddMembers(c *gin.Context) {

	var newMembers = []domain.AddMembers{}
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)
	memberRole := c.Query("memberrole")
	fmt.Println("role ", role)

	c.Bind(&newMembers)
	fmt.Println("newMembers", newMembers)
	if role > "1" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err := cr.userUseCase.AddMembers(newMembers, memberRole, Organization_id)
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
// @Tags User-Organization Management
// @Produce json
// @Param  Organization_id   query  string  true  "Organization_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/get-organization [get]
func (cr *UserHandler) GetOrganization(c *gin.Context) {

	Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	organization, err := cr.userUseCase.FindOrganizationById(Organization_id)

	fmt.Println("organization:", organization)

	if err != nil {
		response := response.ErrorResponse("error while getting organization from database", err.Error(), nil)
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
// @Tags User-Organization Management
// @Produce json
// @Security BearerAuth
// @Param  Organization_id   query  string  true  "Organization_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/join-organization [patch]
func (cr *UserHandler) JoinOrganization(c *gin.Context) {

	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	fmt.Println("Organization Id ", Organization_id)
	err := cr.userUseCase.JoinOrganization(Organization_id, user_id)

	if err != nil {
		response := response.ErrorResponse("Joining organization failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Requested to join organization ", Organization_id)
	utils.ResponseJSON(*c, response)

}

// @Summary list all registered organizations for user
// @ID list all registered organizations
// @Tags User-Organization Management
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

	response := response.SuccessResponse(true, "Listed All Organization in applications", result)
	utils.ResponseJSON(*c, response)

}

// @Summary Create Organization
// @ID Create Organizatioin from user
// @Tags User-Organization Management
// @Produce json
// @Security BearerAuth
// @param CreateOrganization body domain.Organizations{} true "Create organization"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/create-organization [post]
// Create Organization
func (cr *UserHandler) CreateOrganization(c *gin.Context) {

	var newOrganization domain.Organizations

	fmt.Println("Creating Organizations")
	//fetching data
	c.Bind(&newOrganization)

	fmt.Println("//////handler organization", newOrganization.OrganizationName)
	newOrganization.CreatedBy, _ = strconv.Atoi(c.Writer.Header().Get("user_id"))
	newOrganization.CreatedAt = time.Now()

	err := cr.userUseCase.CreateOrganization(newOrganization)
	fmt.Println(err)
	log.Println(newOrganization)

	if err != nil {
		response := response.ErrorResponse("Failed to create Organization", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	organization, _ := cr.userUseCase.FindOrganizationByName(newOrganization.OrganizationName)
	response := response.SuccessResponse(true, "SUCCESS", organization)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary update Profileabout
// @ID Update userprofile
// @Tags User Profile
// @Produce json
// @Security BearerAuth
// @param UpdateProfile body domain.Bios{} true "update profile with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/update-profile [patch]
func (cr *UserHandler) UpdateProfile(c *gin.Context) {

	var updatedProfile domain.Bios
	fmt.Println("Updating event")
	//fetching data
	c.Bind(&updatedProfile)

	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)

	//check event exit or not

	err := cr.userUseCase.UpdateProfile(updatedProfile, user_id)
	fmt.Println("error on updaed profile", err)

	log.Println(updatedProfile)

	if err != nil && err != sql.ErrNoRows {
		response := response.ErrorResponse("Failed to Update Profile", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.userUseCase.FindUserById(user_id)
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

type UpdatePassword struct {
	password string
}

// @Summary update password
// @ID Update password
// @Tags User Profile
// @Produce json
// @Param  email   query  string  true  "Email: "
// @param Updatepassword body UpdatePassword{} true "update password with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/update-password [patch]
func (cr *UserHandler) UpdatePassword(c *gin.Context) {

	var updatedPassword = UpdatePassword{}
	fmt.Println("Updating event")
	//fetching data
	c.Bind(&updatedPassword)
	fmt.Println("userPassword", updatedPassword.password)
	email := c.Query("email")

	//check event exit or not

	err := cr.userUseCase.UpdatePassword(updatedPassword.password, email)

	fmt.Println(updatedPassword.password)

	if err != nil {
		response := response.ErrorResponse("Failed to Update Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user, _ := cr.userUseCase.FindUserByName(email)
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// SendVerificationEmail sends the verification email

// @Summary Send verification
// @ID Send verifiation code via email
// @Tags Verification mail
// @Produce json
// @Param  email   query  string  true  "Email: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/send-verification [post]
func (cr *UserHandler) SendVerificationMail(c *gin.Context) {

	var email string
	email = c.Query("email")
	// if err != nil {
	// 	fmt.Println("error on binding the email")
	// }
	var code int
	fmt.Println(code)
	_, err := cr.userUseCase.FindUserByName(email)
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
// @Tags FAQA-user
// @Produce json
// @Param  Event_id   query  string  true  "Event_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-faqas [get]
func (cr *UserHandler) GetPublicFaqas(c *gin.Context) {

	Event_id, _ := strconv.Atoi(c.Query("Event_id"))
	faqas, err := cr.userUseCase.GetPublicFaqas(Event_id)
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
// @Tags FAQA-organization-volunteers>
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id :"
// @Param  Event_id   query  int  true  "Event_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/list-questions [get]
func (cr *UserHandler) GetQuestions(c *gin.Context) {

	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	Event_id, _ := strconv.Atoi(c.Query("Event_id"))
	questions, err := (cr.userUseCase.GetQuestions(Event_id))
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
// @Tags FAQA-user
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id :"
// @Param  Event_id   query  int  true  "Event_id: "
// @param PostQuestion body domain.Faqas{} true "Post question"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/event/post-question [post]
// PostQuesition handles Posting events
func (cr *UserHandler) PostQuestion(c *gin.Context) {

	var question domain.Faqas
	Event_id, _ := strconv.Atoi(c.Query("Event_id"))
	Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	user, _ := cr.userUseCase.FindUserById(user_id)

	c.Bind(&question)

	question.EventId = Event_id
	question.UserId = user_id
	question.OrganizationId = Organization_id
	question.UserName = user.UserName
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
// @Tags FAQA-organization-volunteers>
// @Produce json
// @Security BearerAuth
// @Param Organization_id query int true "Organization_id :"
// @param faqaid query int true "Getting the id of the question"
// @param PostAnswer body domain.Answers{} true "Post Answer"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/post-answer [post]
// PostQuesition handles Posting events
func (cr *UserHandler) PostAnswer(c *gin.Context) {

	var answer domain.Answers
	question_id, _ := strconv.Atoi(c.Query("faqaid"))
	user_id, _ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id ", user_id)
	Organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
	fmt.Println("Organization_id ", Organization_id)
	role := c.Writer.Header().Get("role")
	fmt.Println("role ", role)

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

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
