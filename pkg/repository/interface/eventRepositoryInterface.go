package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type EventRepository interface {
	
	FindEvent(title string) (domain.EventResponse, error)
	FindUser(username string) (string, error)
	CreateEvent(event domain.Events) (int, error)
	UpdateEvent(event domain.Events,title string) (int, error)
	DeleteEvent(title string) error
	AllApprovedEvents(pagenation utils.Filter) ([]domain.EventResponse, utils.Metadata, error)
}