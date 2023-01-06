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
func (*adminRepository) InsertAdmin(admin domain.Admins) (int, error) {
	panic("unimplemented")
}



func NewAdminRespository(db *sql.DB) interfaces.AdminRepository {
	return &adminRepository{
		db: db,
	}
}
