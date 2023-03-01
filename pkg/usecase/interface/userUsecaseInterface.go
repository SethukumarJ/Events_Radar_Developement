package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	PostAnswer(answer domain.Answers, question int) error
	CreateOrganization(organization domain.Organizations) error
	ApplyEvent(applicationForm domain.ApplicationForm) error
	FindApplication(userName string,eventname string) (*domain.ApplicationFormResponse, error)
	JoinOrganization(organizationName string, userName string) error
	VerifyRole(username string, organizationName string) (string, error)
	FindOrganization(organizationName string) (*domain.OrganizationsResponse, error)
	ListOrganizations(pagenation utils.Filter) (*[]domain.OrganizationsResponse, *utils.Metadata, error)
	UpdatePassword(user string, email string) error
	FindUser(email string) (*domain.UserResponse, error)
	SendVerificationEmail(email string) error
	PostQuestion(question domain.Faqas) error
	GetPublicFaqas(title string) (*[]domain.QAResponse, error)
	ListJoinRequests(username string,organizationName string) (*[]domain.Join_StatusResponse, error)
	GetQuestions(title string) (*[]domain.FaqaResponse, error)
	UpdateProfile(user domain.Bios, username string) error
	AddMembers(newMembers []string,memberRole string, organizationName string) error
	AcceptJoinInvitation(username string, organizationName string,role string) error
	AdmitMember(JoinStatusId int , memberRole string) error
    PromoteEvent(promotion domain.Promotion) error
	FeaturizeEvent(orderid string) error
	Prmotion_Success(orderid string,paymentid string) error
	Prmotion_Faliure(orderid string,paymentid string) error
	// PaymentFaliure(orderid string) error


}
