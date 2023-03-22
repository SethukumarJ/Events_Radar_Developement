package mocks

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/mocks/interfaces"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	m *mock.Mock
}


// FindUserByName implements interfaces.UserMockRepository
func (m *UserRepositoryMock) FindUserByName(email string) (domain.UserResponse, error) {
	args := m.m.Called(email)
	return *(args.Get(0).(*domain.UserResponse)), args.Error(1)
}

// InsertUser implements interfaces.UserMockRepository
func (m *UserRepositoryMock) InsertUser(user domain.Users) (int, error) {
	args := m.m.Called(user)
	return args.Int(0), args.Error(1)
}

func NewUserMockRepository(mock *mock.Mock) interfaces.UserMockRepository {
	return &UserRepositoryMock{
		m: mock,
	}
}

// UpdateProfile implements interfaces.UserMockRepository
func (*UserRepositoryMock) UpdateProfile(user domain.Bios, user_id int) (int, error) {
	panic("unimplemented")
}

// FindUserById implements interfaces.UserMockRepository
func (*UserRepositoryMock) FindUserById(user_id int) (domain.UserResponse, error) {
	panic("unimplemented")
}
