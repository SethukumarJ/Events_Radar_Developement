package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUserByName(email string) (domain.UserResponse, error)
	FindUserById(user_id int) (domain.UserResponse, error)
	FindOrganizationByName(organizationName string) (domain.OrganizationsResponse, error)
	FindOrganizationById(organizaiton_id int) (domain.OrganizationsResponse, error)
	FindApplication(user_id int,event_id int) (domain.ApplicationFormResponse, error)
	FindRole(user_id int, organizaiton_id int) (string, error)
	FindRelation(user_id int, organizaiton_id int) (string, error)
	UpdateProfile(user domain.Bios, user_id int) (int, error)
	UpdatePassword(password string, username string) (int, error)
	GetPublicFaqas(event_id int) ([]domain.QAResponse, error)
	ListJoinRequests(user_id int,organizaiton_id int) ([]domain.Join_StatusResponse, error)
	GetQuestions(evnent_id int) ([]domain.FaqaResponse, error)
	InsertUser(user domain.Users) (int, error)
	PostQuestion(question domain.Faqas) (int, error)
	PostAnswer(answer domain.Answers, question int) (int, error)
	StoreVerificationDetails(email string, code string) error
	VerifyAccount(email string, code string) error
	CreateOrganization(organization domain.Organizations) (int, error)
	ApplyEvent(applicationForm domain.ApplicationForm) (int , error)
	ListOrganizations(pagenation utils.Filter) ([]domain.OrganizationsResponse, utils.Metadata, error)
	JoinOrganization(organization_id int , user_id int) (int, error)
	AcceptJoinInvitation(user_id int,organizaiton_id int, memberRole string) (int, error)
	AdmitMember(JoinStatusId int,memberRole string) error
	FindJoinStatus(JoinStatusId int) (int ,int,error)
	PromoteEvent(promotion domain.Promotion) error
	FeaturizeEvent(orderid string) error
	Prmotion_Success(orderid string,paymentid string) error
	Prmotion_Faliure(orderid string,paymentid string) error
	ListMembers(memberRole string, organizaiton_id int) ([]domain.UserOrganizationConnectionResponse,error)
	DeleteMember(user_id int, organizaiton_id int) error
	UpdateRole(user_id int, organizaiton_id int, updatedRole string) error 
	
}
