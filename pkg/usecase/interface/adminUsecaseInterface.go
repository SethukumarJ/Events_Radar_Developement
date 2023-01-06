package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type AdminUsecase interface {
	CreateAdmin(admin domain.Admins) error
	FindAdmin(email string) (*domain.AdminResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
}