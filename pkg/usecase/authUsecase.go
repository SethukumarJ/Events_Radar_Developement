package usecase

import (
	"errors"
	"fmt"
	"log"

	repository "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

// authUsecase is the struct for the authentication service
type authUsecase struct {
	userRepo  repository.UserRepository
	adminRepo repository.AdminRepository
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	adminRepo repository.AdminRepository,
) usecase.AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		adminRepo: adminRepo,
	}
}

// VerifyAccount implements interfaces.UserUseCase
func (c *authUsecase) VerifyAccount(email string, code string) error {
	err := c.userRepo.VerifyAccount(email, code)

	if err != nil {
		return err
	}
	return nil
}

// VerifyUser verifies the user credentials
func (c *authUsecase) VerifyUser(email string, password string) error {

	user, err := c.userRepo.FindUserByName(email)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	if !user.Verification {
		return errors.New("failed to login. user not verified")
	}

	return nil
}

// VerifyUser verifies the user credentials
func (c *authUsecase) VerifyAdmin(email string, password string) error {

	admin, err := c.adminRepo.FindAdminByName(email)

	fmt.Println("admin.Password", admin.Password)
	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(admin.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

func VerifyPassword(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
