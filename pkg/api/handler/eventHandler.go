package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type EventHandler struct {
	eventUsecase usecase.EventUsecase
}

func NewEventHandler(usecase usecase.EventUsecase) EventHandler {
	return EventHandler{
		eventUsecase: usecase,
	}
}


// @Summary list all pending application of participant
// @ID list all application with status
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Param  applicationStatus   query  string  true  "List organization based on status: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/list-application [get]
func (cr *EventHandler) ListApplications(c *gin.Context) {


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

	applications, metadata, err := cr.eventUsecase.ListApplications(pagenation, applicationStatus)

	fmt.Println("applications:", applications)

	result := struct {
		ApplicationForm *[]domain.ApplicationFormResponse
		Meta  *utils.Metadata
	}{
		ApplicationForm: applications,
		Meta:  metadata,
	}

	if err != nil {
		response := response.ErrorResponse("error while getting event applications from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All participant applications", result)
	utils.ResponseJSON(*c, response)

}




// @Summary accept the application for participate in the event
// @ID Accept application
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  applicationstsid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/accept-application [patch]
func (cr *EventHandler) AcceptApplication(c *gin.Context)  {
	
	applicationstsid,_ := strconv.Atoi(c.Query("applicationstsid"))

	err := cr.eventUsecase.AcceptApplication(applicationstsid)

	if err != nil {
		response := response.ErrorResponse("accepting request failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "application accepted", applicationstsid)
	utils.ResponseJSON(*c, response)

}
// @Summary Rejects the application for participate in the event
// @ID Reject application
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  applicationstsid   query  int  true  "applicationstsid  : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/reject-application [patch]
func (cr *EventHandler) RejectApplication(c *gin.Context)  {
	
	applicationstsid,_ := strconv.Atoi(c.Query("applicationstsid"))

	err := cr.eventUsecase.RejectApplication(applicationstsid)

	if err != nil {
		response := response.ErrorResponse("Rjecting applicaiton failed!", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Application rejected", applicationstsid)
	utils.ResponseJSON(*c, response)

}

// @Summary delete event
// @ID Delete event
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  title   query  string  true  "Title: "
// @Param organizationName query string true "organizationName: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/delete-event [delete]
func (cr *EventHandler) DeleteEvent(c *gin.Context) {


	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	title := c.Query("title")

	err := cr.eventUsecase.DeleteEvent(title)

	if err != nil {
		response := response.ErrorResponse("Could not delete event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Deleted event successfully!", title)
	utils.ResponseJSON(*c, response)

}

// @Summary update event
// @ID Update event
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @param title query string true "event title"
// @Param organizationName query string true "organizationName: "
// @param UpdateEvent body domain.Events{} true "update Event with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/update-event [patch]
func (cr *EventHandler) UpdateEvent(c *gin.Context) {



	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	var updatedEvent domain.Events
	fmt.Println("Updating event")
	//fetching data
	err:=c.Bind(&updatedEvent)

	fmt.Println("event ////////////", updatedEvent, "errror",err)
	title := c.Query("title")

	//check event exit or not

	err = cr.eventUsecase.UpdateEvent(updatedEvent, title)

	log.Println(updatedEvent)

	if err != nil {
		response := response.ErrorResponse("Failed to Update Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	event, _ := cr.eventUsecase.FindEvent(updatedEvent.Title)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary Create event
// @ID Create event from user
// @Tags User
// @Produce json
// @Security BearerAuth
// @param CreateEvent body domain.Events{} true "Create event"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/create-event [post]
// Create events
func (cr *EventHandler) CreateEventUser(c *gin.Context) {

	var newEvent domain.Events

	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newEvent)


	fmt.Println("event", newEvent)
	newEvent.OrganizerName = c.Writer.Header().Get("userName")
	newEvent.CreatedAt = time.Now()
	vip,err:= cr.eventUsecase.FindUser(newEvent.OrganizerName)
	if err != nil {
		fmt.Println(err)
	}

	if vip {
		newEvent.Approved = true
	}


	//check event exit or not

	err = cr.eventUsecase.CreateEvent(newEvent)

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

// @Summary Create event
// @ID Create event from admin
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @param CreateEvent body domain.Events{} true "Create event"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/create-event [post]
// Create events
func (cr *EventHandler) CreateEventAdmin(c *gin.Context) {

	var newEvent domain.Events
	fmt.Println("event", newEvent)
	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newEvent)
	newEvent.OrganizerName = c.Writer.Header().Get("userName")
	newEvent.CreatedAt = time.Now()
	newEvent.Approved = true
	

	fmt.Println(newEvent.OrganizerName,newEvent.CreatedAt,newEvent.Approved,newEvent.Title)

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

 
// @Summary Create event by organization
// @ID Create event from organization
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param organizationName query string true "organizationName"
// @param CreateEvent body domain.Events{} true "Create event"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/create-event [post]
// Create events
func (cr *EventHandler) CreateEventOrganization(c *gin.Context) {

	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}


	var newEvent domain.Events
	fmt.Println("event", newEvent)
	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newEvent)
	newEvent.OrganizerName = c.Writer.Header().Get("organizationName")
	newEvent.CreatedAt = time.Now()
	newEvent.Approved = true
	

	fmt.Println(newEvent.OrganizerName,newEvent.CreatedAt,newEvent.Approved,newEvent.Title)

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

// @Summary get event by title
// @ID Get event by title
// @Tags User
// @Produce json
// @Param  title   query  string  true  "Title: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/geteventbytitle [get]
func (cr *EventHandler) GetEventByTitle(c *gin.Context) {

	title := c.Query("title")

	event, err := cr.eventUsecase.FindEvent(title)

	fmt.Println("event:", event)

	if err != nil {
		response := response.ErrorResponse("error while getting event from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Showing the event", event)
	utils.ResponseJSON(*c, response)

}
 

// @Summary Create Poster by organization
// @ID Create Poster from organization
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param EventName query string true "EventName"
// @param CreatePoster body domain.Posters{} true "Create poster"
// @Param organizationName query string true "organizationName: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /event/create-poster [post]
// Create posters
func (cr *EventHandler) CreatePosterOrganization(c *gin.Context) {

	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}


	var newPoster domain.Posters
	fmt.Println("event", newPoster)
	fmt.Println("Creating event")
	//fetching data
	c.Bind(&newPoster)
	now := time.Now()
	newPoster.Date  =  now.Format("2006-01-02")
	

	//check event exit or not

	err := cr.eventUsecase.CreatePoster(newPoster)

	log.Println(newPoster)

	if err != nil {
		response := response.ErrorResponse("Failed to create Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	poster, _ := cr.eventUsecase.FindPoster(newPoster.Name,int(newPoster.EventId))
	response := response.SuccessResponse(true, "SUCCESS", poster)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}


// @Summary Search Event from user side
// @ID search sdf with string by user
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param  search   body  string  true  "List event by approved non approved : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/search-event [get]
func (cr *EventHandler) SearchEventUser(c *gin.Context) {
	var search string
	c.Bind(&search)
	
	events, err := cr.eventUsecase.SearchEventUser(search)

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
// @Summary Search Event from user side
// @ID search event with string by user
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param  Eventid   query  int  true  "Posters under event : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /event/get-Posters [get]
func (cr *EventHandler) PostersByEvent(c *gin.Context) {
	

	eventId,_ := strconv.Atoi(c.Query("Eventid"))

	posters, err := cr.eventUsecase.PostersByEvent(eventId)

	fmt.Println("events:", posters)

	if err != nil {
		response := response.ErrorResponse("error while getting posters from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}


	response := response.SuccessResponse(true, "posters", posters)
	utils.ResponseJSON(*c, response)

}

// @Summary delete poster
// @ID Delete poster
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  title   query  string  true  "Title: "
// @Param  eventid   query  int  true  "Title: "
// @Param organizationName query string true "organizationName: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /event/delete-poster [delete]
func (cr *EventHandler) DeletePoster(c *gin.Context) {


	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	title := c.Query("title")
	eventId,_ := strconv.Atoi(c.Query("Eventid"))
	err := cr.eventUsecase.DeletePoster(title,eventId)

	if err != nil {
		response := response.ErrorResponse("Could not delete event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Deleted event successfully!", title)
	utils.ResponseJSON(*c, response)

}

// @Summary Get poster by title
// @ID Get event by id
// @Tags User
// @Produce json
// @Param  title   query  string  true  "Title: "
// @Param  eventid   query  int  true  "Title: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /event/getposterbytitle [get]
func (cr *EventHandler) GetPosterByTitle(c *gin.Context) {

	title := c.Query("title")
	eventId,_ := strconv.Atoi(c.Query("Eventid"))
	poster, err := cr.eventUsecase.FindPoster(title,eventId)

	fmt.Println("event:", poster)

	if err != nil {
		response := response.ErrorResponse("error while getting poster from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Showing the event", poster)
	utils.ResponseJSON(*c, response)

}
// @Summary list all approved upcoming events
// @ID list all approved events
// @Tags User
// @Produce json
// @Param  page   query  string  true  "Page number: "
// @Param  cusatonly   query  bool  true  "Cusat only: "
// @Param  online   query  bool  true  "Online: "
// @Param  sex   query  string  true  "sex: "
// @Param  pagesize   query  string  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list/approved-events [get]
func (cr *EventHandler) ViewAllApprovedEvents(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	cusatonly := c.Query("cusatonly")
	cusatOnly, _ := strconv.ParseBool(cusatonly)
	online := c.Query("online")
	onLine, _ := strconv.ParseBool(online)
	sex := c.Query("sex")

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	log.Println(page, "   ", pageSize)

	fmt.Println("page :", page)
	fmt.Println("pagesize", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	filter := utils.FilterEvent{
		CusatOnly: cusatOnly,
		Online: onLine,
		Sex: sex,
	}


	fmt.Println("pagenation", pagenation)

	events, metadata, err := cr.eventUsecase.AllApprovedEvents(pagenation,filter)

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
