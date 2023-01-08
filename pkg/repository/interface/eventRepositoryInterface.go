package interfaces

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

// UserRepository represent the users's repository contract
type EventRepository interface {
	FindEvent(title string) (domain.EventResponse, error)
	CreateEvent(event domain.Events) (int, error)
	AllApprovedEvents(pagenation utils.Filter) ([]domain.EventResponse, utils.Metadata, error)
}
