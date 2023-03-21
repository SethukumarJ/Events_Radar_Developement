package repository_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	repository "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	// Test Case 1: Successful insert
	user := domain.Users{
		UserName:    "testuser",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "johndoe@example.com",
		PhoneNumber: "555-555-5555",
		Password:    "password",
		Profile:     "http://example.com/profile",
	}

	mock.ExpectQuery(`INSERT INTO users\(user_name,first_name,last_name,email,phone_number,password,profile\)VALUES\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)RETURNING user_id`).
		WithArgs(user.UserName, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password, user.Profile).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

	mock.ExpectQuery(`INSERT INTO bios\(user_id\)VALUES\(\$1\)RETURNING bio_id`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"bio_id"}).AddRow(1))

	userID, err := repo.InsertUser(user)
	assert.NoError(t, err)
	assert.Equal(t, 1, userID)

	// Test Case 2: Duplicate username
	mock.ExpectQuery(`INSERT INTO users\(user_name,first_name,last_name,email,phone_number,password,profile\)VALUES\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)RETURNING user_id`).
		WithArgs(user.UserName, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password, user.Profile).
		WillReturnError(errors.New("duplicate key value violates unique constraint"))

	userID, err = repo.InsertUser(user)
	assert.Error(t, err)
	assert.Equal(t, 0, userID)

	// Test Case 1: Successful insert
	user2 := domain.Users{
		UserName:    "testuser2",
		FirstName:   "John2",
		LastName:    "Doe2",
		Email:       "johndoe2@example.com",
		PhoneNumber: "555-555-55552",
		Password:    "password2",
		Profile:     "http://example.com/profile2",
	}

	// Test Case 2: Duplicate username
	mock.ExpectQuery(`INSERT INTO users\(user_name,first_name,last_name,email,phone_number,password,profile\)VALUES\(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)RETURNING user_id`).
		WithArgs(user2.UserName, user2.FirstName, user2.LastName, user2.Email, user2.PhoneNumber, user2.Password, user2.Profile).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

	userID, err = repo.InsertUser(user2)
	assert.Error(t, err)
	assert.Equal(t, 1, userID)

}


func TestFindUserById(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %s", err)
	}
	defer db.Close()

	// Create a new userRepository instance with the mock database connection
	userRepo := repository.NewUserRepository(db)

	// Define a test case
	testCase := struct {
		userID       int
		expectedUser domain.UserResponse
		expectedErr  error
	}{
		userID: 1,
		expectedUser: domain.UserResponse{
			UserId:       1,
			UserName:     "john.doe",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john.doe@example.com",
			Password:     "passw0rd",
			Verification: true,
			Vip:          false,
			PhoneNumber:  "1234567890",
			Profile:      "https://example.com/john.doe",
		},
		expectedErr: nil,
	}

	
	// Define the expected SQL query and result
	
	mock.ExpectQuery(`SELECT user_id,user_name,firs_name,last_name,email,password,phone_number,profile,verification FROM users WHERE user_id = \$1;`).WithArgs(testCase.userID).WillReturnRows(
		sqlmock.NewRows([]string{"user_id", "user_name", "first_name", "last_name", "email", "password", "phone_number", "profile", "verification"}).
			AddRow(testCase.expectedUser.UserId, testCase.expectedUser.UserName, testCase.expectedUser.FirstName, testCase.expectedUser.LastName, testCase.expectedUser.Email, testCase.expectedUser.Password, testCase.expectedUser.PhoneNumber, testCase.expectedUser.Profile, testCase.expectedUser.Verification))
	

	// Call the FindUserById method with the test case user ID
	user, err := userRepo.FindUserById(testCase.userID)

	// Check the result and error against the expected values
	assert.Equal(t, testCase.expectedErr, err)
	assert.Equal(t, testCase.expectedUser, user)
}

