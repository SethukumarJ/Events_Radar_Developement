package usecase

import (
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
func (*AdminUsecase) CreateAdmin(admin domain.Admins) error {
	panic("unimplemented")
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
