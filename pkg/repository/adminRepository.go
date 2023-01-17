package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type adminRepository struct {
	db *sql.DB
}
const (
	listPendingOrganizations = `SELECT COUNT(*) OVER() AS total_records,org.organization_id,org.organization_name,
	org.created_by,org.logo,org.about,org.created_at,org.linked_in,org.website_link,org.verified ,status.org_status_id 
	FROM organizations AS org INNER JOIN org_statuses AS status 
	ON org.organization_name = status.pending LIMIT $1 OFFSET $2;`
	listregisteredOrganizations = `SELECT COUNT(*) OVER() AS total_records,org.organization_id,org.organization_name,
	org.created_by,org.logo,org.about,org.created_at,org.linked_in,org.website_link,org.verified ,status.org_status_id 
	FROM organizations AS org INNER JOIN org_statuses AS status 
	ON org.organization_name = status.registered LIMIT $1 OFFSET $2;`
	listRejectedOrganizations = `SELECT COUNT(*) OVER() AS total_records,org.organization_id,org.organization_name,
	org.created_by,org.logo,org.about,org.created_at,org.linked_in,org.website_link,org.verified ,status.org_status_id 
	FROM organizations AS org INNER JOIN org_statuses AS status 
	ON org.organization_name = status.rejected LIMIT $1 OFFSET $2;`

)

