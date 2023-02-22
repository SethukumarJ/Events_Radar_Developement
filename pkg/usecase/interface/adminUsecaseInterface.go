package interfaces

import (
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type AdminUsecase interface {
	CreateAdmin(admin domain.Admins) error
	FindAdmin(email string) (*domain.AdminResponse, error)
	AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	AllEvents(pagenation utils.Filter, approved string) (*[]domain.EventResponse, *utils.Metadata, error)
	SearchEvent(search string) (*[]domain.EventResponse, error)
	ListOrgRequests(pagenation utils.Filter, applicationStatus string) (*[]domain.OrganizationsResponse, *utils.Metadata, error)
	ApproveEvent(title string) error
	RegisterOrganization(orgstatusId int) error
	RejectOrganization(orgstatusId int) error
	VipUser(username string) error
}
