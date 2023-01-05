package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// CreateUser implements interfaces.UserUseCase
func (c *userUseCase) CreateUser(user domain.Users) error {
	fmt.Println("create user from service")
	_, err := c.userRepo.FindUser(user.Email)
	fmt.Println("found user", err)

	if err == nil {
		return errors.New("Username already exists")
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
func (*userUseCase) FindUser(email string) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// SendVerificationEmail implements interfaces.UserUseCase
func (*userUseCase) SendVerificationEmail(email string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.UserUseCase
func (*userUseCase) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
