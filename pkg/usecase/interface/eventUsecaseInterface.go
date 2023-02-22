package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type EventUsecase interface {
	CreateEvent(event domain.Events) error
	UpdateEvent(event domain.Events, title string) error
	DeleteEvent(title string) error
	FindEvent(title string) (*domain.EventResponse, error)
	FindUser(username string) (bool,error)
	AllApprovedEvents(pagenation utils.Filter , filter utils.FilterEvent) (*[]domain.EventResponse, *utils.Metadata, error)
	SearchEventUser(search string) (*[]domain.EventResponse, error)
}
