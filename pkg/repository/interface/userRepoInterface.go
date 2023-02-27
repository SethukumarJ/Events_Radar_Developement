package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (domain.UserResponse, error)
	FindOrganization(organizationName string) (domain.OrganizationsResponse, error)
	FindApplication(username string) (domain.ApplicationFormResponse, error)
	FindRole(username string, organizationName string) (string, error)
	FindRelation(username string, organizationName string) (string, error)
	UpdateProfile(user domain.Bios, username string) (int, error)
	UpdatePassword(user string, username string) (int, error)
	GetPublicFaqas(approved string) ([]domain.QAResponse, error)
	ListJoinRequests(username string,organizationName string) ([]domain.Join_StatusResponse, error)
	GetQuestions(title string) ([]domain.FaqaResponse, error)
	InsertUser(user domain.Users) (int, error)
	PostQuestion(question domain.Faqas) (int, error)
	PostAnswer(answer domain.Answers, question int) (int, error)
	StoreVerificationDetails(email string, code string) error
	VerifyAccount(email string, code string) error
	CreateOrganization(organization domain.Organizations) (int, error)
	ApplyEvent(applicationForm domain.ApplicationForm) (int , error)
	ListOrganizations(pagenation utils.Filter) ([]domain.OrganizationsResponse, utils.Metadata, error)
	JoinOrganization(organizatinName string , username string) (int, error)
	AcceptJoinInvitation(newMember string, memberRole string, organizationName string) (int, error)
	AdmitMember(JoinStatusId int,memberRole string) error
	FindJoinStatus(JoinStatusId int) (string ,string,error)
	PromoteEvent(promotion domain.Promotion) error
	FeaturizeEvent(orderid string) error
	
}
