package repository_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	repository "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository"
)

func TestCreateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewEventRepository(db)

	// Test Case 1: Successful insert
	event := domain.Events{
		Title:                  "Example Event",
		OrganizationId:         123,
		User_id:                456,
		CreatedBy:              "John Doe",
		EventPic:               "https://example.com/event.jpg",
		ShortDiscription:       "A short description of the event",
		LongDiscription:        "A longer description of the event",
		EventDate:              "2023-04-01",
		Location:               "Example Venue",
		CreatedAt:              time.Now(),
		Approved:               true,
		Paid:                   true,
		Amount:                 "10.00",
		Sex:                    "male",
		CusatOnly:              false,
		Archived:               false,
		SubEvents:              "",
		Online:                 true,
		MaxApplications:        101,
		ApplicationClosingDate: "2023-03-31",
		ApplicationLink:        "https://example.com/apply",
		WebsiteLink:            "https://example.com",
		ApplicationLeft:        101,
		Featured:               true,
	}

	mock.ExpectQuery(`INSERT INTO events\(title,organization_id,user_id,created_by,event_pic,short_discription,long_discription,event_date,location,created_at,approved,paid,amount,sex,cusat_only,archived,sub_events,online,max_applications,application_closing_date,application_link,website_link,application_left\)VALUES\(\$1, \$2, \$3, \$4, \$5, \$6,\$7,\$8, \$9, \$10, \$11, \$12, \$13,\$14,\$15, \$16, \$17, \$18, \$19,\$20,\$21,\$22,\$23\)RETURNING event_id;`).
		WithArgs(event.Title,
			event.OrganizationId,
			event.User_id,
			event.CreatedBy,
			event.EventPic,
			event.ShortDiscription,
			event.LongDiscription,
			event.EventDate,
			event.Location,
			event.CreatedAt,
			event.Approved,
			event.Paid,
			event.Amount,
			event.Sex,
			event.CusatOnly,
			event.Archived,
			event.SubEvents,
			event.Online,
			event.MaxApplications,
			event.ApplicationClosingDate,
			event.ApplicationLink,
			event.WebsiteLink, event.ApplicationLeft).
		WillReturnRows(sqlmock.NewRows([]string{"event_id"}).AddRow(1))

	eventId, err := repo.CreateEvent(event)
	assert.NoError(t, err)
	assert.Equal(t, 1, eventId)

	// Test Case 2: Duplicate username
	mock.ExpectQuery(`INSERT INTO events\(title,organization_id,user_id,created_by,event_pic,short_discription,long_discription,event_date,location,created_at,approved,paid,amount,sex,cusat_only,archived,sub_events,online,max_applications,application_closing_date,application_link,website_link,application_left\)VALUES\(\$1, \$2, \$3, \$4, \$5, \$6,\$7,\$8, \$9, \$10, \$11, \$12, \$13,\$14,\$15, \$16, \$17, \$18, \$19,\$20,\$21,\$22,\$23\)RETURNING event_id;`).
		WithArgs(event.Title,
			event.OrganizationId,
			event.User_id,
			event.CreatedBy,
			event.EventPic,
			event.ShortDiscription,
			event.LongDiscription,
			event.EventDate,
			event.Location,
			event.CreatedAt,
			event.Approved,
			event.Paid,
			event.Amount,
			event.Sex,
			event.CusatOnly,
			event.Archived,
			event.SubEvents,
			event.Online,
			event.MaxApplications,
			event.ApplicationClosingDate,
			event.ApplicationLink,
			event.WebsiteLink, event.ApplicationLeft).
		WillReturnError(errors.New("duplicate key value violates unique constraint"))

	eventId, err = repo.CreateEvent(event)
	assert.Error(t, err)
	assert.Equal(t, 0, eventId)

}

