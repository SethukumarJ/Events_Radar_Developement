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
	_, err := c.userRepo.FindApplication(applicationForm.UserName)
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
func (c *userUseCase) FindApplication(userName string) (*domain.ApplicationFormResponse, error) {
	Application, err := c.userRepo.FindApplication(userName)

	if err != nil {
		return nil, err
	}

	return &Application, nil
}

// AdmitMember implements interfaces.UserUseCase
func (c *userUseCase) AdmitMember(JoinStatusId int, memberRole string) error {

	userName, organizationName, err := c.userRepo.FindJoinStatus(JoinStatusId)
	if err != nil {
		return errors.New("no join statuses found")
	}

	_, err = c.userRepo.FindRelation(userName, organizationName)

	if err == nil {
		return errors.New("relation allready exist with this credentials")

	}

	err = c.userRepo.AdmitMember(JoinStatusId, memberRole)

	if err != nil {
		return err
	}
	return nil
}

// ListJoinRequests implements interfaces.UserUseCase
func (c *userUseCase) ListJoinRequests(username string, organizationName string) (*[]domain.Join_StatusResponse, error) {
	fmt.Println("get requests  from usecase called")
	requests, err := c.userRepo.ListJoinRequests(username, organizationName)
	fmt.Println("requests:", requests)
	if err != nil {
		fmt.Println("error from listjoinRequests usecase:", err)
		return nil, err
	}

	return &requests, nil
}

// AcceptJoinInvitation implements interfaces.UserUseCase
func (c *userUseCase) AcceptJoinInvitation(username string, organizationName string, role string) error {

	_, err := c.userRepo.FindRelation(username, organizationName)

	if err == nil {
		return errors.New("relation allready exist with this credentials")

	}

	_, err = c.userRepo.AcceptJoinInvitation(username, organizationName, role)
	if err != nil {
		return err
	}
	return nil
}

// AddMembers implements interfaces.UserUseCase
func (c *userUseCase) AddMembers(newMembers []string, memberRole string, organizationName string) error {
	_, err := c.userRepo.FindOrganization(organizationName)
	fmt.Println("found organization", err)

	if err != nil {
		return err
	}

	for _, v := range newMembers {
		user, err := c.userRepo.FindUser(v)

		if err == nil {
			c.SendInvitationMail(user.Email, organizationName, memberRole)
		} else if err == sql.ErrNoRows {
			c.SendInvitationMail(v, organizationName, memberRole)
		} else {
			fmt.Println("coud'nt invite :", v)
		}
	}

	return nil
}

func (c *userUseCase) SendInvitationMail(email string, organizationName string, memberRole string) error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":         email,
		"organizationName": organizationName,
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
			"    <a class=\"blue-button\" href=\"http://localhost:3000/accept-invitation?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
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

func (c *userUseCase) VerifyRole(username string, organizationName string) (string, error) {

	role, err := c.userRepo.FindRole(username, organizationName)

	if err != nil {
		return "", err
	}

	return role, nil
}

// JoinOrganization implements interfaces.UserUseCase
func (c *userUseCase) JoinOrganization(organizationName string, userName string) error {
	_, err := c.userRepo.JoinOrganization(organizationName, userName)

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
func (c *userUseCase) FindOrganization(organizationName string) (*domain.OrganizationsResponse, error) {
	organization, err := c.userRepo.FindOrganization(organizationName)

	if err != nil {
		return nil, err
	}

	return &organization, nil
}

// CreateOrganization implements interfaces.UserUseCase
func (c *userUseCase) CreateOrganization(organization domain.Organizations) error {
	fmt.Println("create organization from service")
	_, err := c.userRepo.FindOrganization(organization.OrganizationName)
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
func (c *userUseCase) GetQuestions(title string) (*[]domain.FaqaResponse, error) {
	fmt.Println("get questions  from usecase called")
	qustions, err := c.userRepo.GetQuestions(title)
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
func (c *userUseCase) GetPublicFaqas(title string) (*[]domain.QAResponse, error) {
	fmt.Println("get faqas  from usecase called")
	faqas, err := c.userRepo.GetPublicFaqas(title)
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
func (c *userUseCase) UpdatePassword(user string, email string) error {
	_, err := c.userRepo.UpdatePassword(user, email)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProfile implements interfaces.UserUseCase
func (c *userUseCase) UpdateProfile(user domain.Bios, username string) error {
	fmt.Println("update user from service")

	_, err := c.userRepo.UpdateProfile(user, username)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser implements interfaces.UserUseCase
func (c *userUseCase) CreateUser(user domain.Users) error {
	fmt.Println("create user from service")
	_, err := c.userRepo.FindUser(user.Email)
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
func (c *userUseCase) FindUser(email string) (*domain.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)

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
			"    <a class=\"blue-button\" href=\"http://localhost:3000/user/verify-account?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
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
