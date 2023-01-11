package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type EventUsecase interface {
	CreateEvent(event domain.Events) error
	UpdateEvent(event domain.Events, title string) error
	DeleteEvent(title string) error
	FindEvent(title string) (*domain.EventResponse, error)
	FindUser(username string) (bool,error)
	AllApprovedEvents(pagenation utils.Filter) (*[]domain.EventResponse, *utils.Metadata, error)
}
