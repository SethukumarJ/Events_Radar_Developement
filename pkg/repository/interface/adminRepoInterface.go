package interfaces

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

// UserRepository represent the users's repository contract
type AdminRepository interface {
	FindAdmin(email string) (domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error)
	AllEvents(pagenation utils.Filter,approved string) ([]domain.EventResponse, utils.Metadata, error)
	ListOrgRequests(pagenation utils.Filter,applicationStatus string) ([]domain.OrganizationsResponse, utils.Metadata, error)
	CreateAdmin(admin domain.Admins) (int, error)
	VipUser(username string) error
	ApproveEvent(title string) error
	RegisterOrganization(orgStatudId int) error
	RejectOrganization(orgStatudId int) error
}
