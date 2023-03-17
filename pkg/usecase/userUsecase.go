package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	interfaces "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	services "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo   interfaces.UserRepository
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
}

// UpdateRole implements interfaces.UserUseCase
func (c *userUseCase) UpdateRole(user_id int, organization_id int, updatedRole string) error {
	_, err := c.userRepo.FindRelation(user_id, organization_id)

	if err != nil {
		return errors.New("relation does'nt exist with this credentials")

	}
	err = c.userRepo.UpdateRole(user_id, organization_id,updatedRole)

	if err != nil {
		return err
	}

	return nil
}

// DeleteMember implements interfaces.UserUseCase
func (c *userUseCase) DeleteMember(user_id int, organization_id int) error {

	_, err := c.userRepo.FindRelation(user_id, organization_id)

	if err != nil {
		return errors.New("relation does'nt exist with this credentials")

	}
	err = c.userRepo.DeleteMember(user_id, organization_id)

	if err != nil {
		return err
	}

	return nil
}

// ListMembers implements interfaces.UserUseCase
func (c *userUseCase) ListMembers(memberRole string, organization_id int) (*[]domain.UserOrganizationConnectionResponse, error) {
	fmt.Println("get membets  from usecase called")
	members, err := c.userRepo.ListMembers(memberRole, organization_id)
	fmt.Println("members:", members)
	if err != nil {
		fmt.Println("error from list members usecase:", err)
		return nil, err
	}

	return &members, nil
}

