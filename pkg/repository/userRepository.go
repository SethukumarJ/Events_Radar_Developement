package repository

import (
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// FindUser implements interfaces.UserRepository
func (*userDatabase) FindUser(email string) (domain.UserResponse, error) {
	panic("unimplemented")
}

// InsertUser implements interfaces.UserRepository
func (*userDatabase) InsertUser(user domain.Users) (int, error) {
	panic("unimplemented")
}

// StoreVerificationDetails implements interfaces.UserRepository
func (*userDatabase) StoreVerificationDetails(email string, code int) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.UserRepository
func (*userDatabase) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
