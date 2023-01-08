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

type EventHandler struct {
	eventUsecase usecase.EventUsecase
}

func NewEventHandler(usecase usecase.EventUsecase) EventHandler {
	return EventHandler{
		eventUsecase: usecase,
	}
}

func (cr *EventHandler) CreateEvent(c *gin.Context) {

	var newEvent domain.Events
	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newEvent)
	fmt.Println("event id", newEvent.EventId)

	//check event exit or not

	err := cr.eventUsecase.CreateEvent(newEvent)

	log.Println(newEvent)

	if err != nil {
		response := response.ErrorResponse("Failed to create Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	event, _ := cr.eventUsecase.FindEvent(newEvent.Title)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

func (cr *EventHandler) ViewAllApprovedEvents(c *gin.Context) {

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

	events, metadata, err := cr.eventUsecase.AllApprovedEvents(pagenation)

	fmt.Println("events:", events)

	result := struct {
		Events *[]domain.EventResponse
		Meta   *utils.Metadata
	}{
		Events: events,
		Meta:   metadata,
	}

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
