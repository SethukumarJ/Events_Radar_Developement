package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
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
// @Tags Admin
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
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  orgstatusid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/organization/register [patch]
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
	response := response.SuccessResponse(true, "Organization rejected", orgStatusId)
	utils.ResponseJSON(*c, response)

}
// @Summary Rejects the organization
// @ID Reject organization
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  orgstatusid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/organization/reject [patch]
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
	response := response.SuccessResponse(true, "Organization registered", orgStatusId)
	utils.ResponseJSON(*c, response)

}


// @Summary makes the user vip
// @ID make vip user
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  username   query  string  true  "User Name : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/vipuser [patch]
func (cr *AdminHandler) VipUser(c *gin.Context)  {
	
	username := c.Query("username")

	err := cr.adminUsecase.VipUser(username)

	if err != nil {
		response := response.ErrorResponse("making user into vip faled!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "User made into vip", username)
	utils.ResponseJSON(*c, response)

}

// @Summary approves the event for admin
// @ID approves event
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  title   query  string  true  "Event Name : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/approveevent [patch]
func (cr *AdminHandler) ApproveEvent(c *gin.Context)  {
	
	title := c.Query("title")

	err := cr.adminUsecase.ApproveEvent(title)

	if err != nil {
		response := response.ErrorResponse("approving event failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "event approved", title)
	utils.ResponseJSON(*c, response)

}


// @Summary list all upcoming events for admin
// @ID list all upcoming events
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Param  approved   query  bool  true  "List event by approved non approved : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listEvents [get]
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
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param  page   query  string  true  "Page number: "
// @Param  pagesize   query  string  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listUsers [get]
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
