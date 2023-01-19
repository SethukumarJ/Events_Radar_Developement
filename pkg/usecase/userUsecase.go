package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo   interfaces.UserRepository
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
}

// AddMembers implements interfaces.UserUseCase
func (c *userUseCase) AddMembers(newMembers []string,memberRole string, organizationName string) error {
	_, err := c.userRepo.FindOrganization(organizationName)
	fmt.Println("found organization", err)

	if err != nil{
		return err
	}

	_ ,err = c.userRepo.AddMembers(newMembers, memberRole,organizationName)

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
func (c *userUseCase) UpdatePassword(user domain.Users, email string) error {
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
	fmt.Println("TokenString", tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	subject := "Account Verification"
	body := "Please click on the link to verify your account: http://localhost:3000/user/verify/account?token=" + tokenString
	message := "To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body

	// send random code to user's email
	if err := c.mailConfig.SendMail(c.config, email, message); err != nil {
		return err
	}
	fmt.Println("email sent: ", email)

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
