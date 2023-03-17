package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type EventRepository interface {
	
	CreatePoster(event domain.Posters) (int,error)
	DeletePoster(poster_id int, eventid int) error
	FindPosterByName(title string, eventid int) (domain.PosterResponse, error)
	FindPosterById(poster_id int, eventid int) (domain.PosterResponse, error)
	PostersByEvent(eventid int) ([]domain.PosterResponse, error)
	FindEventByTitle(title string) (domain.EventResponse, error)
	FindEventById(event_id int) (domain.EventResponse, error)
	FindUser(username string) (string, error)
	CreateEvent(event domain.Events) (int, error)
	UpdateEvent(event domain.Events,event_id int) (int, error)
	DeleteEvent(event_id int) error
	AllApprovedEvents(pagenation utils.Filter, filter  utils.FilterEvent) ([]domain.EventResponse, utils.Metadata, error)
	SearchEventUser(search string) ([]domain.EventResponse,error)
	ListApplications(pagenation utils.Filter,applicationStatus string,event_id int) ([]domain.ApplicationFormResponse, utils.Metadata, error)
	AcceptApplication(applicationStatusId int,event_id int) error
	RejectApplication(applicationStatusId int,event_id int) error
	FindOrganizationById(organizaiton_id int) (domain.OrganizationsResponse, error)
}