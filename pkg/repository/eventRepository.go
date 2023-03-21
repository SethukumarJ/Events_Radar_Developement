package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	interfaces "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository/interface"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
)

type eventRepository struct {
	db *sql.DB
}



const (
	listPendingAppications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_id,
	app.applied_at,app.first_name,app.last_name,app.event_id,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_id = status.event_id WHERE app.user_id = status.pending AND app.event_id = $1 LIMIT $2 OFFSET $3;`
	listAcceptedApplications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_id,
	app.applied_at,app.first_name,app.last_name,app.event_id,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_id = status.event_id WHERE app.user_id = status.accepted AND app.event_id = $1 LIMIT $2 OFFSET $3;`
	listRejectedApplications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_id,
	app.applied_at,app.first_name,app.last_name,app.event_id,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_id = status.event_id WHERE app.user_id = status.rejected AND app.event_id = $1 LIMIT $2 OFFSET $3;`
)





// FindOrganizationById implements interfaces.EventRepository
func (c *eventRepository) FindOrganizationById(organizaiton_id int) (domain.OrganizationsResponse, error) {
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

	err := c.db.QueryRow(query, organizaiton_id).Scan(
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
// AcceptApplication implements interfaces.EventRepository
func (c *eventRepository) AcceptApplication(applicationStatusId int, event_id int) error {
	var eventId int
	var userId int

	query := `SELECT app.event_id,status.pending
				FROM application_forms AS app INNER JOIN appllication_statuses AS status 
				ON app.user_id = status.pending WHERE status.application_status_id = $1;`
	err := c.db.QueryRow(query, applicationStatusId).Scan(&eventId, &userId)
	if err != nil {
		fmt.Println("1//////////", applicationStatusId, "////////////////////")
		return err
	}

	query2 := `UPDATE appllication_statuses SET pending = null, accepted = $1 WHERE application_status_id = $2;`
	err = c.db.QueryRow(query2, userId, applicationStatusId).Err()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		fmt.Println("2//////////////////////////////")
		return err

	}
	query3 := `UPDATE events SET application_left = application_left - 1 WHERE event_id = $1;`
	err = c.db.QueryRow(query3, event_id).Err()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		fmt.Println("3//////////////////////////////")
		return err

	}

	return nil
}

// ListApplications implements interfaces.EventRepository
func (c *eventRepository) ListApplications(pagenation utils.Filter, applicationStatus string, event_id int) ([]domain.ApplicationFormResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var applications []domain.ApplicationFormResponse
	var rows *sql.Rows
	var err error
	if applicationStatus == "pending" {
		rows, err = c.db.Query(listPendingAppications, event_id, pagenation.Limit(), pagenation.Offset())
	} else if applicationStatus == "accepted" {
		rows, err = c.db.Query(listAcceptedApplications, event_id, pagenation.Limit(), pagenation.Offset())
	} else if applicationStatus == "rejected" {
		rows, err = c.db.Query(listRejectedApplications, event_id, pagenation.Limit(), pagenation.Offset())
	}
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	fmt.Println("List applications called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("applications called from repo")

	for rows.Next() {
		var application domain.ApplicationFormResponse
		fmt.Println("username :", application.UserId)
		err = rows.Scan(
			&totalRecords,
			&application.ApplicationId,
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
			&application.ApplicationStatusId,
		)

		fmt.Println("username", application.UserId)

		if err != nil {
			return applications, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		applications = append(applications, application)
	}

	if err := rows.Err(); err != nil {
		return applications, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(applications)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return applications, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// RejectApplication implements interfaces.EventRepository
func (c *eventRepository) RejectApplication(applicationStatusId int, event_id int) error {
	var eventId int
	var userId int

	query := `SELECT app.event_id,status.pending
				FROM application_forms AS app INNER JOIN appllication_statuses AS status 
				ON app.user_id = status.pending WHERE status.application_status_id = $1;`
	err := c.db.QueryRow(query, applicationStatusId).Scan(&eventId, &userId)
	if err != nil {
		return err
	}

	query2 := `UPDATE appllication_statuses SET pending = null, rejected = $1 WHERE application_status_id = $2;`
	err = c.db.QueryRow(query2, userId, applicationStatusId).Err()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		return err

	}

	return nil
}

// CreatePoster implements interfaces.EventRepository
func (c *eventRepository) CreatePoster(poster domain.Posters) (int, error) {
	var id int

	posterName := poster.Name
	fmt.Println(posterName)
	fmt.Println("evet_id ", poster.EventId)
	query := `INSERT INTO posters(event_id,
									name,
									image,
									discription,
									date,
									colour)VALUES($1, $2, $3, $4, $5, $6)
									RETURNING poster_id;`

	err := c.db.QueryRow(query,
		poster.EventId,
		poster.Name,
		poster.Image,
		poster.Discription,
		poster.Date,
		poster.Colour).Scan(&id)

	fmt.Println(poster.Name, "from repository poster anname")
	if err != nil {
		return id, err
	}

	fmt.Println("id", id)
	return id, err
}

// DeletePoster implements interfaces.EventRepository
func (c *eventRepository) DeletePoster(poster_id int, eventid int) error {
	var id int
	query := `DELETE FROM posters WHERE poster_id = $1 AND event_id = $2 RETURNING poster_id;`

	err := c.db.QueryRow(query, poster_id, eventid).Scan(&id)
	fmt.Println("id deleted:", id)
	if err != nil {
		return err
	}
	return nil
}

// FindPoster implements interfaces.EventRepository
func (c *eventRepository) FindPosterByName(title string, eventid int) (domain.PosterResponse, error) {
	var poster domain.PosterResponse

	query := `SELECT poster_id,name,image,discription,date,colour,event_id FROM posters WHERE name  = $1 AND event_id = $2;`

	err := c.db.QueryRow(query, title, eventid).Scan(
		&poster.PosterId,
		&poster.Name,
		&poster.Image,
		&poster.Discription,
		&poster.Date,
		&poster.Colour,
		&poster.EventId,
	)

	fmt.Println("poster from find poster :", poster)
	return poster, err
}

// FindPoster implements interfaces.EventRepository
func (c *eventRepository) FindPosterById(Poster_id int, eventid int) (domain.PosterResponse, error) {
	var poster domain.PosterResponse

	query := `SELECT poster_id,name,image,discription,date,colour,event_id FROM posters WHERE poster_id  = $1 AND event_id = $2;`

	err := c.db.QueryRow(query, Poster_id, eventid).Scan(
		&poster.PosterId,
		&poster.Name,
		&poster.Image,
		&poster.Discription,
		&poster.Date,
		&poster.Colour,
		&poster.EventId,
	)

	fmt.Println("poster from find poster :", poster)
	return poster, err
}

// PostersByEvent implements interfaces.EventRepository
func (c *eventRepository) PostersByEvent(eventid int) ([]domain.PosterResponse, error) {
	fmt.Println("all posters called from repo")
	var posters []domain.PosterResponse

	query := `SELECT COUNT(*) OVER(),poster_id,name,image,discription,date,colour,event_id FROM posters WHERE event_id = $1ORDER BY date DESC;`

	rows, err := c.db.Query(query, eventid)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}

	fmt.Println("all posters called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("all posters called from repo")

	for rows.Next() {
		var poster domain.PosterResponse
		fmt.Println("postername", poster.Name)
		err = rows.Scan(
			&totalRecords,
			&poster.PosterId,
			&poster.Name,
			&poster.Image,
			&poster.Discription,
			&poster.Date,
			&poster.Colour,
			&poster.EventId,
		)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("title", poster.Name)

		posters = append(posters, poster)
	}

	if err = rows.Err(); err != nil {
		return posters, err
	}
	log.Println(posters)

	return posters, nil
}

// SearchEventUser implements interfaces.EventRepository
func (c *eventRepository) SearchEventUser(search string) ([]domain.EventResponse, error) {
	fmt.Println("allevents called from repo")
	var events []domain.EventResponse

	now := time.Now()
	dateString := now.Format("2006-01-02")
	fmt.Println("currentdate:", dateString)

	query := `SELECT 
		COUNT(*) OVER(),
		event_id,
		title,
		organization_id,
		user_id,
		created_by,
		event_pic,
		short_discription,
		long_discription,
		event_date,
		location,
		created_at,
		approved,
		paid,
		sex,
		cusat_only,
		archived,
		sub_events,
		online,
		max_applications,
		application_closing_date,
		application_link,
		website_link 
			FROM events WHERE event_date > $1 AND approved = true
			AND concat(event_id::text, title, short_discription, long_discription, location) LIKE '%' || $2 || '%' 
			ORDER BY event_date DESC;`

	rows, err := c.db.Query(query, dateString, search)
	fmt.Println("rows", rows)
	if err != nil {
		return nil, err
	}

	fmt.Println("allevents called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allevents called from repo")

	for rows.Next() {
		var event domain.EventResponse
		fmt.Println("username :", event.Title)
		err = rows.Scan(
			&totalRecords,
			&event.EventId,
			&event.Title,
			&event.OrganizationId,
			&event.User_id,
			&event.CreatedBy,
			&event.EventPic,
			&event.ShortDiscription,
			&event.LongDiscription,
			&event.EventDate,
			&event.Location,
			&event.CreatedAt,
			&event.Approved,
			&event.Paid,
			&event.Sex,
			&event.CusatOnly,
			&event.Archived,
			&event.SubEvents,
			&event.Online,
			&event.MaxApplications,
			&event.ApplicationClosingDate,
			&event.ApplicationLink,
			&event.WebsiteLink,
		)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("title", event.Title)

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return events, err
	}
	log.Println(events)

	return events, nil
}

// FindUser implements interfaces.EventRepository
func (c *eventRepository) FindUser(username string) (string, error) {
	query := `SELECT vip FROM users WHERE user_name = $1;`
	var vip string
	err := c.db.QueryRow(query, username).Scan(&vip)
	if err != nil {
		return "", err
	}
	return vip, nil
}

// DeleteEvent implements interfaces.EventRepository
func (c *eventRepository) DeleteEvent(event_id int) error {

	var id int
	query := `DELETE FROM events WHERE event_id = $1 RETURNING event_id;`

	err := c.db.QueryRow(query, event_id).Scan(&id)
	fmt.Println("id deleted:", id)
	if err != nil {
		return err
	}
	return nil

}

// UpdateEvent implements interfaces.EventRepository
func (c *eventRepository) UpdateEvent(event domain.Events, event_id int) (int, error) {
	var id int

	query := `UPDATE events SET
								
								event_pic = $1,
								short_discription = $2,
								long_discription = $3,
								event_date = $4,
								location = $5,
								paid = $6,
								sex = $7,
								cusat_only = $8,
								sub_events = $9,
								online = $10,
								max_applications = $11,
								application_closing_date = $12,
								application_link = $13,
								title = $14,
								website_link = $15 WHERE event_id = $16;`

	err := c.db.QueryRow(query,

		event.EventPic,
		event.ShortDiscription,
		event.LongDiscription,
		event.EventDate,
		event.Location,
		event.Paid,
		event.Sex,
		event.CusatOnly,
		event.SubEvents,
		event.Online,
		event.MaxApplications,
		event.ApplicationClosingDate,
		event.ApplicationLink, event.Title,
		event.WebsiteLink, event_id).Err()

	fmt.Println("err", err)
	return id, err
}

// AllApprovedEvents implements interfaces.EventRepository
func (c *eventRepository) AllApprovedEvents(pagenation utils.Filter, filter utils.FilterEvent) ([]domain.EventResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var events []domain.EventResponse
	var rows *sql.Rows
	var err error
	now := time.Now()
	dateString := now.Format("2006-01-02")
	fmt.Println("currentdate:", dateString)

	query := `SELECT 
			COUNT(*) OVER(),
			event_id,
	title,
	organization_id,
	user_id,
	created_by,
	event_pic,
	short_discription,
	long_discription,
	event_date,
	location,
	created_at,
	approved,
	paid,
	amount,
	sex,
	cusat_only,
	archived,
	sub_events,
	online,
	max_applications,
	application_left,
	application_closing_date,
	application_link,
	website_link  FROM events 
					WHERE event_date > $1 AND approved = true 
					AND cusat_only = $2 AND sex = $3 AND online = $4 ORDER BY event_date DESC  
					LIMIT $5 OFFSET $6;`
	query2 := `SELECT 
			COUNT(*) OVER(),
			event_id,
	title,
	organization_id,
	user_id,
	created_by,
	event_pic,
	short_discription,
	long_discription,
	event_date,
	location,
	created_at,
	approved,
	paid,
	amount,
	sex,
	cusat_only,
	archived,
	sub_events,
	online,
	max_applications,
	application_left,
	application_closing_date,
	application_link,
	website_link   FROM events 
					WHERE event_date > $1 AND approved = true 
					AND cusat_only = $2 AND online = $3 ORDER BY event_date DESC  
					LIMIT $4 OFFSET $5;`

	if filter.Sex == "any" {
		rows, err = c.db.Query(query, dateString, filter.CusatOnly, filter.Sex, filter.Online, pagenation.Limit(), pagenation.Offset())
		fmt.Println("rows", rows)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
	} else {
		rows, err = c.db.Query(query2, dateString, filter.CusatOnly, filter.Online, pagenation.Limit(), pagenation.Offset())
		fmt.Println("rows", rows)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
	}

	fmt.Println("allevents called from repo")

	var totalRecords int

	defer rows.Close()
	fmt.Println("allevents called from repo")

	for rows.Next() {
		var event domain.EventResponse
		fmt.Println("username :", event.Title)
		err = rows.Scan(
			&totalRecords,
		&event.EventId,
		&event.Title,
		&event.OrganizationId,
		&event.User_id,
		&event.CreatedBy,
		&event.EventPic,
		&event.ShortDiscription,
		&event.LongDiscription,
		&event.EventDate,
		&event.Location,
		&event.CreatedAt,
		&event.Approved,
		&event.Paid,
		&event.Amount,
		&event.Sex,
		&event.CusatOnly,
		&event.Archived,
		&event.SubEvents,
		&event.Online,
		&event.MaxApplications,
		&event.ApplicationLeft,
		&event.ApplicationClosingDate,
		&event.ApplicationLink,
		&event.WebsiteLink,
		)

		fmt.Println("title", event.Title)

		if err != nil {
			return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	log.Println(events)
	log.Println(utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return events, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil
}

// FindUser implements interfaces.UserRepository
func (c *eventRepository) FindEventByTitle(title string) (domain.EventResponse, error) {
	var event domain.EventResponse
	var totalRecords int
	query := `SELECT 
	COUNT(*) OVER(),
	event_id,
	title,
	organization_id,
	user_id,
	created_by,
	event_pic,
	short_discription,
	long_discription,
	event_date,
	location,
	created_at,
	approved,
	paid,
	amount,
	sex,
	cusat_only,
	archived,
	sub_events,
	online,
	max_applications,
	application_left,
	application_closing_date,
	application_link,
	website_link  FROM events 
	WHERE title = $1;`

	err := c.db.QueryRow(query, title).Scan(
		&totalRecords,
		&event.EventId,
		&event.Title,
		&event.OrganizationId,
		&event.User_id,
		&event.CreatedBy,
		&event.EventPic,
		&event.ShortDiscription,
		&event.LongDiscription,
		&event.EventDate,
		&event.Location,
		&event.CreatedAt,
		&event.Approved,
		&event.Paid,
		&event.Amount,
		&event.Sex,
		&event.CusatOnly,
		&event.Archived,
		&event.SubEvents,
		&event.Online,
		&event.MaxApplications,
		&event.ApplicationLeft,
		&event.ApplicationClosingDate,
		&event.ApplicationLink,
		&event.WebsiteLink,
	)

	fmt.Println("event from find evnet :", event)
	return event, err
}

// FindUser implements interfaces.UserRepository
func (c *eventRepository) FindEventById(event_id int) (domain.EventResponse, error) {
	var event domain.EventResponse
	var totalRecords int
	query := `SELECT 
	COUNT(*) OVER(),
	event_id,
	title,
	organization_id,
	user_id,
	created_by,
	event_pic,
	short_discription,
	long_discription,
	event_date,
	location,
	created_at,
	approved,
	paid,
	amount,
	sex,
	cusat_only,
	archived,
	sub_events,
	online,
	max_applications,
	application_left,
	application_closing_date,
	application_link,
	website_link FROM events 
	WHERE event_id = $1;`

	err := c.db.QueryRow(query, event_id).Scan(
		&totalRecords,
		&event.EventId,
		&event.Title,
		&event.OrganizationId,
		&event.User_id,
		&event.CreatedBy,
		&event.EventPic,
		&event.ShortDiscription,
		&event.LongDiscription,
		&event.EventDate,
		&event.Location,
		&event.CreatedAt,
		&event.Approved,
		&event.Paid,
		&event.Amount,
		&event.Sex,
		&event.CusatOnly,
		&event.Archived,
		&event.SubEvents,
		&event.Online,
		&event.MaxApplications,
		&event.ApplicationLeft,
		&event.ApplicationClosingDate,
		&event.ApplicationLink,
		&event.WebsiteLink,
	)

	fmt.Println("event from find evnet :", event)
	return event, err
}

// InsertUser implements interfaces.UserRepository
func (c *eventRepository) CreateEvent(event domain.Events) (int, error) {
	var id int
	event.ApplicationLeft = event.MaxApplications
	query := `INSERT INTO events(title,organization_id,user_id,created_by,event_pic,short_discription,long_discription,event_date,location,created_at,approved,paid,amount,sex,cusat_only,archived,sub_events,online,max_applications,application_closing_date,application_link,website_link,application_left)VALUES($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11, $12, $13,$14,$15, $16, $17, $18, $19,$20,$21,$22,$23)RETURNING event_id;`

	err := c.db.QueryRow(query, event.Title,
		event.OrganizationId,
		event.User_id,
		event.CreatedBy,
		event.EventPic,
		event.ShortDiscription,
		event.LongDiscription,
		event.EventDate,
		event.Location,
		event.CreatedAt,
		event.Approved,
		event.Paid,
		event.Amount,
		event.Sex,
		event.CusatOnly,
		event.Archived,
		event.SubEvents,
		event.Online,
		event.MaxApplications,
		event.ApplicationClosingDate,
		event.ApplicationLink,
		event.WebsiteLink, event.ApplicationLeft).Scan(&id)

	fmt.Println(event.Title, "from repository event title")
	if err != nil {
		return id, err
	}

	fmt.Println("id", id)
	return id, err
}

func NewEventRepository(db *sql.DB) interfaces.EventRepository {
	return &eventRepository{
		db: db,
	}
}
