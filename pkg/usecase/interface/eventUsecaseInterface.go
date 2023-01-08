package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type EventUsecase interface {
	CreateEvent(event domain.Events) error
	FindEvent(email string) (*domain.EventResponse, error)
	AllApprovedEvents(pagenation utils.Filter) (*[]domain.EventResponse, *utils.Metadata, error)
}
