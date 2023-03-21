package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
	userUsecase  usecase.UserUseCase
	eventUsecase usecase.EventUsecase
}

func NewAdminHandler(
	adminUsecase usecase.AdminUsecase,
	userUsecase usecase.UserUseCase,
	eventUsecase usecase.EventUsecase,
) AdminHandler {
	return AdminHandler{
		adminUsecase: adminUsecase,
		userUsecase:  userUsecase,
		eventUsecase: eventUsecase,
	}
}

// @Summary list all pending organizations for admin
// @ID list all organization with status
// @Tags Admin-Organization Management
// @Produce json
// @Security BearerAuth
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Param  applicationStatus   query  string  true  "List organization based on status: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-organizations [get]
func (cr *AdminHandler) ListOrgRequests(c *gin.Context) {


	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	applicationStatus:= c.Query("applicationStatus")
	fmt.Println("applicationStatus",applicationStatus)

	log.Println(page, "   ", pageSize)

	fmt.Println("page :", page)
	fmt.Println("pagesize", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	fmt.Println("pagenation", pagenation)

	organizations, metadata, err := cr.adminUsecase.ListOrgRequests(pagenation, applicationStatus)

	fmt.Println("events:", organizations)

	result := struct {
		Organizations *[]domain.OrganizationsResponse
		Meta  *utils.Metadata
	}{
		Organizations: organizations,
		Meta:  metadata,
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




// @Summary Resginter the organization
// @ID Register organization
// @Tags Admin-Organization Management
// @Produce json
// @Security BearerAuth
// @Param  orgstatusid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/register-organization [patch]
func (cr *AdminHandler) RegisterOrganization(c *gin.Context)  {
	
	orgStatusId,_ := strconv.Atoi(c.Query("orgstatusid"))

	err := cr.adminUsecase.RegisterOrganization(orgStatusId)

	if err != nil {
		response := response.ErrorResponse("Registering organization failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Organization registered", orgStatusId)
	utils.ResponseJSON(*c, response)

}
// @Summary Rejects the organization
// @ID Reject organization
// @Tags Admin-Organization Management
// @Produce json
// @Security BearerAuth
// @Param  orgstatusid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/reject-organization [patch]
func (cr *AdminHandler) RejectOrganization(c *gin.Context)  {
	
	orgStatusId,_ := strconv.Atoi(c.Query("orgstatusid"))

	err := cr.adminUsecase.RejectOrganization(orgStatusId)

	if err != nil {
		response := response.ErrorResponse("Registering organization failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Organization rejected", orgStatusId)
	utils.ResponseJSON(*c, response)

}


// @Summary makes the user vip
// @ID make vip user
// @Tags Admin-User Profile
// @Produce json
// @Security BearerAuth
// @Param  User_id   query  int  true  "User Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/vipfy-user [patch]
func (cr *AdminHandler) VipUser(c *gin.Context)  {
	
	User_id,_ := strconv.Atoi(c.Query("User_id"))


	err := cr.adminUsecase.VipUser(User_id)

	if err != nil {
		response := response.ErrorResponse("making user into vip faled!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "User made into vip", User_id)
	utils.ResponseJSON(*c, response)

}

// @Summary approves the event for admin
// @ID approves event
// @Tags Admin-Event Management
// @Produce json
// @Security BearerAuth
// @Param  Event_id   query  string  true  "Event Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/approve-event [patch]
func (cr *AdminHandler) ApproveEvent(c *gin.Context)  {
	
	Event_id,_ := strconv.Atoi(c.Query("Event_id"))

	err := cr.adminUsecase.ApproveEvent(Event_id)

	if err != nil {
		response := response.ErrorResponse("approving event failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "event approved", Event_id)
	utils.ResponseJSON(*c, response)

}

// @Summary Search Event
// @ID search event with string
// @Tags Admin-Event Management
// @Produce json
// @Param  search   query  string  true  "search string: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/search-event [get]
func (cr *AdminHandler) SearchEvent(c *gin.Context) {
	var search = c.Query("search")
	
	
	events, err := cr.adminUsecase.SearchEvent(search)

	fmt.Println("events:", events)

	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Search result", events)
	utils.ResponseJSON(*c, response)

}



// @Summary list all upcoming events for admin
// @ID list all upcoming events
// @Tags Admin-Event Management
// @Produce json
// @Security BearerAuth
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Param  approved   query  bool  true  "List event by approved non approved : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-events [get]
func (cr *AdminHandler) ViewAllEvents(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	approved:= c.Query("approved")
	fmt.Println("approved",approved)

	log.Println(page, "   ", pageSize)

	fmt.Println("page :", page)
	fmt.Println("pagesize", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	fmt.Println("pagenation", pagenation)

	events, metadata, err := cr.adminUsecase.AllEvents(pagenation,approved)

	fmt.Println("events:", events)

	result := struct {
		Events *[]domain.EventResponse
		Meta  *utils.Metadata
	}{
		Events: events,
		Meta:  metadata,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Users", result)
	utils.ResponseJSON(*c, response)

}


// @Summary list all active users for admin
// @ID list all active users
// @Tags Admin-User Profile
// @Produce json
// @Security BearerAuth
// @Param  page   query  string  true  "Page number: "
// @Param  pagesize   query  string  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-users [get]
func (cr *AdminHandler) ViewAllUsers(c *gin.Context) {

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

	users, metadata, err := cr.adminUsecase.AllUsers(pagenation)

	fmt.Println("users:", users)

	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Users", result)
	utils.ResponseJSON(*c, response)

}
