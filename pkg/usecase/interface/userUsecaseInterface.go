package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type UserUseCase interface {
	CreateUser(user domain.Users) error
	FindUserByName(email string) (*domain.UserResponse, error)
	FindUserById(user_id int) (*domain.UserResponse, error)
	PostAnswer(answer domain.Answers, question int) error
	ApplyEvent(applicationForm domain.ApplicationForm) error
	FindApplication(user_id int,event_id int) (*domain.ApplicationFormResponse, error)
	JoinOrganization(organization_id int, user_id int) error
	VerifyRole(user_id int, organization_id int) (string, error)
	CreateOrganization(organization domain.Organizations) error
	FindOrganizationById(organization_id int) (*domain.OrganizationsResponse, error)
	FindOrganizationByName(organizationName string) (*domain.OrganizationsResponse, error)
	ListOrganizations(pagenation utils.Filter) (*[]domain.OrganizationsResponse, *utils.Metadata, error)
	UpdatePassword(password string, email string) error
	SendVerificationEmail(email string) error
	PostQuestion(question domain.Faqas) error
	GetPublicFaqas(event_id int) (*[]domain.QAResponse, error)
	ListJoinRequests(user_id int,organization_id int) (*[]domain.Join_StatusResponse, error)
	GetQuestions(event_id int) (*[]domain.FaqaResponse, error)
	UpdateProfile(user domain.Bios, user_id int) error
	AddMembers(newMembers []domain.AddMembers,memberRole string, organization_id int) error
	AcceptJoinInvitation(user_id int, organization_id int,role string) error
	AdmitMember(JoinStatusId int , memberRole string) error
    PromoteEvent(promotion domain.Promotion) error
	FeaturizeEvent(order_id string) error
	Prmotion_Success(order_id string,payment_id string) error
	Prmotion_Faliure(order_id string,payment_id string) error
	ListMembers(memberRole string, organization_id int) (*[]domain.UserOrganizationConnectionResponse, error)
	DeleteMember(user_id int,organization_id int) error
	UpdateRole(user_id int,organization_id int,updatedRole string) error
 
	// PaymentFaliure(orderid string) error


}
