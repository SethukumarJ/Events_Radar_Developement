package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventUsecase usecase.EventUsecase
	adminUsecase usecase.AdminUsecase
	userUsecase  usecase.UserUseCase
}

func NewEventHandler(adminUsecase usecase.AdminUsecase,
	userUsecase usecase.UserUseCase,
	eventUsecase usecase.EventUsecase,) EventHandler {
	return EventHandler{
		adminUsecase: adminUsecase,
		userUsecase:  userUsecase,
		eventUsecase: eventUsecase,
	}
}


// @Summary list all pending application of participant
// @ID list all application with status
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param Event_id query int true "Event_id: "
// @Param  page   query  int  true  "Page number: "
// @Param  pagesize   query  int  true  "Page capacity : "
// @Param Organization_id query int true "Organization_id: "
// @Param  applicationStatus   query  string  true  "List application based on status: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/list-applications [get]
func (cr *EventHandler) ListApplications(c *gin.Context) {


	page, _ := strconv.Atoi(c.Query("page"))
	
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	Event_id,_ := strconv.Atoi(c.Query("Event_id"))
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

	applications, metadata, err := cr.eventUsecase.ListApplications(pagenation, applicationStatus,Event_id)

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
// @Param Organization_id query int true "Organization_id: "
// @Param Event_id query int true "Event_id: "
// @Param  applicationstsid   query  int  true  "orgStatus id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/accept-application [patch]
func (cr *EventHandler) AcceptApplication(c *gin.Context)  {
	
	applicationstsid,_ := strconv.Atoi(c.Query("applicationstsid"))
	Event_id,_ :=  strconv.Atoi(c.Query("Event_id"))
	err := cr.eventUsecase.AcceptApplication(applicationstsid,Event_id)

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
// @Param Organization_id query int true "Organization_id: "
// @Param Event_id query int true "Event_id: "
// @Param  applicationstsid   query  int  true  "applicationstsid  : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/reject-application [patch]
func (cr *EventHandler) RejectApplication(c *gin.Context)  {
	
	applicationstsid,_ := strconv.Atoi(c.Query("applicationstsid"))
	Event_id,_ :=  strconv.Atoi(c.Query("Event_id"))
	err := cr.eventUsecase.RejectApplication(applicationstsid,Event_id)

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
// @Param  Event_id   query  int  true  "Event_id: "
// @Param Organization_id query int true "Organizatiion_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/delete-event [delete]
func (cr *EventHandler) DeleteEvent(c *gin.Context) {


	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	Event_id,_ :=  strconv.Atoi(c.Query("Event_id"))

	err := cr.eventUsecase.DeleteEvent(Event_id)

	if err != nil {
		response := response.ErrorResponse("Could not delete event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Deleted event successfully!", Event_id)
	utils.ResponseJSON(*c, response)

}

// @Summary update event
// @ID Update event
// @Tags Organization
// @Produce json
// @Security BearerAuth
// @Param  Event_id   query  int  true  "Event_id: "
// @Param Organization_id query int true "Organizatiion_id: "
// @param UpdateEvent body domain.Events{} true "update Event with new body"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/update-event [patch]
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
	Event_id,_ :=  strconv.Atoi(c.Query("Event_id"))

	//check event exit or not

	err = cr.eventUsecase.UpdateEvent(updatedEvent, Event_id)

	log.Println(updatedEvent)

	if err != nil {
		response := response.ErrorResponse("Failed to Update Event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	event, _ := cr.eventUsecase.FindEventById(Event_id)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary Create event
// @ID Create event from user
// @Tags User-Event Management
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
	newEvent.CreatedBy = "user"
	newEvent.User_id,_ =  strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user id", newEvent.User_id)

	newEvent.CreatedAt = time.Now()
	user, err := cr.userUsecase.FindUserById(newEvent.User_id)
	fmt.Println("user", user)
	if err != nil {
		err = errors.New("error while getting user")
		response := response.ErrorResponse("FAIL", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
		
	}
	vip,err:= cr.eventUsecase.FindUser(user.UserName)
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

	event, _ := cr.eventUsecase.FindEventByTitle(newEvent.Title)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary Create event
// @ID Create event from admin
// @Tags Admin-Event Management
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
	
	newEvent.CreatedBy = "admin"
	newEvent.User_id,_ =  strconv.Atoi(c.Writer.Header().Get("user_id"))
	newEvent.CreatedAt = time.Now()
	newEvent.Approved = true
	

	fmt.Println(newEvent.CreatedBy,newEvent.CreatedAt,newEvent.Approved,newEvent.Title)

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

	event, _ := cr.eventUsecase.FindEventByTitle(newEvent.Title)
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
// @Param Organization_id query int true "Organization_id"
// @param CreateEvent body domain.Events{} true "Create event"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/create-event [post]
// Create events
func (cr *EventHandler) CreateEventOrganization(c *gin.Context) {

	role := c.Writer.Header().Get("role")
	organization_id,_ := strconv.Atoi(c.Writer.Header().Get("organization_id"))
	fmt.Println("organization_id",organization_id)
	user_id,_ := strconv.Atoi(c.Writer.Header().Get("user_id"))
	fmt.Println("user_id",user_id)
	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	var newEvent domain.Events
	c.Bind(&newEvent)

	_,err :=cr.userUsecase.FindOrganizationById(organization_id)

	if err != nil {
		response := response.ErrorResponse("No Organization found", "no value", err)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	
	fmt.Println("event", newEvent)
	fmt.Println("Creating event",newEvent.EventDate)
	//fetching data
	newEvent.CreatedBy = "organization"
	newEvent.User_id = user_id
	newEvent.OrganizationId,_ = strconv.Atoi(c.Writer.Header().Get("Organization_id"))
	newEvent.CreatedAt = time.Now()
	newEvent.Approved = true
	

	fmt.Println(newEvent.CreatedBy,newEvent.CreatedAt,newEvent.Approved,newEvent.Title, newEvent.ApplicationClosingDate)
	newEvent.ApplicationLeft= newEvent.MaxApplications
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

	event, _ := cr.eventUsecase.FindEventByTitle(newEvent.Title)
	response := response.SuccessResponse(true, "SUCCESS", event)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}

// @Summary get event by title
// @ID Get event by title
// @Tags User-Event Management
// @Produce json
// @Param  Event_id   query  int  true  "Evnet_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/geteventbyid [get]
func (cr *EventHandler) GetEventById(c *gin.Context) {

	event_id,_ := strconv.Atoi(c.Query("Event_id"))

	event, err := cr.eventUsecase.FindEventById(event_id)

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
// @Tags Organization-Event-Poster Management
// @Produce json
// @Security BearerAuth
// @Param EventId query int true "Event id"
// @param CreatePoster body domain.Posters{} true "Create poster"
// @Param Organization_id query string true "Organization_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/create-poster [post]
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
	newPoster.EventId,_ =  strconv.Atoi(c.Query("EventId"))
	newPoster.Date  =  now.Format("2006-01-02")
	

	//check event exit or not

	err := cr.eventUsecase.CreatePoster(newPoster)

	log.Println(newPoster)

	if err != nil {
		response := response.ErrorResponse("Failed to create Posterr", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	poster, _ := cr.eventUsecase.FindPosterByName(newPoster.Name,int(newPoster.EventId))
	response := response.SuccessResponse(true, "SUCCESS", poster)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}


type SearchEvent struct {
	Search string
}

// @Summary Search Event from user side
// @ID search event with string by user
// @Tags User-Event Management
// @Produce json
// @Security BearerAuth
// @Param  search   query  string  true  "Search Eventt: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/search-event [get]
func (cr *EventHandler) SearchEventUser(c *gin.Context) {
	var search = c.Query("search")

	
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
// @ID Get posters Order by Event
// @Tags Organization-Event-Poster Management
// @Produce json
// @Security BearerAuth
// @Param  Eventid   query  int  true  "Posters under event : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/get-posters [get]
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
// @Tags Organization-Event-Poster Management
// @Produce json
// @Security BearerAuth
// @Param  Poster_id   query  int  true  "Poster_id: "
// @Param  event_id   query  int  true  "Event_id: "
// @Param Organization_id query string true "Organization_id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/delete-poster [delete]
func (cr *EventHandler) DeletePoster(c *gin.Context) {


	role := c.Writer.Header().Get("role")

	if role > "2" {
		response := response.ErrorResponse("Your role is not eligible for this action", "no value", nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	Poster_id,_ :=strconv.Atoi(c.Query("Poster_id"))
	event_id,_ := strconv.Atoi(c.Query("event_id"))
	err := cr.eventUsecase.DeletePoster(Poster_id,event_id)

	if err != nil {
		response := response.ErrorResponse("Could not delete event", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "Deleted poster successfully!", Poster_id)
	utils.ResponseJSON(*c, response)

}

// @Summary Get poster by title
// @ID Get event by id
// @Tags Organization-Event-Poster Management
// @Produce json
// @Param  Poster_id   query  int  true  "Poster_id: "
// @Param  eventid   query  int  true  "event id: "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /organization/event/get-posterbyid [get]
func (cr *EventHandler) GetPosterById(c *gin.Context) {

	Poster_id,_ :=strconv.Atoi(c.Query("Poster_id"))
	eventId,_ := strconv.Atoi(c.Query("eventid"))
	poster, err := cr.eventUsecase.FindPosterById(int(Poster_id),int(eventId))

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
// @Tags User-Event Management
// @Produce json
// @Param  page   query  string  true  "Page number: "
// @Param  cusatonly   query  bool  true  "Cusat only: "
// @Param  online   query  bool  true  "Online: "
// @Param  sex   query  string  true  "Sex (options: male, female, any)"
// @Param  pagesize   query  string  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-approved-events [get]
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
