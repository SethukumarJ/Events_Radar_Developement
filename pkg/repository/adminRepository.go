package repository

import (
	"database/sql"
	"fmt"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type adminRepository struct {
	db *sql.DB
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
