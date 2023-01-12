package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo   interfaces.UserRepository
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
}

// PostAnswer implements interfaces.UserUseCase
func (c *userUseCase) PostAnswer(answer domain.Answers, question int) error {
	_, err := c.userRepo.PostAnswer(answer,question)
	if err != nil {
		return err
	}
	return nil
}

// GetPublicFaqas implements interfaces.UserUseCase
func (c *userUseCase) GetPublicFaqas(title string) (*[]domain.FaqaResponse, error) {
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

// SendVerificationEmail implements interfaces.UserUseCase
func (c *userUseCase) SendVerificationEmail(email string) error {
	//to generate random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)

	fmt.Println("code: ", code)

	message := fmt.Sprintf(
		"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for usingEvents.\n with regards Team Events radar.",
		code,
	)

	// send random code to user's email
	if err := c.mailConfig.SendMail(c.config, email, message); err != nil {
		return err
	}
	fmt.Println("email sent: ", email)

	err := c.userRepo.StoreVerificationDetails(email, code)

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
