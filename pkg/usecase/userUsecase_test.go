package usecase

import (
	"errors"
	"testing"

	"github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	mock "github.com/SethukumarJ/Events_Radar_Developement/pkg/mock/repoMock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockUserRepository(ctrl)

	userusecase := NewUserUseCase(c, nil, nil, config.Config{})
	testData := []struct {
		name       string
		email      string
		beforeTest func(userRepo *mock.MockUserRepository)
		expectErr  error
	}{
		{
			name:  "Test success response",
			email: "jon",
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUserByName("jon").Return(domain.UserResponse{
					UserName: "jon",
					Password: "12345",
				}, nil)
			},
			expectErr: nil,
		},
		{
			name:  "Test when user alredy exist response",
			email: "jon",
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUserByName("jon").Return(domain.UserResponse{}, errors.New("repo error"))
			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			actualUser, err := userusecase.FindUserByName(tt.email)
			assert.Equal(t, tt.expectErr, err)
			if err == nil {
				assert.Equal(t, tt.email, actualUser.UserName)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockUserRepository(ctrl)

	userusecase := NewUserUseCase(c, nil, nil, config.Config{})
	testData := []struct {
		name       string
		user       domain.Users
		beforeTest func(userRepo *mock.MockUserRepository)
		expectErr  error
	}{
		{
			name: "Test success response",
			user: domain.Users{
				Email:    "jon",
				Password: "12345",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUserByName("jon").Return(domain.UserResponse{}, nil)

			},
			expectErr: errors.New("username already exists"),
		},
		{
			name: "Test Repo err response",
			user: domain.Users{
				Email:    "jon",
				Password: "12345",
			},
			beforeTest: func(userRepo *mock.MockUserRepository) {
				userRepo.EXPECT().FindUserByName("jon").Return(domain.UserResponse{}, errors.New("repo error"))

			},
			expectErr: errors.New("repo error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(c)
			err := userusecase.CreateUser(tt.user)
			assert.Equal(t, tt.expectErr, err)
		})
	}
}
