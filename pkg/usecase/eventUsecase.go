package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecases "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type eventUsecase struct {
	eventRepo interfaces.EventRepository
}

// DeleteEvent implements interfaces.EventUsecase
func (c *eventUsecase) DeleteEvent(title string) error {
	err := c.eventRepo.DeleteEvent(title)

	if err != nil {
		return nil
	}

	return nil
}

// UpdateEvent implements interfaces.EventUsecase
func (c *eventUsecase) UpdateEvent(event domain.Events, title string) error {
	fmt.Println("update event from service")
	_, err := c.eventRepo.FindEvent(title)
	fmt.Println("found event", err)

	if err == nil {
		log.Printf("found event")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	_, err = c.eventRepo.UpdateEvent(event, title)
	if err != nil {
		return err
	}
	return nil
}

// AllEvents implements interfaces.EventUsecase
func (c *eventUsecase) AllApprovedEvents(pagenation utils.Filter) (*[]domain.EventResponse, *utils.Metadata, error) {
	fmt.Println("allevents from usecase called")
	events, metadata, err := c.eventRepo.AllApprovedEvents(pagenation)
	fmt.Println("events:", events)
	if err != nil {
		fmt.Println("error from allevents usecase:", err)
		return nil, &metadata, err
	}

	return &events, &metadata, nil
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
	events, err := c.eventRepo.FindEvent(event.Title)
	fmt.Println("found event", events.Title)

	if err == nil {
		return errors.New("eventtitle already exists")
	}

	if err == nil && err != sql.ErrNoRows {
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
