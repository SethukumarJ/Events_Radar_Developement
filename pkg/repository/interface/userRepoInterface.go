package interfaces

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (domain.UserResponse, error)
	FindOrganization(organizationName string) (domain.OrganizationsResponse, error)
	FindRole(username string, organizationName string) (string, error)
	UpdateProfile(user domain.Bios, username string) (int, error)
	UpdatePassword(user domain.Users, username string) (int, error)
	GetPublicFaqas(approved string) ([]domain.QAResponse, error)
	GetQuestions(title string) ([]domain.FaqaResponse, error)
	InsertUser(user domain.Users) (int, error)
	PostQuestion(question domain.Faqas) (int, error)
	PostAnswer(answer domain.Answers, question int) (int, error)
	StoreVerificationDetails(email string, code string) error
	VerifyAccount(email string, code string) error
	CreateOrganization(organization domain.Organizations) (int, error)
	ListOrganizations(pagenation utils.Filter) ([]domain.OrganizationsResponse, utils.Metadata, error)
	JoinOrganization(organizatinName string , username string) (int, error)
	AddMembers(newMembers []string, memberRole string, organizationName string) (int, error)
}
