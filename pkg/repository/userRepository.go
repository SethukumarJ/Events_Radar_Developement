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





// PostAnswer implements interfaces.UserRepository
func (c *userRepository) PostAnswer(answer domain.Answers, question int) (int, error) {
	var id int

	query := `INSERT INTO answers(answer)VALUES($1)RETURNING answer_id;`

	err := c.db.QueryRow(query,
		answer.Answer).Scan(&id)

	if err != nil {
		return 0, err
	}
	query2 := `UPDATE faqas SET answer_id = $1 , public = $2 WHERE faqa_id = $3`
	err = c.db.QueryRow(query2,
		id, true, question).Err()

	fmt.Println("id", id)
	return id, err
}

// GetPublicFaqas implements interfaces.UserRepository
func (c *userRepository) GetPublicFaqas(title string) ([]domain.QAResponse , error) {
	fmt.Println("faqas called from repo")

	var queanda []map[interface{}]interface{}
	var qa []domain.QAResponse
	
	query := `SELECT COUNT(*) OVER() AS total_records,que.faqa_id,que.question,que.answer_id,
	que.title,que.created_at,que.user_name,que.organizer_name ,ans.answer 
	FROM faqas AS que INNER JOIN answers AS ans 
	ON que.answer_id = ans.answer_id WHERE que.public = $1 AND title = $2;`
					

	rows, err := c.db.Query(query, true, title)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}

	fmt.Println("faqas called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("faqas called from repo")

	for rows.Next() {
		var faqas domain.QAResponse
		fmt.Println("username :", faqas.Title)
		err = rows.Scan(
			&totalRecords,
			&faqas.FaqaId,
			&faqas.Question,
			&faqas.AnswerId,
			&faqas.Title,
			&faqas.CreatedAt,
			&faqas.OrganizerName,
			&faqas.OrganizerName,
		    &faqas.Answer)

		fmt.Println("title", faqas.Title)
		
		if err != nil {
			return nil, err
		}
		
		qa = append(qa,faqas)
	}
	fmt.Println("queanda",queanda)
	if err := rows.Err(); err != nil {
		return qa, err
	}
	log.Println(qa)

	return qa, nil
}

// PostQuestion implements interfaces.UserRepository
func (c *userRepository) PostQuestion(question domain.Faqas) (int, error) {
	var id int

	query := `INSERT INTO faqas(question,
		title,
		created_at,
		user_name,
		organizer_name
		)VALUES($1, $2, $3, $4,$5)RETURNING faqa_id;`

	err := c.db.QueryRow(query,
		question.Question,
		question.Title,
		question.CreatedAt,
		question.UserName,
		question.OrganizerName).Scan(&id)

	fmt.Println("id", id)
	return id, err
}

// UpdatePassword implements interfaces.UserRepository
func (c *userRepository) UpdatePassword(user domain.Users, email string) (int, error) {

	query := `UPDATE users SET password =$1 WHERE email = $2`

	err := c.db.QueryRow(query, user.Password, email).Err()

	if err != nil {
		return 0, err
	}
	return 0, nil

}

// UpdateProfile implements interfaces.UserRepository
func (c *userRepository) UpdateProfile(profile domain.Bios, username string) (int, error) {
	var id int
	query := `UPDATE bios SET 
							about=$1,
							twitter = $2,
							github = $3,
							linked_in = $4,
							skills =$5,
							qualification=$6,
							dev_folio=$7,
							website_link=$8 WHERE user_name = $9;`
	err := c.db.QueryRow(query, profile.About,
		profile.Twitter,
		profile.Github,
		profile.LinkedIn,
		profile.Skills,
		profile.Qualification,
		profile.DevFolio,
		profile.WebsiteLink, username).Scan(&id)

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

	query2 := `INSERT INTO bios(user_name)VALUES($1);`
	c.db.QueryRow(query2, user.UserName)

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
func (c *userRepository) VerifyAccount(email string, code string) error {
	var useremail string

	query := `SELECT email FROM verifications 
			  WHERE email = $1 AND code = $2;`
	query3 := `DELETE FROM verifications WHERE email = $1;`
	err := c.db.QueryRow(query, email, code).Scan(&useremail)

	fmt.Println("useremail", useremail)

	if err == sql.ErrNoRows {
		err = c.db.QueryRow(query3, email).Err()
		fmt.Println("deleting the verification code.")
		if err != nil {
			return err
		}

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
