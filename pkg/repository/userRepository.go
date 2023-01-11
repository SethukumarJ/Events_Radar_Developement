package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type userRepository struct {
	db *sql.DB
}

// UpdateProfile implements interfaces.UserRepository
func (c *userRepository) UpdateProfile(profile domain.Bios, username string) (int, error) {
	var id int
	query := `UPDATE bios SET 
							about=$1,
							twitter_link=$2,
							github_link=$3,
							linked_in=$4,
							skills=$5,
							qualification=$,
							devfolio=$,6
							website_link=$7 WHERE user_name = $8;`
	err := c.db.QueryRow(query, profile.About,
		profile.Twitter,
		profile.Github,
		profile.Github,
		profile.Qualification,
		profile.WebsiteLink).Scan(&id)

	fmt.Println("id", id)
	return id, err

}

// FindUser implements interfaces.UserRepository
func (c *userRepository) FindUser(email string) (domain.UserResponse, error) {

	var user domain.UserResponse

	query := `SELECT user_id,user_name,first_name,
			  		last_name,email,password,
					phone_number,profile,verification FROM users 
					WHERE email = $1;`

	err := c.db.QueryRow(query, email).Scan(&user.UserId,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.Profile,
		&user.Verification,
	)

	fmt.Println("user from find user :", user)
	return user, err
}

// InsertUser implements interfaces.UserRepository
func (c *userRepository) InsertUser(user domain.Users) (int, error) {
	var id int

	query := `INSERT INTO users(user_name,first_name,last_name,
								email,phone_number,password,
								profile)VALUES($1, $2, $3, $4, $5, $6,$7)
								RETURNING user_id;`

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
func (u *userRepository) StoreVerificationDetails(email string, code int) error {
	var err error
	query := `INSERT INTO verifications (email, code) 
										VALUES ($1, $2);`

	err = u.db.QueryRow(query, email, code).Err()
	return err
}

// VerifyAccount implements interfaces.UserRepository
func (c *userRepository) VerifyAccount(email string, code int) error {
	var useremail string

	query := `SELECT email FROM verifications 
			  WHERE email = $1 AND code = $2;`
	err := c.db.QueryRow(query, email, code).Scan(&useremail)

	fmt.Println("useremail", useremail)

	if err == sql.ErrNoRows {
		return errors.New("invalid verification code/Email")
	}

	if err != nil {
		return err
	}

	query2 := `UPDATE users SET verification = $1 WHERE email = $2`

	err = c.db.QueryRow(query2, true, email).Err()
	log.Println("Updating User verification: ", err)
	if err != nil {
		return err
	}

	query3 := `DELETE FROM verifications WHERE email = $1;`

	err = c.db.QueryRow(query3, email).Err()
	fmt.Println("deleting the verification code.")
	if err != nil {
		return err
	}

	return nil
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
