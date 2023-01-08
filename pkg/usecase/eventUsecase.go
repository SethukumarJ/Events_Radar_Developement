package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecases "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type eventUsecase struct {
	eventRepo interfaces.EventRepository
}

// AllEvents implements interfaces.EventUsecase
func (*eventUsecase) AllEvents(pagenation utils.Filter) (*[]domain.EventResponse, *utils.Metadata, error) {
	panic("unimplemented")
}

func NewEventUseCase(
	eventRepo interfaces.EventRepository,
) usecases.EventUsecase {
	return &eventUsecase{
		eventRepo: eventRepo,
	}
}

// CreateUser implements interfaces.UserUseCase
func (c *eventUsecase) CreateEvent(event domain.Events) error {
	fmt.Println("create event from service")
	_, err := c.eventRepo.FindEvent(event.Title)
	fmt.Println("found event", err)

	if err == nil {
		return errors.New("eventtitle already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	_, err = c.eventRepo.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

// FindUser implements interfaces.UserUseCase
func (c *eventUsecase) FindEvent(title string) (*domain.EventResponse, error) {
	event, err := c.eventRepo.FindEvent(title)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
