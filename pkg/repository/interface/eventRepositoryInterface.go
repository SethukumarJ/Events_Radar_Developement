package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type EventRepository interface {
	
	CreatePoster(event domain.Posters) (int,error)
	DeletePoster(name string, eventid int) error
	FindPoster(title string, eventid int) (domain.PosterResponse, error)
	PostersByEvent(eventid int) ([]domain.PosterResponse, error)
	FindEvent(title string) (domain.EventResponse, error)
	FindUser(username string) (string, error)
	CreateEvent(event domain.Events) (int, error)
	UpdateEvent(event domain.Events,title string) (int, error)
	DeleteEvent(title string) error
	AllApprovedEvents(pagenation utils.Filter, filter  utils.FilterEvent) ([]domain.EventResponse, utils.Metadata, error)
	SearchEventUser(search string) ([]domain.EventResponse,error)
	ListApplications(pagenation utils.Filter,applicationStatus string) ([]domain.ApplicationFormResponse, utils.Metadata, error)
	AcceptApplication(applicationStatusId int) error
	RejectApplication(applicationStatusId int) error
}