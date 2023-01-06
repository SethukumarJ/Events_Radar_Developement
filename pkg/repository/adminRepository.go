package repository

import (
	"database/sql"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type adminRepository struct {
	db *sql.DB
}

// FindAdmin implements interfaces.AdminRepository
func (*adminRepository) FindAdmin(email string) (domain.AdminResponse, error) {
	panic("unimplemented")
}

// InsertAdmin implements interfaces.AdminRepository
func (c *adminRepository) CreateAdmin(admin domain.Admins) (int, error) {
	var id int
	query := `INSERT INTO admins (admin_name,password)
				VALUES($1, $2)
				RETURNING admin_id;`
	err := c.db.QueryRow(query, admin.AdminName,admin.Password,).Scan(&id)
	return id ,err
}



func NewAdminRespository(db *sql.DB) interfaces.AdminRepository {
	return &adminRepository{
		db: db,
	}
}
