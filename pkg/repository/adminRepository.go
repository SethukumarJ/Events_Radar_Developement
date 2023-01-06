package repository

import (
	"database/sql"
	"log"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type adminRepository struct {
	db *sql.DB
}

// FindAdmin implements interfaces.AdminRepository
func (c *adminRepository) FindAdmin(adminName string) (domain.AdminResponse, error) {
	log.Println("username of admin:", adminName)
	var admin domain.AdminResponse

	query := `SELECT
			admin_id, 
			admin_name,
			password
			FROM admins WHERE admin_name = $1;`

	err := c.db.QueryRow(query, adminName).Scan(
		&admin.AdminId,
		&admin.AdminName,
		&admin.Password)

	return admin, err
}

// InsertAdmin implements interfaces.AdminRepository
func (c *adminRepository) CreateAdmin(admin domain.Admins) (int, error) {
	var id int
	query := `INSERT INTO admins (admin_name,password)
				VALUES($1, $2)
				RETURNING admin_id;`
	err := c.db.QueryRow(query, admin.AdminName, admin.Password).Scan(&id)
	return id, err
}

func NewAdminRespository(db *sql.DB) interfaces.AdminRepository {
	return &adminRepository{
		db: db,
	}
}
