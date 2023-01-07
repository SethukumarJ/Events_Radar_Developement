package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type userRepository struct {
	db *sql.DB
}

// AllUsers implements interfaces.UserRepository
func (c *userRepository) AllUsers(pagenation utils.Filter) ([]domain.UserResponse, utils.Metadata, error) {
	
	fmt.Println("allusers called from repo")
	var users []domain.UserResponse

	query := `SELECT 
				COUNT(*) OVER(),
				user_id,
				first_name,
				last_name,
				user_name,
				email,
				phone_number,
				profile,
				vip,
				FROM users
				LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(query, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("allusers called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allusers called from repo")

	for rows.Next() {
		var User domain.UserResponse
		fmt.Println("username :", User.UserName)
		err = rows.Scan(
			&totalRecords,
			&User.UserId,
			&User.FirstName,
			&User.LastName,
			&User.UserName,
			&User.Email,
			&User.PhoneNumber,
			&User.Vip,
			&User.Profile,
		)

		fmt.Println("username", User.UserName)

		if err != nil {
			return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		users = append(users, User)
	}

	if err := rows.Err(); err != nil {
		return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(users)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return users, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

}

// FindUser implements interfaces.UserRepository
func (c *userRepository) FindUser(email string) (domain.UserResponse, error) {

	var user domain.UserResponse

	query := `SELECT user_id,user_name,first_name,
			  		last_name,email,password,
					phone_number,profile FROM users 
					WHERE email = $1;`

	err := c.db.QueryRow(query, email).Scan(&user.UserId,
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

	if err == sql.ErrNoRows {
		return errors.New("invalid verification code/Email")
	}

	if err != nil {
		return err
	}

	query = `UPDATE users SET verification = $1
			WHERE email = $2 ;`

	err = c.db.QueryRow(query, true, email).Err()
	log.Println("Updating User verification: ", err)
	if err != nil {
		return err
	}

	query = `DELETE FROM verifications WHERE email = $1;`

	err = c.db.QueryRow(query, email).Err()
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
