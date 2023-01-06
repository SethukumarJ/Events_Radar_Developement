package interfaces


import "github.com/thnkrn/go-gin-clean-arch/pkg/domain"

// UserRepository represent the users's repository contract
type AdminRepository interface {
	FindAdmin(email string) (domain.AdminResponse, error)
	InsertUser(admin domain.Admins) (int, error)
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
}
