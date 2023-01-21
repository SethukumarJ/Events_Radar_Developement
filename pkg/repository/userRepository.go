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




// AdmitMember implements interfaces.UserRepository
func (c *userRepository) AdmitMember(JoinStatusId int, memberRole string) error {
	var organizationName string
	var userName string

	query := `SELECT pending, orgnaization_name FROM join_statuses WHERE join_status_id = $1;`
	err := c.db.QueryRow(query, JoinStatusId).Scan(&userName,&organizationName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	query2 := `UPDATE join_statuses SET pending = null, joined = $1 WHERE join_status_id = $2;`
	err = c.db.QueryRow(query2, organizationName,JoinStatusId).Scan(&organizationName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}
	query3 := `INSERT INTO user_organization_connections(organization_name,user_name,role)
	VALUES($1,$2,$3);`
	err = c.db.QueryRow(query3, organizationName,userName,memberRole).Scan(&organizationName)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}

	return nil
}

// ListJoinRequests implements interfaces.UserRepository
func (c *userRepository) ListJoinRequests(username string, organizationName string) ([]domain.Join_StatusResponse, error) {
	var Requests []domain.Join_StatusResponse

	query := `SELECT COUNT(*) OVER(),join_status_id pending, organization_name FROM join_statuses WHERE organization_name = $1;`

	rows, err := c.db.Query(query, organizationName)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}
	fmt.Println("join statuses called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("joinstatuses called from repo")

	for rows.Next() {
		var joinStatuses domain.Join_StatusResponse
		fmt.Println("organizatioinName :", joinStatuses.OrganizationName)
		err = rows.Scan(
			&totalRecords,
			&joinStatuses.JoinStatusId,
			&joinStatuses.Pending,
			&joinStatuses.OrganizationName,
		)

		fmt.Println("organizatioinName :", joinStatuses.OrganizationName)

		if err != nil {
			return nil, err
		}

		Requests = append(Requests, joinStatuses)
	}
	fmt.Println("Requests", Requests)
	if err := rows.Err(); err != nil {
		return Requests, err
	}
	log.Println(Requests)

	return Requests, nil
}

// FindRelation implements interfaces.UserRepository
func (c *userRepository) FindRelation(username string, organizationName string) (string, error) {
	var role string
	findRole := `SELECT role FROM user_organization_connections WHERE organization_name = $1 AND user_name = $2;`

	err := c.db.QueryRow(findRole, organizationName, username).Scan(&role)
	fmt.Println("role,", role)

	return role, err
}

// AddMembers implements interfaces.UserRepository
func (c *userRepository) AcceptJoinInvitation(newMember string, organizationName string, memberRole string) (int, error) {
	var id int
	var err error
	query := `INSERT INTO user_organization_connections(user_name,organization_name,role)VALUES($1,$2,$3);`

	err = c.db.QueryRow(query, newMember, organizationName, memberRole).Err()
	fmt.Println("err", err)

	fmt.Println("id", id)
	return id, err
}

// FindRole implements interfaces.UserRepository
func (c *userRepository) FindRole(username string, organizationName string) (string, error) {

	var role string
	findRole := `SELECT role FROM user_organization_connections WHERE organization_name = $1 AND user_name = $2;`

	err := c.db.QueryRow(findRole, organizationName, username).Scan(&role)
	fmt.Println("role,", role)

	return role, err
}

// JoinOrganization implements interfaces.UserRepository
func (c *userRepository) JoinOrganization(organizatinName string, username string) (int, error) {
	var id int

	query := `INSERT INTO join_statuses(pending,organization_name)VALUES($1,$2);`
	err := c.db.QueryRow(query, username, organizatinName).Err()

	fmt.Println("id", id)
	return id, err

}

