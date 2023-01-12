package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	PostAnswer(answer domain.Answers, question int) error
	UpdatePassword(user domain.Users,email string) error
	FindUser(email string) (*domain.UserResponse, error)
	SendVerificationEmail(email string) (error)
	PostQuestion(question domain.Faqas) error
	GetPublicFaqas(title string) (*[]domain.QAResponse, error) 
	UpdateProfile(user domain.Bios, username string) error
}