func TestFindPosterById(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %s", err)
	}
	defer db.Close()

	// Create a new userRepository instance with the mock database connection
	eventRepo := repository.NewEventRepository(db)

	// Define a test case
	testCase := struct {
		posterID       int
		eventID        int
		expectedPoster domain.PosterResponse
		expectedErr    error
	}{
		posterID: 1,
		eventID:  1,
		expectedPoster: domain.PosterResponse{
			PosterId:    1,
			Name:        "Sample Poster",
			Image:       "https://example.com/poster.jpg",
			Discription: "This is a sample poster",
			Date:        "2022-01-01",
			Colour:      "red",
			EventId:     1,
		},
		expectedErr: nil,
	}

	// Define the expected SQL query and result

	mock.ExpectQuery(`SELECT poster_id,name,image,discription,date,colour,event_id FROM posters WHERE poster_id  = \$1 AND event_id = \$2;`).WithArgs(testCase.posterID, testCase.eventID).WillReturnRows(
		sqlmock.NewRows([]string{"poster_id", "name", "image", "discription", "date", "colour", "event_id"}).
			AddRow(testCase.expectedPoster.PosterId, testCase.expectedPoster.Name, testCase.expectedPoster.Image, testCase.expectedPoster.Discription, testCase.expectedPoster.Date, testCase.expectedPoster.Colour, testCase.expectedPoster.EventId))

	// Call the FindUserById method with the test case user ID
	poster, err := eventRepo.FindPosterById(testCase.posterID, testCase.eventID)

	// Check the result and error against the expected values
	assert.Equal(t, testCase.expectedErr, err)
	assert.Equal(t, testCase.expectedPoster, poster)
}

func TestFindPosterByName(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %s", err)
	}
	defer db.Close()

	// Create a new userRepository instance with the mock database connection
	eventRepo := repository.NewEventRepository(db)

	// Define a test case
	testCase := struct {
		posterName     string
		eventID        int
		expectedPoster domain.PosterResponse
		expectedErr    error
	}{
		posterName: "Sample Poster",
		eventID:    1,
		expectedPoster: domain.PosterResponse{
			PosterId:    1,
			Name:        "Sample Poster",
			Image:       "https://example.com/poster.jpg",
			Discription: "This is a sample poster",
			Date:        "2022-01-01",
			Colour:      "red",
			EventId:     1,
		},
		expectedErr: nil,
	}

	// Define the expected SQL query and result

	mock.ExpectQuery(`SELECT poster_id,name,image,discription,date,colour,event_id FROM posters WHERE name  = \$1 AND event_id = \$2;`).WithArgs(testCase.posterName, testCase.eventID).WillReturnRows(
		sqlmock.NewRows([]string{"poster_id", "name", "image", "discription", "date", "colour", "event_id"}).
			AddRow(testCase.expectedPoster.PosterId, testCase.expectedPoster.Name, testCase.expectedPoster.Image, testCase.expectedPoster.Discription, testCase.expectedPoster.Date, testCase.expectedPoster.Colour, testCase.expectedPoster.EventId))

	// Call the FindUserById method with the test case user ID
	poster, err := eventRepo.FindPosterByName(testCase.posterName, testCase.eventID)

	// Check the result and error against the expected values
	assert.Equal(t, testCase.expectedErr, err)
	assert.Equal(t, testCase.expectedPoster, poster)
}


func TestEventRepository_PostersByEvent(t *testing.T) {
	// Create a new mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %s", err)
	}
	defer db.Close()

	// Create a new event repository using the mock DB
	eventRepo := repository.NewEventRepository(db)

	// Create test data
	eventID := 1
	expectedPosters := []domain.PosterResponse{
		{
			PosterId:     1,
			Name:         "Poster 1",
			Image:        "image1.jpg",
			Discription:  "Description 1",
			Date:         "2022-01-01",
			Colour:       "Red",
			EventId:      1,
		},
		{
			PosterId:     2,
			Name:         "Poster 2",
			Image:        "image2.jpg",
			Discription:  "Description 2",
			Date:         "2022-01-02",
			Colour:       "Green",
			EventId:      1,
		},
	}
	rows := sqlmock.NewRows([]string{"count", "poster_id", "name", "image", "discription", "date", "colour", "event_id"})
	for _, poster := range expectedPosters {
		rows = rows.AddRow(len(expectedPosters), poster.PosterId, poster.Name, poster.Image, poster.Discription, poster.Date, poster.Colour, poster.EventId)
	}

	// Set expectations on the mock DB
	mock.ExpectQuery("^SELECT COUNT(.+) FROM posters WHERE event_id = (.+)$").WillReturnRows(rows)
	fmt.Println("rows",rows)

	// Call the function being tested
	actualPosters, err := eventRepo.PostersByEvent(eventID)
	if err != nil {
		t.Fatalf("error calling PostersByEvent: %s", err)
	}

	// Check the results
	assert.Equal(t, expectedPosters, actualPosters)
}