// ListOrgRequests implements interfaces.AdminRepository
func (c *adminRepository) ListOrgRequests(pagenation utils.Filter, applicationStatus string) ([]domain.OrganizationsResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var organizations []domain.OrganizationsResponse

	if applicationStatus == "pending" {

		rows, err := c.db.Query(listPendingOrganizations, pagenation.Limit(), pagenation.Offset())
		fmt.Println("rows", rows)
		if err != nil {
			return nil, utils.Metadata{}, err
		}

		fmt.Println("List organizations called from repo")

		var totalRecords int

		defer rows.Close()
		fmt.Println("allevents called from repo")

		for rows.Next() {
			var organization domain.OrganizationsResponse
			fmt.Println("username :", organization.OrganizationName)
			err = rows.Scan(
				&totalRecords,
				&organization.OrganizationId,
				&organization.OrganizationName,
				&organization.CreatedBy,
				&organization.Logo,
				&organization.About,
				&organization.CreatedAt,
				&organization.LinkedIn,
				&organization.WebsiteLink,
				&organization.Verified,
				&organization.OrgStatusId,
			)

			fmt.Println("organization", organization.OrganizationName)

			if err != nil {
				return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
			}
			organizations = append(organizations, organization)
		}

		if err := rows.Err(); err != nil {
			return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		log.Println(organizations)
		log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
		return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

	} else if applicationStatus == "registered" {

		rows, err := c.db.Query(listregisteredOrganizations, pagenation.Limit(), pagenation.Offset())
		fmt.Println("rows", rows)
		if err != nil {
			return nil, utils.Metadata{}, err
		}

		fmt.Println("List organizations called from repo")

		var totalRecords int

		defer rows.Close()
		fmt.Println("allevents called from repo")

		for rows.Next() {
			var organization domain.OrganizationsResponse
			fmt.Println("username :", organization.OrganizationName)
			err = rows.Scan(
				&totalRecords,
				&organization.OrganizationId,
				&organization.OrganizationName,
				&organization.CreatedBy,
				&organization.Logo,
				&organization.About,
				&organization.CreatedAt,
				&organization.LinkedIn,
				&organization.WebsiteLink,
				&organization.Verified,
				&organization.OrgStatusId,
			)

			fmt.Println("organization", organization.OrganizationName)

			if err != nil {
				return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
			}
			organizations = append(organizations, organization)
		}

		if err := rows.Err(); err != nil {
			return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		log.Println(organizations)
		log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
		return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
	} 

	rows, err := c.db.Query(listRejectedOrganizations, pagenation.Limit(), pagenation.Offset())
		fmt.Println("rows", rows)
		if err != nil {
			return nil, utils.Metadata{}, err
		}

		fmt.Println("List organizations called from repo")

		var totalRecords int

		defer rows.Close()
		fmt.Println("allevents called from repo")

		for rows.Next() {
			var organization domain.OrganizationsResponse
			fmt.Println("username :", organization.OrganizationName)
			err = rows.Scan(
				&totalRecords,
				&organization.OrganizationId,
				&organization.OrganizationName,
				&organization.CreatedBy,
				&organization.Logo,
				&organization.About,
				&organization.CreatedAt,
				&organization.LinkedIn,
				&organization.WebsiteLink,
				&organization.Verified,
				&organization.OrgStatusId,
			)

			fmt.Println("organization", organization.OrganizationName)

			if err != nil {
				return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
			}
			organizations = append(organizations, organization)
		}

		if err := rows.Err(); err != nil {
			return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		log.Println(organizations)
		log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
		return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// RegisterOrganization implements interfaces.AdminRepository
func (c *adminRepository) RegisterOrganization(orgStatudId int) error {
	var organizationName string
	var userName string

	query := `SELECT org.created_by,status.pending
				FROM organizations AS org INNER JOIN org_statuses AS status 
				ON org.organization_name = status.pending WHERE status.org_status_id = $1;`
	err := c.db.QueryRow(query, orgStatudId).Scan(&userName,&organizationName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	query2 := `UPDATE org_statuses SET pending = null, registered = $1;`
	err = c.db.QueryRow(query2, organizationName).Scan(&organizationName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}
	query3 := `INSERT INTO user_organization_connections(organization_name,user_name,role)
	VALUES($1,$2,$3);`
	err = c.db.QueryRow(query3, organizationName,userName,"1").Scan(&organizationName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}

	return nil
}

// RejectOrganization implements interfaces.AdminRepository
func (c *adminRepository) RejectOrganization(orgStatudId int) error {
	var organizationName string
	query := `SELECT pending FROM org_statuses WHERE org_status_id = $1;`
	err := c.db.QueryRow(query, orgStatudId).Scan(&organizationName)
	if err != nil {
		return err
	}

	query2 := `UPDATE org_statuses SET pending = null, rejected = $1;`
	err = c.db.QueryRow(query2, organizationName).Scan(&organizationName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}

	return nil
}

// ApproveEvent implements interfaces.AdminRepository
func (c *adminRepository) ApproveEvent(title string) error {
	var event_name string

	query := `SELECT title FROM 
				events WHERE 
				title = $1;`
	err := c.db.QueryRow(query, title).Scan(&event_name)

	if err == sql.ErrNoRows {
		return errors.New("invalid title")
	}

	if err != nil {
		return err
	}

	query = `UPDATE events SET
				approved = $1
				WHERE
				title = $2 ;`
	err = c.db.QueryRow(query, true, title).Err()
	log.Println("approved the event successfully", err)
	if err != nil {
		return err
	}

	return nil
}

// AllEvents implements interfaces.AdminRepository
func (c *adminRepository) AllEvents(pagenation utils.Filter, approved string) ([]domain.EventResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var events []domain.EventResponse

	now := time.Now()
	dateString := now.Format("2006-01-02")
	fmt.Println("currentdate:", dateString)

	query := `SELECT 
					COUNT(*) OVER(),
					event_id,
					title,
					organizer_name,
					event_pic,
					short_discription,
					long_discription,
					event_date,
					location,
					created_at,
					approved,
					paid,
					sex,
					cusat_only,
					archived,
					sub_events,
					online,
					max_applications,
					application_closing_date,
					application_link,
					website_link FROM events WHERE event_date > $1 AND approved = $2
					LIMIT $3 OFFSET $4;`

	rows, err := c.db.Query(query, dateString, approved, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("allevents called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allevents called from repo")

	for rows.Next() {
		var event domain.EventResponse
		fmt.Println("username :", event.Title)
		err = rows.Scan(
			&totalRecords,
			&event.EventId,
			&event.Title,
			&event.OrganizerName,
			&event.EventPic,
			&event.ShortDiscription,
			&event.LongDiscription,
			&event.EventDate,
			&event.Location,
			&event.CreatedAt,
			&event.Approved,
			&event.Paid,
			&event.Sex,
			&event.CusatOnly,
			&event.Archived,
			&event.SubEvents,
			&event.Online,
			&event.MaxApplications,
			&event.ApplicationClosingDate,
			&event.ApplicationLink,
			&event.WebsiteLink,
		)

		fmt.Println("title", event.Title)

		if err != nil {
			return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(events)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// VipUser implements interfaces.AdminRepository
func (c *adminRepository) VipUser(username string) error {
	var user_name string

	query := `SELECT user_name FROM 
				users WHERE 
				user_name = $1;`
	err := c.db.QueryRow(query, username).Scan(&user_name)

	if err == sql.ErrNoRows {
		return errors.New("invalid title")
	}

	if err != nil {
		return err
	}

	query = `UPDATE users SET
				vip = $1
				WHERE
				user_name = $2 ;`
	err = c.db.QueryRow(query, true, username).Err()
	log.Println("Updating vip status to true ", err)
	if err != nil {
		return err
	}

	return nil
}

// AllUsers implements interfaces.UserRepository
func (c *adminRepository) AllUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {

	fmt.Println("allusers called from repo")
	var users []domain.UserResponse

	query := `SELECT 
				COUNT(*) OVER(),
				user_id,
				first_name,
				last_name,
				user_name,
				email,
				phone_number,
				vip,
				profile,
				verification
				FROM users
				LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("allusers called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allusers called from repo")

	for rows.Next() {
		var User domain.UserResponse
		fmt.Println("username :", User.UserName)
		err = rows.Scan(
			&totalRecords,
			&User.UserId,
			&User.FirstName,
			&User.LastName,
			&User.UserName,
			&User.Email,
			&User.PhoneNumber,
			&User.Vip,
			&User.Profile,
			&User.Verification,
		)

		fmt.Println("username", User.UserName)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		users = append(users, User)
	}

	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

}

// FindUser implements interfaces.UserRepository
func (c *adminRepository) FindAdmin(email string) (domain.AdminResponse, error) {

	var admin domain.AdminResponse

	query := `SELECT admin_id,admin_name,email,password,
					phone_number FROM admins 
					WHERE email = $1;`

	err := c.db.QueryRow(query, email).Scan(&admin.AdminId,
		&admin.AdminName,
		&admin.Email,
		&admin.Password,
		&admin.PhoneNumber,
	)

	fmt.Println("admin from find admin :", admin)
	return admin, err
}

// InsertUser implements interfaces.UserRepository
func (c *adminRepository) CreateAdmin(admin domain.Admins) (int, error) {
	var id int

	query := `INSERT INTO admins(admin_name,
								email,phone_number,password)VALUES($1, $2, $3, $4)
								RETURNING admin_id;`

	err := c.db.QueryRow(query, admin.AdminName,

		admin.Email,
		admin.PhoneNumber,
		admin.Password).Scan(&id)

	fmt.Println("id", id)
	return id, err
}

func NewAdminRespository(db *sql.DB) interfaces.AdminRepository {
	return &adminRepository{
		db: db,
	}
}
