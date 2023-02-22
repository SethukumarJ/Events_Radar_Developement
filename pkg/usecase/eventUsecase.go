package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	interfaces "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	usecases "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type eventUsecase struct {
	eventRepo interfaces.EventRepository
}

// FindUser implements interfaces.EventUsecase
func (c *eventUsecase) FindUser(username string) (bool, error) {
	vip,err := c.eventRepo.FindUser(username) 
	if err != nil {
		return false ,err
	}

	if vip == "false"{
		return false,nil
	} 
	return true, nil


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
	fmt.Println("found event", err,)

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
func (c *eventUsecase) AllApprovedEvents(pagenation utils.Filter, filter utils.FilterEvent) (*[]domain.EventResponse, *utils.Metadata, error) {
	fmt.Println("allevents from usecase called")
	events, metadata, err := c.eventRepo.AllApprovedEvents(pagenation,filter)
	fmt.Println("events:", events)
	if err != nil {
		fmt.Println("error from allevents usecase:", err)
		return nil, &metadata, err
	}

	return &events, &metadata, nil
}

// CreateUser implements interfaces.UserUseCase
func (c *eventUsecase) CreateEvent(event domain.Events) error {
	fmt.Println("create user from service")
	_, err := c.eventRepo.FindEvent(event.Title)
	fmt.Println("found event", err)

	if err == nil {
		return errors.New("event title already exists")
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

func NewEventUseCase(
	eventRepo interfaces.EventRepository,
) usecases.EventUsecase {
	return &eventUsecase{
		eventRepo: eventRepo,
	}
}
