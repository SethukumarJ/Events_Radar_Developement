package usecase_test

import (
	"database/sql"
	"testing"



	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/mocks"
	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	// Create a new controller for managing mock objects
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock object for your user repository interface
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	// Create a mock user object
	user := domain.Users{
		UserName:    "testuser",
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "johndoe@example.com",
		PhoneNumber: "555-555-5555",
		Password:    "password",
		Profile:     "http://example.com/profile",
	}

	// Set the expectations for the mock object
	mockUserRepo.EXPECT().FindUserByEmail(user.Email).Return(nil, sql.ErrNoRows)
	mockUserRepo.EXPECT().InsertUser(user).Return(nil)

	// Create a new user usecase with the mock objects
	userUC := &userUseCase{
		userRepo:   mockUserRepo,
		adminRepo:  nil,
		mailConfig: nil,
		config:     nil,
	}

	// Call the function to be tested
	err := userUC.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}
