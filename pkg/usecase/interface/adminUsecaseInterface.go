package interfaces

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type AdminUsecase interface {
	CreateAdmin(admin domain.Admins) error
	FindAdmin(email string) (*domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	AllEvents(pagenation utils.Filter) (*[]domain.EventResponse, *utils.Metadata, error)
	ApproveEvent(title string) error
	VipUser(username string) error
}
