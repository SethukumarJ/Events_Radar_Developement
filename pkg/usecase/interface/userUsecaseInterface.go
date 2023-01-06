package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	FindUser(email string) (*domain.UserResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
}
