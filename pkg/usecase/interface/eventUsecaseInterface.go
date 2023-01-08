package interfaces

import domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"

type EventUseCase interface {
	CreateEvent(event domain.Events) error
	FindEvent(email string) (*domain.EventResponse, error)
}