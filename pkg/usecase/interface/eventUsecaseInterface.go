package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type EventUsecase interface {
	CreateEvent(event domain.Events) error
	DeleteEvent(title string) error
	FindEvent(title string) (*domain.EventResponse, error)
	AllApprovedEvents(pagenation utils.Filter , filter utils.FilterEvent) (*[]domain.EventResponse, *utils.Metadata, error)
	CreatePoster(event domain.Posters) error
	DeletePoster(name string) error
	FindPoster(title string) (*domain.PosterResponse, error)
	PostersByEvent(eventid int) (*[]domain.PosterResponse, error)
	FindUser(username string) (bool,error)
	UpdateEvent(event domain.Events, title string) error
	SearchEventUser(search string) (*[]domain.EventResponse, error)
}
