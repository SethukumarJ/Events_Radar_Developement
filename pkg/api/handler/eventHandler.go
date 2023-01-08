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

type EventHandler struct {
	eventUsecase usecase.EventUseCase
}

func NewEventHandler(usecase usecase.EventUseCase) EventHandler {
	return EventHandler{
		eventUsecase: usecase,
	}
}

func (cr *EventHandler) CreateEvent(c *gin.Context) {

	var newEvent domain.Events
	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newEvent)

	//check username exit or not

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
