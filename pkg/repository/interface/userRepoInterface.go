package interfaces

import "github.com/thnkrn/go-gin-clean-arch/pkg/domain"

// UserRepository represent the users's repository contract
type UserRepository interface {
	FindUser(email string) (domain.UserResponse, error)
	InsertUser(user domain.Users) (int, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
}
