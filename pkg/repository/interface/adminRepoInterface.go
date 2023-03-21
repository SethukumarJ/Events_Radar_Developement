package interfaces

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

// UserRepository represent the users's repository contract
type AdminRepository interface {
	FindAdminByName(email string) (domain.AdminResponse, error)
	FindAdminById(admin_id int) (domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error)
	AllEvents(pagenation utils.Filter,approved string) ([]domain.EventResponse, utils.Metadata, error)
	SearchEvent(search string) ([]domain.EventResponse,error)
	CreateAdmin(admin domain.Admins) (int, error)
	VipUser(user_id int) error
	ApproveEvent(event_id int) error
	ListOrgRequests(pagenation utils.Filter,applicationStatus string) ([]domain.OrganizationsResponse, utils.Metadata, error)
	RegisterOrganization(orgStatudId int) error
	RejectOrganization(orgStatudId int) error
}
