package repository

import (
	"database/sql"
	"fmt"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type userRepository struct {
	db *sql.DB
}

// FindUser implements interfaces.UserRepository
func (c *userRepository) FindUser(email string) (domain.UserResponse, error) {

	var user domain.UserResponse

	query := `SELECT userid,username,firstname,
			  		lastname,email,password,
					phonenumber,profileFROM users 
					WHERE email = $1;`

	err := c.db.QueryRow(query,email).Scan(	&user.UserId,
											&user.UserName,
											&user.FirstName,
											&user.LastName,
											&user.Email,
											&user.Password,
											&user.PhoneNumber,
											&user.Profile,
										)							

	fmt.Println("user from find user :", user)
	return user, err
}

// InsertUser implements interfaces.UserRepository
func (c *userRepository) InsertUser(user domain.Users) (int, error) {
	var id int

	query := `INSERT INTO users(username,firstname,lastname,
								email,phonenumber,password,
								profile)VALUES($1, $2, $3, $4, $5, $6,$7)
								RETURNING id;`

	err := c.db.QueryRow(query, user.UserName,
								user.FirstName,
								user.LastName,
								user.Email,
								user.PhoneNumber,
								user.Password,
								user.Profile).Scan(&id)

	fmt.Println("id", id)
	return id, err
}


// StoreVerificationDetails implements interfaces.UserRepository
func (*userRepository) StoreVerificationDetails(email string, code int) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.UserRepository
func (*userRepository) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

// UserId       uint   `json:"userid"`
// UserName     string `json:"username"`
// FirstName    string `json:"firstname"`
// LastName     string `json:"lastname"`
// Password     string `json:"password"`
// Email        string `json:"email"`
// Verification bool   `json:"verification" `
// Vip          bool   `json:"vip" `
// PhoneNumber  string `json:"phonenumber"`
// Profile      string `json:"profile"`
// Token        string `json:"token"`
