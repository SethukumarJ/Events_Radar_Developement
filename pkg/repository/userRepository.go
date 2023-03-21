package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	interfaces "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type userRepository struct {
	db *sql.DB
}

// FindUser implements interfaces.UserRepository
func (c *userRepository) FindUserById(user_id int) (domain.UserResponse, error) {

	var user domain.UserResponse

	query := `SELECT user_id,user_name,firs_name,last_name,email,password,phone_number,profile,verification FROM users WHERE user_id = $1;`

	err := c.db.QueryRow(query, user_id).Scan(&user.UserId,
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


func (c *userRepository) InsertUser(user domain.Users) (int, error) {
	var id int
	var bio int
	query := `INSERT INTO users(user_name,first_name,last_name,email,phone_number,password,profile)VALUES($1, $2, $3, $4, $5, $6, $7)RETURNING user_id;`

	err := c.db.QueryRow(query, user.UserName,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.Profile).Scan(&id)

	if err != nil {
		return id, err
	}

	query2 := `INSERT INTO bios(user_id)VALUES($1)RETURNING bio_id;`
	err = c.db.QueryRow(query2, id).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		println(bio)
	}
	fmt.Println("id", id)
	return id, err
}

// UpdateRole implements interfaces.UserRepository
func (c *userRepository) UpdateRole(user_id int, organization_id int, updatedRole string) error {

	query := `UPDATE user_organization_connections SET role = $1 WHERE user_id = $2 AND organization_id = $3;`

	err := c.db.QueryRow(query, updatedRole, user_id, organization_id).Err()
	fmt.Println("id updted:")
	if err != nil {
		return err
	}
	return nil
}

// DeleteMember implements interfaces.UserRepository
func (c *userRepository) DeleteMember(user_id int, organization_id int) error {

	query := `DELETE FROM user_organization_connections WHERE user_id = $1 AND organization_id = $2;`

	err := c.db.QueryRow(query, user_id, organization_id).Err()
	fmt.Println("id deleted:")
	if err != nil {
		return err
	}
	return nil
}

// ListMembers implements interfaces.UserRepository
func (c *userRepository) ListMembers(memberRole string, organizaition_id int) ([]domain.UserOrganizationConnectionResponse, error) {

	var members []domain.UserOrganizationConnectionResponse

	query := `SELECT COUNT(*) OVER(),organization_id,user_id, role FROM user_organization_connections WHERE organization_id = $1 AND role = $2;`

	rows, err := c.db.Query(query, organizaition_id, memberRole)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}
	fmt.Println("list memberscalled from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("list members called from repo")

	for rows.Next() {
		var member domain.UserOrganizationConnectionResponse
		fmt.Println("member name :", member.UserId)
		err = rows.Scan(
			&totalRecords,
			&member.OrganizationId,
			&member.UserId,
			&member.Role,
		)

		fmt.Println("organizationName :", member.OrganizationId)

		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}
	fmt.Println("Requests", members)
	if err := rows.Err(); err != nil {
		return members, err
	}
	log.Println(members)

	return members, nil

}

// Prmotion_Success implements interfaces.UserRepository
func (c *userRepository) Prmotion_Success(orderid string, paymentid string) error {
	fmt.Println(orderid, paymentid, "from repo///////////////")
	updatePromotion := `UPDATE promotions SET status = true,payment_id = $1 WHERE order_id =$2;`
	err := c.db.QueryRow(updatePromotion, paymentid, orderid).Err()
	if err != nil {
		return err
	}

	return nil
}

// Prmotion_Success implements interfaces.UserRepository
func (c *userRepository) Prmotion_Faliure(orderid string, paymentid string) error {

	fmt.Println(orderid, paymentid, "from repo hello///////////////")
	updatePromotion := `UPDATE promotions SET status = false, payment_id = $1 WHERE order_id = $2;`
	err := c.db.QueryRow(updatePromotion, paymentid, orderid).Err()
	if err != nil {
		return err
	}
	var event int
	getEvent := `SELECT event_id FROM promotions WHERE order_id = $1;`
	err = c.db.QueryRow(getEvent, orderid).Scan(&event)
	if err != nil {
		return err
	}

	updatePackage := `UPDATE packages SET basic = false, standard = false , premium = false WHERE event_id = $1;`
	err = c.db.QueryRow(updatePackage, event).Err()
	if err != nil {
		return err
	}

	unfeature := `UPDATE events SET featured = false WHERE event_id = $1;`
	err = c.db.QueryRow(unfeature, event).Err()
	if err != nil {
		return err
	}
	return nil
}

// FeaturizeEvent implements interfaces.UserRepository
func (c *userRepository) FeaturizeEvent(orderid string) error {

	var plan, insertPackage string
	var event_id int

	query := `SELECT event_id, plan FROM promotions WHERE order_id = $1;`

	err := c.db.QueryRow(query, orderid).Scan(&event_id, &plan)
	if err != nil {
		fmt.Println("1////////////////////////////")
		return err
	}

	feature := `UPDATE events SET featured = true WHERE event_id = $1`

	err = c.db.QueryRow(feature, event_id).Err()
	if err != nil {
		fmt.Println("2////////////////////////////")
		return err
	}

	packages := `INSERT INTO packages(event_id)VALUES($1)`

	err = c.db.QueryRow(packages, event_id).Err()
	if err != nil {
		fmt.Println("3////////////////////////////")
		return err
	}

	basic := `UPDATE packages SET basic = true WHERE event_id = $1`
	standard := `UPDATE packages SET standard = true WHERE event_id = $1`
	premium := `UPDATE packages SET premium = true WHERE event_id = $1`

	if plan == "basic" {
		insertPackage = basic
	} else if plan == "standard" {
		insertPackage = standard
	} else if plan == "premium" {
		insertPackage = premium
	}

	err = c.db.QueryRow(insertPackage, event_id).Err()
	if err != nil {
		fmt.Println("4////////////////////////////")
		return err
	}

	if plan == "basic" {
		// Start a goroutine to update the featured column to false after 7 days
		go func() {
			time.Sleep(7 * 24 * time.Hour) // Wait for 7 days
			update := `UPDATE events SET featured = false WHERE event_id = $1`
			err := c.db.QueryRow(update, event_id).Err()
			if err != nil {
				fmt.Println("Error updating featured column to false:", err)
			}
		}()
	} else if plan == "standard" {
		// Start a goroutine to update the featured column to false after 7 days
		go func() {
			time.Sleep(10 * 24 * time.Hour) // Wait for 7 days
			update := `UPDATE events SET featured = false WHERE event_id = $1`
			err := c.db.QueryRow(update, event_id).Err()
			if err != nil {
				fmt.Println("Error updating featured column to false:", err)
			}
		}()
	} else if plan == "premium" {

		// Start a goroutine to update the featured column to false after 7 days
		go func() {
			time.Sleep(14 * 24 * time.Hour) // Wait for 7 days
			update := `UPDATE events SET featured = false WHERE event_id = $1`
			err := c.db.QueryRow(update, event_id).Err()
			if err != nil {
				fmt.Println("Error updating featured column to false:", err)
			}
		}()
	}

	return nil
}

// PromoteEvent implements interfaces.UserRepository
func (c *userRepository) PromoteEvent(promotion domain.Promotion) error {
	var id int

	query := `INSERT INTO promotions(order_id,event_id,promoted_by,amount,plan)VALUES($1, $2, $3, $4, $5)RETURNING promotion_id;`

	err := c.db.QueryRow(query,
		promotion.OrderId,
		promotion.EventId,
		promotion.PromotedBy,
		promotion.Amount,
		promotion.Plan).Scan(&id)

	if err != nil {
		return err
	}
	return nil
}

// ApplyEvent implements interfaces.UserRepository
func (c *userRepository) ApplyEvent(applicationForm domain.ApplicationForm) (int, error) {
	var id int

	query := `INSERT INTO application_forms(user_id,
		applied_at,
		first_name,
		last_name,
		event_id,
		proffession,
		college,
		company,
		about,
		email,
		github,
		linkedin)VALUES($1, $2, $3, $4, $5, $6,$7,$8,$9,$10,$11,$12)
										RETURNING application_id;`

	err := c.db.QueryRow(query,
		applicationForm.UserId,
		applicationForm.AppliedAt,
		applicationForm.FirstName,
		applicationForm.LastName,
		applicationForm.EventId,
		applicationForm.Proffession,
		applicationForm.College,
		applicationForm.Company,
		applicationForm.About,
		applicationForm.Email,
		applicationForm.Github,
		applicationForm.Linkedin).Scan(&id)
	if err != nil {
		return -1, err
	}

	query2 := `INSERT INTO appllication_statuses(pending,event_id)VALUES($1,$2);`
	err = c.db.QueryRow(query2, applicationForm.UserId, applicationForm.EventId).Err()
	if err != nil {
		return -1, err
	}
	fmt.Println("id", id)
	return id, err
}

// FindApplication implements interfaces.UserRepository
func (c *userRepository) FindApplication(user_id int, event_id int) (domain.ApplicationFormResponse, error) {
	var application domain.ApplicationFormResponse

	query := `SELECT user_id,
	applied_at,
	first_name,
	last_name,
	event_id,
	proffession,
	college,
	company,
	about,
	email,
	github,
	linkedin FROM application_forms 
					WHERE user_id = $1 AND event_id = $2;`

	err := c.db.QueryRow(query, user_id, event_id).Scan(
		&application.UserId,
		&application.AppliedAt,
		&application.FirstName,
		&application.LastName,
		&application.EventId,
		&application.Proffession,
		&application.College,
		&application.Company,
		&application.About,
		&application.Email,
		&application.Github,
		&application.Linkedin,
	)

	fmt.Println("application from find application :", application)
	return application, err
}

// AdmitMember implements interfaces.UserRepository
func (c *userRepository) AdmitMember(JoinStatusId int, memberRole string) error {
	var organizationId int
	var userId int

	query := `SELECT pending, organization_id FROM join_statuses WHERE join_status_id = $1;`
	err := c.db.QueryRow(query, JoinStatusId).Scan(&userId, &organizationId)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	query3 := `INSERT INTO user_organization_connections(organization_id,user_id,role, joined_at)
	VALUES($1,$2,$3,$4);`
	joined_at := time.Now()
	err = c.db.QueryRow(query3, organizationId, userId, memberRole, joined_at).Scan(&organizationId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}
	query2 := `UPDATE join_statuses SET pending = null, joined = $1 WHERE join_status_id = $2;`
	err = c.db.QueryRow(query2, userId, JoinStatusId).Scan(&organizationId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}

	return nil
}

func (c *userRepository) FindJoinStatus(JoinStatusId int) (int, int, error) {
	var organization_id int
	var user_id int

	query := `SELECT pending, organization_id FROM join_statuses WHERE join_status_id = $1 AND pending IS NOT NULL;`
	err := c.db.QueryRow(query, JoinStatusId).Scan(&user_id, &organization_id)
	if err != nil && err != sql.ErrNoRows {
		return -1, -1, err
	}
	return user_id, organization_id, nil
}

// ListJoinRequests implements interfaces.UserRepository
func (c *userRepository) ListJoinRequests(user_id int, organization_id int) ([]domain.Join_StatusResponse, error) {
	var Requests []domain.Join_StatusResponse

	query := `SELECT COUNT(*) OVER(),join_status_id, pending, organization_id FROM join_statuses WHERE organization_id = $1 AND pending IS NOT NULL;`

	rows, err := c.db.Query(query, organization_id)
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
		fmt.Println("organizatioinName :", joinStatuses.OrganizationId)
		err = rows.Scan(
			&totalRecords,
			&joinStatuses.JoinStatusId,
			&joinStatuses.Pending,
			&joinStatuses.OrganizationId,
		)

		fmt.Println("organizatioinName :", joinStatuses.OrganizationId)

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
func (c *userRepository) FindRelation(user_id int, organizaition_id int) (string, error) {
	var role string
	findRole := `SELECT role FROM user_organization_connections WHERE organization_id = $1 AND user_id = $2;`

	err := c.db.QueryRow(findRole, organizaition_id, user_id).Scan(&role)
	fmt.Println("role,", role)

	return role, err
}

// AddMembers implements interfaces.UserRepository
func (c *userRepository) AcceptJoinInvitation(user_id int, organization_id int, memberRole string) (int, error) {
	var id int
	var err error
	query := `INSERT INTO user_organization_connections(user_id,organization_id,role)VALUES($1,$2,$3);`

	err = c.db.QueryRow(query, user_id, organization_id, memberRole).Err()
	fmt.Println("err", err)

	fmt.Println("id", id)
	return id, err
}

// FindRole implements interfaces.UserRepository
func (c *userRepository) FindRole(user_id int, organization_id int) (string, error) {

	var role string
	findRole := `SELECT role FROM user_organization_connections WHERE organization_id = $1 AND user_id = $2;`

	err := c.db.QueryRow(findRole, organization_id, user_id).Scan(&role)
	fmt.Println("role,", role)

	return role, err
}

// JoinOrganization implements interfaces.UserRepository
func (c *userRepository) JoinOrganization(organizaition_id int, user_id int) (int, error) {
	var id int

	query := `INSERT INTO join_statuses(pending,organization_id)VALUES($1,$2);`
	err := c.db.QueryRow(query, user_id, organizaition_id).Err()

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
	ON org.organization_id = status.registered LIMIT $1 OFFSET $2;`

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
func (c *userRepository) FindOrganizationByName(organizationName string) (domain.OrganizationsResponse, error) {
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

// FindOrganization implements interfaces.UserRepository
func (c *userRepository) FindOrganizationById(organization_id int) (domain.OrganizationsResponse, error) {
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
					WHERE organization_id = $1;`

	err := c.db.QueryRow(query, organization_id).Scan(
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
	c.db.QueryRow(query2, id)

	fmt.Println("id", id)
	return id, err

}

// GetQuestions implements interfaces.UserRepository
func (c *userRepository) GetQuestions(event_id int) ([]domain.FaqaResponse, error) {

	var questions []domain.FaqaResponse

	query := `SELECT COUNT(*) OVER(),faqa_id, event_id, question, created_at,user_name FROM faqas WHERE event_id = $1 AND answer_id = '0';`

	rows, err := c.db.Query(query, event_id)
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
		fmt.Println("title :", faqas.EventId)
		err = rows.Scan(
			&totalRecords,
			&faqas.FaqaId,
			&faqas.EventId,
			&faqas.Question,
			&faqas.CreatedAt,
			&faqas.UserName,
		)

		fmt.Println("title", faqas.EventId)

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
func (c *userRepository) GetPublicFaqas(event_id int) ([]domain.QAResponse, error) {
	fmt.Println("faqas called from repo")

	var qa []domain.QAResponse

	query := `SELECT COUNT(*) OVER() AS total_records,que.faqa_id,que.question,que.answer_id,
	que.event_id,que.created_at,que.user_name,que.organization_id ,ans.answer 
	FROM faqas AS que INNER JOIN answers AS ans 
	ON que.answer_id = ans.answer_id WHERE que.public = $1 AND event_id = $2;`

	rows, err := c.db.Query(query, true, event_id)
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
		fmt.Println("username :", faqas.EventId)
		err = rows.Scan(
			&totalRecords,
			&faqas.FaqaId,
			&faqas.Question,
			&faqas.AnswerId,
			&faqas.EventId,
			&faqas.CreatedAt,
			&faqas.UserName,
			&faqas.OrganizaitionId,
			&faqas.Answer)

		fmt.Println("title", faqas.EventId)

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
		event_id,
		created_at,
		user_name,
		organization_id
		)VALUES($1, $2, $3, $4,$5)RETURNING faqa_id;`

	err := c.db.QueryRow(query,
		question.Question,
		question.EventId,
		question.CreatedAt,
		question.UserName,
		question.OrganizationId).Scan(&id)

	fmt.Println("id", id)
	return id, err
}

// UpdatePassword implements interfaces.UserRepository
func (c *userRepository) UpdatePassword(password string, email string) (int, error) {

	query := `UPDATE users SET password =$1 WHERE email = $2`

	err := c.db.QueryRow(query, password, email).Err()

	if err != nil {
		return 0, err
	}
	return 0, nil

}

// UpdateProfile implements interfaces.UserRepository
func (c *userRepository) UpdateProfile(profile domain.Bios, user_id int) (int, error) {
	var id int
	query := `UPDATE bios SET 
							about=$1,
							twitter = $2,
							github = $3,
							linked_in = $4,
							skills =$5,
							qualification=$6,
							dev_folio=$7,
							website_link=$8 WHERE user_id = $9;`
	err := c.db.QueryRow(query, profile.About,
		profile.Twitter,
		profile.Github,
		profile.LinkedIn,
		profile.Skills,
		profile.Qualification,
		profile.DevFolio,
		profile.WebsiteLink, user_id).Scan(&id)

	fmt.Println("id", id)
	return id, err

}

// FindUser implements interfaces.UserRepository
func (c *userRepository) FindUserByName(email string) (domain.UserResponse, error) {

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
