package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type AdminUsecase interface {
	CreateAdmin(admin domain.Admins) error
	FindAdminByName(email string) (*domain.AdminResponse, error)
	FindAdminById(admin_id int) (*domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	AllEvents(pagenation utils.Filter, approved string) (*[]domain.EventResponse, *utils.Metadata, error)
	SearchEvent(search string) (*[]domain.EventResponse, error)
	ApproveEvent(event_id int) error
	VipUser(user_id int) error
	ListOrgRequests(pagenation utils.Filter, applicationStatus string) (*[]domain.OrganizationsResponse, *utils.Metadata, error)
	RegisterOrganization(orgstatus_id int) error
	RejectOrganization(orgstatus_id int) error
	
	
}
