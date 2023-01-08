package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type AdminUsecase interface {
	CreateAdmin(admin domain.Admins) error
	FindAdmin(email string) (*domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	VipUser(username string) error
}