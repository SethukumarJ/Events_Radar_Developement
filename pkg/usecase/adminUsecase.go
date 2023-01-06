package usecase

import (
	"database/sql"
	"errors"
	"log"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type AdminUsecase struct {
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
}

// CreateAdmin implements interfaces.AdminUsecase
func (c *AdminUsecase) CreateAdmin(admin domain.Admins) error {
	_, err := c.adminRepo.FindAdmin(admin.AdminName)

	if err == nil {
		return errors.New("admin already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	admin.Password = HashPassword(admin.Password)
	_,err = c.adminRepo.CreateAdmin(admin)

	if err != nil {
		log.Println(err)
		return errors.New("error while signup")
	}
	return nil
}

// FindAdmin implements interfaces.AdminUsecase
func (*AdminUsecase) FindAdmin(email string) (*domain.AdminResponse, error) {
	panic("unimplemented")
}

func NewAdminUsecase(
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig,
	config config.Config) usecase.AdminUsecase {
	return &AdminUsecase{
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
		config:     config,
	}
}
