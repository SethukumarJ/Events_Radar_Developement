package usecase

import (
	"errors"
	"fmt"
	"log"

	repository "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

// authUsecase is the struct for the authentication service
type authUsecase struct {
	userRepo  repository.UserRepository
	adminRepo repository.AdminRepository
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
) usecase.AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

// VerifyUser verifies the user credentials
func (c *authUsecase) VerifyUser(email string, password string) error {

	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

// VerifyAdmin verifies the admin credentials
func (c *authUsecase) VerifyAdmin(email, password string) error {

	admin, err := c.adminRepo.FindAdmin(email)

	//_, err = c.adminRepo.FindAdmin(email)

	if err != nil {
		return errors.New("invalid Username/ password, failed to login")
	}

	fmt.Println("adminpassword", admin.Password)
	fmt.Println("password:", password)

	isValidPassword := VerifyPassword(admin.Password, []byte(password))
	if !isValidPassword {
		return errors.New("invalid username/ Password, failed to login")
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
