package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	PostAnswer(answer domain.Answers, question int) error
	CreateOrganization(organization domain.Organizations) error
	JoinOrganization(organizationName string, userName string) error
	VerifyRole(username string, organizationName string) (string, error)
	FindOrganization(organizationName string) (*domain.OrganizationsResponse, error)
	ListOrganizations(pagenation utils.Filter) (*[]domain.OrganizationsResponse, *utils.Metadata, error)
	UpdatePassword(user domain.Users, email string) error
	FindUser(email string) (*domain.UserResponse, error)
	SendVerificationEmail(email string) error
	PostQuestion(question domain.Faqas) error
	GetPublicFaqas(title string) (*[]domain.QAResponse, error)
	ListJoinRequests(username string,organizationName string) (*[]domain.Join_StatusResponse, error)
	GetQuestions(title string) (*[]domain.FaqaResponse, error)
	UpdateProfile(user domain.Bios, username string) error
	AddMembers(newMembers []string,memberRole string, organizationName string) error
	AcceptJoinInvitation(username string, organizationName string,role string) error
	AdmitMember(JoinStatusId int) error
}
