package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type EventUsecase interface {
	CreateEvent(event domain.Events) error
	DeleteEvent(event_id int) error
	FindEventById(event_id int) (*domain.EventResponse, error)
	FindEventByTitle(title string) (*domain.EventResponse, error)
	AllApprovedEvents(pagenation utils.Filter , filter utils.FilterEvent) (*[]domain.EventResponse, *utils.Metadata, error)
	FindOrganizationById(organization_id int) (*domain.OrganizationsResponse, error)
	CreatePoster(event domain.Posters) error
	DeletePoster(poster_id int,event_id int) error
	FindPosterByName(name string,event_id int) (*domain.PosterResponse, error)
	FindPosterById(poster_id int,event_id int) (*domain.PosterResponse, error)
	PostersByEvent(event_id int) (*[]domain.PosterResponse, error)
	FindUser(username string) (bool,error)
	UpdateEvent(event domain.Events, event_id int) error
	SearchEventUser(search string) (*[]domain.EventResponse, error)
	ListApplications(pagenation utils.Filter, applicationStatus string,event_id int) (*[]domain.ApplicationFormResponse, *utils.Metadata, error)
	AcceptApplication(applicationStatusId int,event_id int) error
	RejectApplication(applicationStatusId int,event_id int) error
}
