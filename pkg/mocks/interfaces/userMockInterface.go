package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
)

type UserMockRepository interface {
	FindUserByName(email string) (domain.UserResponse, error)
	FindUserById(user_id int) (domain.UserResponse, error)
	InsertUser(user domain.Users) (int, error)
	UpdateProfile(user domain.Bios, user_id int) (int, error)
}
