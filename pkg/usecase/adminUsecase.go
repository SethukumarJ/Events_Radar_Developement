package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type adminUsecase struct {
	adminRepo  interfaces.AdminRepository
	userRepo    interfaces.UserRepository
	mailConfig config.MailConfig
	config     config.Config
}

// AllUsers implements interfaces.AdminUsecase
func (c *adminUsecase) AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	
	users, metadata, err := c.userRepo.AllUsers(pagenation)
	if err != nil {
		return nil, &metadata, err
	}

	return &users, &metadata, nil
}

// CreateUser implements interfaces.UserUseCase
func (c *adminUsecase) CreateAdmin(admin domain.Admins) error {
	fmt.Println("create admin from usecase")
	_, err := c.adminRepo.FindAdmin(admin.Email)
	fmt.Println("found admin", err)

	if err == nil {
		return errors.New("adminname already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	admin.Password = HashPassword(admin.Password)
	fmt.Println("password", admin.Password)
	_, err = c.adminRepo.CreateAdmin(admin)
	if err != nil {
		return err
	}
	return nil
}

// FindUser implements interfaces.UserUseCase
func (c *adminUsecase) FindAdmin(email string) (*domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdmin(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func NewAdminUsecase(
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig,
	config config.Config) usecase.AdminUsecase {
	return &adminUsecase{
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
		config:     config,
	}
}
