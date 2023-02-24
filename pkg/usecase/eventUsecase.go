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

// AcceptApplication implements interfaces.EventUsecase
func (c *eventUsecase) AcceptApplication(applicationStatusId int) error {
	err := c.eventRepo.AcceptApplication(applicationStatusId)

	if err != nil {
		return err
	}
	return nil
}

// ListApplications implements interfaces.EventUsecase
func (c *eventUsecase) ListApplications(pagenation utils.Filter, applicationStatus string) (*[]domain.ApplicationFormResponse, *utils.Metadata, error) {
	fmt.Println("List applilcation from usecase called")
	applicaition, metadata, err := c.eventRepo.ListApplications(pagenation, applicationStatus)
	fmt.Println("applicaition:", applicaition)
	if err != nil {
		fmt.Println("error from list applicaition from usecase:", err)
		return nil, &metadata, err
	}

	return &applicaition, &metadata, nil
}

// RejectApplication implements interfaces.EventUsecase
func (c *eventUsecase) RejectApplication(applicationStatusId int) error {
	err := c.eventRepo.RejectApplication(applicationStatusId)

	if err != nil {
		return err
	}
	return nil
}

// CreatePoster implements interfaces.EventUsecase
func (c *eventUsecase) CreatePoster(poster domain.Posters) error {
	fmt.Println("create poster from service")
	_, err := c.eventRepo.FindPoster(poster.Name, int(poster.EventId))
	fmt.Println("found poster", err)

	if err == nil {
		return errors.New("poster name already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	_, err = c.eventRepo.CreatePoster(poster)
	if err != nil {
		return err

	}
	return nil
}

// DeletePoster implements interfaces.EventUsecase
func (c *eventUsecase) DeletePoster(name string, eventid int) error {
	err := c.eventRepo.DeletePoster(name, eventid)

	if err != nil {
		return nil
	}

	return nil
}

// FindPoster implements interfaces.EventUsecase
func (c *eventUsecase) FindPoster(title string, eventid int) (*domain.PosterResponse, error) {
	poster, err := c.eventRepo.FindPoster(title, eventid)

	if err != nil {
		return nil, err
	}

	return &poster, nil
}

// PostersByEvent implements interfaces.EventUsecase
func (c *eventUsecase) PostersByEvent(eventid int) (*[]domain.PosterResponse, error) {
	fmt.Println("Poster by evnet called from usecase called")
	Posters, err := c.eventRepo.PostersByEvent(eventid)
	fmt.Println("posters:", Posters)
	if err != nil {
		fmt.Println("error from list organization from usecase:", err)
		return nil, err
	}

	return &Posters, nil
}

// FindUser implements interfaces.EventUsecase
func (c *eventUsecase) FindUser(username string) (bool, error) {
	vip, err := c.eventRepo.FindUser(username)
	if err != nil {
		return false, err
	}

	if vip == "false" {
		return false, nil
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
func (c *eventUsecase) AllApprovedEvents(pagenation utils.Filter, filter utils.FilterEvent) (*[]domain.EventResponse, *utils.Metadata, error) {
	fmt.Println("allevents from usecase called")
	events, metadata, err := c.eventRepo.AllApprovedEvents(pagenation, filter)
	fmt.Println("events:", events)
	if err != nil {
		fmt.Println("error from allevents usecase:", err)
		return nil, &metadata, err
	}

	return &events, &metadata, nil
}

// SearchEventUser implements interfaces.EventUsecase
func (c *eventUsecase) SearchEventUser(search string) (*[]domain.EventResponse, error) {
	fmt.Println("Search event from usecase called")
	SearchList, err := c.eventRepo.SearchEventUser(search)
	fmt.Println("searchList:", SearchList)
	if err != nil {
		fmt.Println("error from list organization from usecase:", err)
		return nil, err
	}

	return &SearchList, nil
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