// Prmotion_Faliure implements interfaces.UserUseCase
func (c *userUseCase) Prmotion_Faliure(orderid string, paymentid string) error {
	err := c.userRepo.Prmotion_Faliure(orderid, paymentid)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Prmotion_Success implements interfaces.UserUseCase
func (c *userUseCase) Prmotion_Success(orderid string, paymentid string) error {
	err := c.userRepo.Prmotion_Success(orderid, paymentid)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// // PaymentFaliure implements interfaces.UserUseCase
// func (c *userUseCase) PaymentFaliure(orderid string) error {

// 	err := c.userRepo.PaymentFaliure(orderid)

// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }

// FeaturizeEvent implements interfaces.UserUseCase
func (c *userUseCase) FeaturizeEvent(orderid string) error {

	err := c.userRepo.FeaturizeEvent(orderid)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// PromoteEvent implements interfaces.UserUseCase
func (c *userUseCase) PromoteEvent(promotion domain.Promotion) error {
	err := c.userRepo.PromoteEvent(promotion)

	if err != nil {
		return err
	}
	return nil
}

// ApplyEvent implements interfaces.UserUseCase
func (c *userUseCase) ApplyEvent(applicationForm domain.ApplicationForm) error {
	fmt.Println("create organization from service")
	_, err := c.userRepo.FindApplication(applicationForm.UserId, applicationForm.EventId)
	fmt.Println("found applicationForm", err)

	if err == nil {
		return errors.New("applicationForm already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	_, err = c.userRepo.ApplyEvent(applicationForm)
	if err != nil {
		return err

	}
	return nil
}

// FindApplication implements interfaces.UserUseCase
func (c *userUseCase) FindApplication(user_id int, event_id int) (*domain.ApplicationFormResponse, error) {
	Application, err := c.userRepo.FindApplication(user_id, event_id)

	if err != nil {
		return nil, err
	}

	return &Application, nil
}

// AdmitMember implements interfaces.UserUseCase
func (c *userUseCase) AdmitMember(JoinStatusId int, memberRole string) error {

	user_id, organization_id, err := c.userRepo.FindJoinStatus(JoinStatusId)
	if err != nil {
		return errors.New("no join statuses found")
	}

	_, err = c.userRepo.FindRelation(user_id, organization_id)

	if err == nil {
		return errors.New("relation already exist with this credentials")

	}

	err = c.userRepo.AdmitMember(JoinStatusId, memberRole)

	if err != nil {
		return err
	}
	return nil
}

// ListJoinRequests implements interfaces.UserUseCase
func (c *userUseCase) ListJoinRequests(user_id int, organizaiton_id int) (*[]domain.Join_StatusResponse, error) {
	fmt.Println("get requests  from usecase called")
	requests, err := c.userRepo.ListJoinRequests(user_id, organizaiton_id)
	fmt.Println("requests:", requests)
	if err != nil {
		fmt.Println("error from listjoinRequests usecase:", err)
		return nil, err
	}

	return &requests, nil
}

// AcceptJoinInvitation implements interfaces.UserUseCase
func (c *userUseCase) AcceptJoinInvitation(user_id int, organizaiton_id int, role string) error {

	_, err := c.userRepo.FindRelation(user_id, organizaiton_id)

	if err == nil {
		return errors.New("relation allready exist with this credentials")

	}

	_, err = c.userRepo.AcceptJoinInvitation(user_id,  organizaiton_id,role)
	if err != nil {
		return err
	}
	return nil
}

// AddMembers implements interfaces.UserUseCase
func (c *userUseCase) AddMembers(newMembers []domain.AddMembers, memberRole string, organizaiton_id int) error {
	organizaition, err := c.userRepo.FindOrganizationById(organizaiton_id)
	fmt.Println("found organization", err)

	if err != nil {
		return err
	}

	for _, v := range newMembers {
		user, err := c.userRepo.FindUserByName(v.Members)

		if err == nil {
			c.SendInvitationMail(user.Email, organizaiton_id,organizaition.OrganizationName, memberRole)
		} else if err == sql.ErrNoRows {
			c.SendInvitationMail(v.Members, organizaiton_id,organizaition.OrganizationName, memberRole)
		} else {
			fmt.Println("coud'nt invite :", v)
		}
	}

	return nil
}

func (c *userUseCase) SendInvitationMail(email string, organization_id int,organizationName string, memberRole string) error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":         email,
		"organizationName": organizationName,
		"organizationId"  : organization_id,
		"memberRole":       memberRole,
		"exp":              time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	fmt.Println("TokenString", tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var role string
	if memberRole == "1" {
		role = "admin"
	} else if memberRole == "2" {
		role = "volunteer"
	} else if memberRole == "3" {
		role = "sponser"
	}

	subject := "Join invitation to organization " + organizationName + " for the role " + role
	message := []byte(
		"From: Events Radar <eventsRadarversion1@gmail.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			"<html>" +
			"  <head>" +
			"    <style>" +
			"      .blue-button {" +
			"        background-color: blue;" +
			"        color: white;" +
			"        padding: 10px 20px;" +
			"        border-radius: 5px;" +
			"        text-decoration: none;" +
			"        font-size: 16px;" +
			"      }" +
			"    </style>" +
			"  </head>" +
			"  <body>" +
			"    <p>Click the button on verify your accout:</p>" +
			"    <a class=\"blue-button\" href=\"https://eventsradar.online/user/accept-invitation?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
			"  </body>" +
			"</html>")

	// send random code to user's email
	if err := c.mailConfig.SendMail(c.config, email, []byte(message)); err != nil {
		return err
	}
	fmt.Println("email sent: ", email)

	err = c.userRepo.StoreVerificationDetails(email, tokenString)

	if err != nil {
		return err
	}

	return nil
}

func (c *userUseCase) VerifyRole(user_id int, organization_id int) (string, error) {

	role, err := c.userRepo.FindRole(user_id, organization_id)

	if err != nil {
		return "", err
	}

	return role, nil
}

// JoinOrganization implements interfaces.UserUseCase
func (c *userUseCase) JoinOrganization(organization_id int, user_id int) error {
	_, err := c.userRepo.JoinOrganization(organization_id, user_id)

	if err != nil {
		return err
	}
	return nil
}

// ListOrganizations implements interfaces.UserUseCase
func (c *userUseCase) ListOrganizations(pagenation utils.Filter) (*[]domain.OrganizationsResponse, *utils.Metadata, error) {
	fmt.Println("List Organization from usecase called")
	OrganizaionList, metadata, err := c.userRepo.ListOrganizations(pagenation)
	fmt.Println("organizations:", OrganizaionList)
	if err != nil {
		fmt.Println("error from list organization from usecase:", err)
		return nil, &metadata, err
	}

	return &OrganizaionList, &metadata, nil
}

// FindOrganization implements interfaces.UserUseCase
func (c *userUseCase) FindOrganizationByName(organizationName string) (*domain.OrganizationsResponse, error) {
	organization, err := c.userRepo.FindOrganizationByName(organizationName)

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// FindOrganization implements interfaces.UserUseCase
func (c *userUseCase) FindOrganizationById(organization_id int) (*domain.OrganizationsResponse, error) {
	organization, err := c.userRepo.FindOrganizationById(organization_id)

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// CreateOrganization implements interfaces.UserUseCase
func (c *userUseCase) CreateOrganization(organization domain.Organizations) error {
	fmt.Println("create organization from service")
	_, err := c.userRepo.FindOrganizationByName(organization.OrganizationName)
	fmt.Println("found organization", err)

	if err == nil {
		return errors.New("organization already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	_, err = c.userRepo.CreateOrganization(organization)
	if err != nil {
		return err

	}
	return nil
}

// GetQuestions implements interfaces.UserUseCase
func (c *userUseCase) GetQuestions(event_id int) (*[]domain.FaqaResponse, error) {
	fmt.Println("get questions  from usecase called")
	qustions, err := c.userRepo.GetQuestions(event_id)
	fmt.Println("questioins:", qustions)
	if err != nil {
		fmt.Println("error from getpublicfaqas usecase:", err)
		return nil, err
	}

	return &qustions, nil
}

// PostAnswer implements interfaces.UserUseCase
func (c *userUseCase) PostAnswer(answer domain.Answers, question int) error {
	_, err := c.userRepo.PostAnswer(answer, question)
	if err != nil {
		return err
	}
	return nil
}

// GetPublicFaqas implements interfaces.UserUseCase
func (c *userUseCase) GetPublicFaqas(event_id int) (*[]domain.QAResponse, error) {
	fmt.Println("get faqas  from usecase called")
	faqas, err := c.userRepo.GetPublicFaqas(event_id)
	fmt.Println("faqas:", faqas)
	if err != nil {
		fmt.Println("error from getpublicfaqas usecase:", err)
		return nil, err
	}

	return &faqas, nil
}

// PostQuestion implements interfaces.UserUseCase
func (c *userUseCase) PostQuestion(question domain.Faqas) error {

	_, err := c.userRepo.PostQuestion(question)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePassword implements interfaces.UserUseCase
func (c *userUseCase) UpdatePassword(password string, email string) error {
	_, err := c.userRepo.UpdatePassword(password, email)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProfile implements interfaces.UserUseCase
func (c *userUseCase) UpdateProfile(user domain.Bios, user_id int) error {
	fmt.Println("update user from service")

	_, err := c.userRepo.UpdateProfile(user, user_id)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser implements interfaces.UserUseCase
func (c *userUseCase) CreateUser(user domain.Users) error {
	fmt.Println("create user from service")
	_, err := c.userRepo.FindUserByName(user.Email)
	fmt.Println("found user", err)

	if err == nil {
		return errors.New("username already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	user.Password = HashPassword(user.Password)
	fmt.Println("password", user.Password)
	_, err = c.userRepo.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

// FindUser implements interfaces.UserUseCase
func (c *userUseCase) FindUserByName(email string) (*domain.UserResponse, error) {
	user, err := c.userRepo.FindUserByName(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUser implements interfaces.UserUseCase
func (c *userUseCase) FindUserById(user_id int) (*domain.UserResponse, error) {
	user, err := c.userRepo.FindUserById(user_id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *userUseCase) SendVerificationEmail(email string) error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	subject := "Account Verification"
	msg := []byte(
		"From: Events Radar <eventsRadarversion1@gmail.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			"<html>" +
			"  <head>" +
			"    <style>" +
			"      .blue-button {" +
			"        background-color: blue;" +
			"        color: white;" +
			"        padding: 10px 20px;" +
			"        border-radius: 5px;" +
			"        text-decoration: none;" +
			"        font-size: 16px;" +
			"      }" +
			"    </style>" +
			"  </head>" +
			"  <body>" +
			"    <p>Click the button on verify your accout:</p>" +
			"    <a class=\"blue-button\" href=\"https://eventsradar.online/user/verify-account?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
			"  </body>" +
			"</html>")

	// send email with HTML message
	if err := c.mailConfig.SendMail(c.config, email, msg); err != nil {
		return err
	}

	err = c.userRepo.StoreVerificationDetails(email, tokenString)
	if err != nil {
		return err
	}

	return nil
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func NewUserUseCase(
	userRepo interfaces.UserRepository,
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig,
	config config.Config) services.UserUseCase {
	return &userUseCase{
		userRepo:   userRepo,
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
		config:     config,
	}
}
