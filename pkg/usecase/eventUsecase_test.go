package usecase

import (
	"errors"
	"testing"

	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	mock "github.com/SethukumarJ/Events_Radar_Developement/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockEventRepository(ctrl)

	eventUsecase := NewEventUseCase(c)
	testData := []struct {
		name       string
		event      domain.Events
		beforeTest func(eventRepo *mock.MockEventRepository)
		expectErr  error
	}{
		{
			name: "Test success response",
			event: domain.Events{

				Title: "event1",
			},
			beforeTest: func(userRepo *mock.MockEventRepository) {
				userRepo.EXPECT().FindEventByTitle("event1").Return(domain.EventResponse{}, nil)

			},
			expectErr: errors.New("event title already exists"),
		},
		{
			name: "Test Repo err response",
			event: domain.Events{
				Title: "event1",
			},
			beforeTest: func(eventRepo *mock.MockEventRepository) {
				eventRepo.EXPECT().FindEventByTitle("event1").Return(domain.EventResponse{}, errors.New("repo error"))

			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			err := eventUsecase.CreateEvent(tt.event)
			assert.Equal(t, tt.expectErr, err)
		})
	}
}


func TestFindEventByTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockEventRepository(ctrl)

	eventUsecase := NewEventUseCase(c)
	testData := []struct {
		name       string
		title      string
		beforeTest func(eventRepo *mock.MockEventRepository)
		expectErr  error
	}{
		{
			name:  "Test success response",
			title: "event1",
			beforeTest: func(eventRepo *mock.MockEventRepository) {
				eventRepo.EXPECT().FindEventByTitle("event1").Return(domain.EventResponse{
					Title: "event1",
				}, nil)
			},
			expectErr: nil,
		},
		{
			name:  "Test when user alredy exist response",
			title: "event1",
			beforeTest: func(eventRepo *mock.MockEventRepository) {
				eventRepo.EXPECT().FindEventByTitle("event1").Return(domain.EventResponse{}, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			actualEvent, err := eventUsecase.FindEventByTitle(tt.title)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.title, actualEvent.Title)
			}
		})
	}
}