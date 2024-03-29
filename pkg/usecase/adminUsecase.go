package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	interfaces "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type adminUsecase struct {
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
}


// SearchEvent implements interfaces.AdminUsecase
func (c *adminUsecase) SearchEvent(search string) (*[]domain.EventResponse, error) {
	fmt.Println("Search event from usecase called")
	SearchList, err := c.adminRepo.SearchEvent(search)
	fmt.Println("searchList:", SearchList)
	if err != nil {
		fmt.Println("error from list organization from usecase:", err)
		return nil, err
	}

	return &SearchList, nil
}

// ListOrgRequests implements interfaces.AdminUsecase
func (c *adminUsecase) ListOrgRequests(pagenation utils.Filter, applicationStatus string) (*[]domain.OrganizationsResponse, *utils.Metadata, error) {
	fmt.Println("List requests from usecase called")
	OrganizaionList, metadata, err := c.adminRepo.ListOrgRequests(pagenation, applicationStatus)
	fmt.Println("events:", OrganizaionList)
	if err != nil {
		fmt.Println("error from list organization from usecase:", err)
		return nil, &metadata, err
	}

	return &OrganizaionList, &metadata, nil
}

// RegisterOrganization implements interfaces.AdminUsecase
func (c *adminUsecase) RegisterOrganization(orgstatusId int) error {
	err := c.adminRepo.RegisterOrganization(orgstatusId)

	if err != nil {
		return err
	}
	return nil
}

// RejectOrganization implements interfaces.AdminUsecase
func (c *adminUsecase) RejectOrganization(orgstatusId int) error {
	err := c.adminRepo.RejectOrganization(orgstatusId)

	if err != nil {
		return err
	}
	return nil
}

// ApproveEvent implements interfaces.AdminUsecase
func (c *adminUsecase) ApproveEvent(event_id int) error {
	err := c.adminRepo.ApproveEvent(event_id)

	if err != nil {
		return err
	}
	return nil
}

// AllEvents implements interfaces.AdminUsecase
func (c *adminUsecase) AllEvents(pagenation utils.Filter, approved string) (*[]domain.EventResponse, *utils.Metadata, error) {
	fmt.Println("allevents from usecase called")
	events, metadata, err := c.adminRepo.AllEvents(pagenation, approved)
	fmt.Println("events:", events)
	if err != nil {
		fmt.Println("error from allevents usecase:", err)
		return nil, &metadata, err
	}

	return &events, &metadata, nil
}

// Vip implements interfaces.AdminUsecase
func (c *adminUsecase) VipUser(user_id int) error {
	err := c.adminRepo.VipUser(user_id)

	if err != nil {
		return err
	}
	return nil
}

// AllUsers implements interfaces.AdminUsecase
func (c *adminUsecase) AllUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	fmt.Println("alluser from usecase called")
	users, metadata, err := c.adminRepo.AllUsers(pagenation)
	fmt.Println("users:", users)
	if err != nil {
		fmt.Println("error from alluserser usecase:", err)
		return nil, &metadata, err
	}

	return &users, &metadata, nil
}

// CreateUser implements interfaces.UserUseCase
func (c *adminUsecase) CreateAdmin(admin domain.Admins) error {
	fmt.Println("create admin from usecase")
	_, err := c.adminRepo.FindAdminByName(admin.Email)
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
func (c *adminUsecase) FindAdminByName(email string) (*domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdminByName(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUser implements interfaces.UserUseCase
func (c *adminUsecase) FindAdminById(admin_id int) (*domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdminById(admin_id)

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
