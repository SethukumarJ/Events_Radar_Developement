package interfaces

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (domain.UserResponse, error)
	InsertUser(user domain.Users) (int, error)
	AllUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
}
