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



func (cr *AdminHandler) ViewAllEvents(c *gin.Context) {

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

	events, metadata, err := cr.adminUsecase.AllEvents(pagenation)

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