// ListOrganizations implements interfaces.UserRepository
func (c *userRepository) ListOrganizations(pagenation utils.Filter) ([]domain.OrganizationsResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var organizations []domain.OrganizationsResponse

	ListregisteredOrganizations := `SELECT COUNT(*) OVER() AS total_records,org.organization_id,org.organization_name,
	org.created_by,org.logo,org.about,org.created_at,org.linked_in,org.website_link,org.verified ,status.org_status_id 
	FROM organizations AS org INNER JOIN org_statuses AS status 
	ON org.organization_name = status.registered LIMIT $1 OFFSET $2;`

	rows, err := c.db.Query(ListregisteredOrganizations, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("List organizations called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("all organization called from repo")

	for rows.Next() {
		var organization domain.OrganizationsResponse
		fmt.Println("username :", organization.OrganizationName)
		err = rows.Scan(
			&totalRecords,
			&organization.OrganizationId,
			&organization.OrganizationName,
			&organization.CreatedBy,
			&organization.Logo,
			&organization.About,
			&organization.CreatedAt,
			&organization.LinkedIn,
			&organization.WebsiteLink,
			&organization.Verified,
			&organization.OrgStatusId,
		)

		fmt.Println("organization", organization.OrganizationName)

		if err != nil {
			return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		organizations = append(organizations, organization)
	}

	if err := rows.Err(); err != nil {
		return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(organizations)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return organizations, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// FindOrganization implements interfaces.UserRepository
func (c *userRepository) FindOrganization(organizationName string) (domain.OrganizationsResponse, error) {
	var organization domain.OrganizationsResponse

	query := `SELECT organization_id,
					organization_name,
					created_by,
					logo,
					about,
					created_at,
					linked_in,
					website_link,
					verified FROM organizations 
					WHERE organization_name = $1;`

	err := c.db.QueryRow(query, organizationName).Scan(
		&organization.OrganizationId,
		&organization.OrganizationName,
		&organization.CreatedBy,
		&organization.Logo,
		&organization.About,
		&organization.CreatedAt,
		&organization.LinkedIn,
		&organization.WebsiteLink,
		&organization.Verified,
	)

	fmt.Println("organization from find orgnanization :", organization)
	return organization, err
}

// CreateOrganization implements interfaces.UserRepository
func (c *userRepository) CreateOrganization(organization domain.Organizations) (int, error) {
	var id int

	query := `INSERT INTO organizations(organization_name,
										created_by,
										logo,
										about,
										created_at,
										linked_in,
										website_link)VALUES($1, $2, $3, $4, $5, $6,$7)
										RETURNING organization_id;`

	err := c.db.QueryRow(query,
		organization.OrganizationName,
		organization.CreatedBy,
		organization.Logo,
		organization.About,
		organization.CreatedAt,
		organization.LinkedIn,
		organization.WebsiteLink).Scan(&id)

	query2 := `INSERT INTO org_statuses(pending)VALUES($1);`
	c.db.QueryRow(query2, organization.OrganizationName)

	fmt.Println("id", id)
	return id, err

}

// GetQuestions implements interfaces.UserRepository
func (c *userRepository) GetQuestions(title string) ([]domain.FaqaResponse, error) {

	var questions []domain.FaqaResponse

	query := `SELECT COUNT(*) OVER(), question, created_at,user_name FROM faqas WHERE title = $1 AND answer_id = '0';`

	rows, err := c.db.Query(query, title)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}
	fmt.Println("faqas called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("faqas called from repo")

	for rows.Next() {
		var faqas domain.FaqaResponse
		fmt.Println("title :", faqas.Title)
		err = rows.Scan(
			&totalRecords,
			&faqas.Question,
			&faqas.CreatedAt,
			&faqas.UserName,
		)

		fmt.Println("title", faqas.Title)

		if err != nil {
			return nil, err
		}

		questions = append(questions, faqas)
	}
	fmt.Println("FAQA", questions)
	if err := rows.Err(); err != nil {
		return questions, err
	}
	log.Println(questions)

	return questions, nil

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
func (c *userRepository) GetPublicFaqas(title string) ([]domain.QAResponse, error) {
	fmt.Println("faqas called from repo")

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

		qa = append(qa, faqas)
	}
	fmt.Println("FAQA", qa)
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
					WHERE email = $1 OR user_name = $2;`

	err := c.db.QueryRow(query, email, email).Scan(&user.UserId,
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
func (u *userRepository) StoreVerificationDetails(email string, code string) error {
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
