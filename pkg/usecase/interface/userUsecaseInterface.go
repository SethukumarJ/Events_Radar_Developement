package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	UpdatePassword(user domain.Users,email string) error
	FindUser(email string) (*domain.UserResponse, error)
	SendVerificationEmail(email string) (error)
	PostQuestion(question domain.Faqa) error
	UpdateProfile(user domain.Bios, username string) error
}
