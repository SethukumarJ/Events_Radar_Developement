package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type adminRepository struct {
	db *sql.DB
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
				profile
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