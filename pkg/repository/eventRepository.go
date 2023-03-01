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
	listPendingAppications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_name,
	app.applied_at,app.first_name,app.last_name,app.event_name,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_name = status.event_name WHERE app.user_name = status.pending AND app.event_name = $1 LIMIT $2 OFFSET $3;`
	listAcceptedApplications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_name,
	app.applied_at,app.first_name,app.last_name,app.event_name,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_name = status.event_name WHERE app.user_name = status.accepted AND app.event_name = $1 LIMIT $2 OFFSET $3;`
	listRejectedApplications = `SELECT COUNT(*) OVER() AS total_records,app.application_id,app.user_name,
	app.applied_at,app.first_name,app.last_name,app.event_name,app.proffession,app.college,app.company,app.about,app.email,app.github,app.linkedin,status.application_status_id 
	FROM application_forms AS app INNER JOIN appllication_statuses AS status 
	ON app.event_name = status.event_name WHERE app.user_name = status.rejected AND app.event_name = $1 LIMIT $2 OFFSET $3;`
)

// AcceptApplication implements interfaces.EventRepository
func (c *eventRepository) AcceptApplication(applicationStatusId int,eventname string) error {
	var eventName string
	var userName string

	query := `SELECT app.event_name,status.pending
				FROM application_forms AS app INNER JOIN appllication_statuses AS status 
				ON app.user_name = status.pending WHERE status.application_status_id = $1;`
	err := c.db.QueryRow(query, applicationStatusId).Scan(&eventName, &userName)
	if err != nil {
		fmt.Println("1//////////",applicationStatusId,"////////////////////")
		return err
	}

	query2 := `UPDATE appllication_statuses SET pending = null, accepted = $1 WHERE application_status_id = $2;`
	err = c.db.QueryRow(query2, userName, applicationStatusId).Err()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		fmt.Println("2//////////////////////////////")
		return err

	}
	query3 := `UPDATE events SET application_left = application_left - 1 WHERE title = $1;`
	err = c.db.QueryRow(query3, eventName).Err()
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("err", err)
		fmt.Println("3//////////////////////////////")
		return err

	}

	return nil
}

// ListApplications implements interfaces.EventRepository
func (c *eventRepository) ListApplications(pagenation utils.Filter, applicationStatus string,eventname string) ([]domain.ApplicationFormResponse, utils.Metadata, error) {
	fmt.Println("allevents called from repo")
	var applications []domain.ApplicationFormResponse
	var rows *sql.Rows
	var err error
	if applicationStatus == "pending" {
		rows, err = c.db.Query(listPendingAppications,eventname, pagenation.Limit(), pagenation.Offset())
	} else if applicationStatus == "accepted" {
		rows, err = c.db.Query(listAcceptedApplications,eventname,  pagenation.Limit(), pagenation.Offset())
	} else if applicationStatus == "rejected" {
		rows, err = c.db.Query(listRejectedApplications,eventname,  pagenation.Limit(), pagenation.Offset())
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
		fmt.Println("username :", application.UserName)
		err = rows.Scan(
			&totalRecords,
			&application.ApplicationId,
			&application.UserName,
			&application.AppliedAt,
			&application.FirstName,
			&application.LastName,
			&application.Event_name,
			&application.Proffession,
			&application.College,
			&application.Company,
			&application.About,
			&application.Email,
			&application.Github,
			&application.Linkedin,
			&application.ApplicationStatusId,
		)

		fmt.Println("username", application.UserName)

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
func (c *eventRepository) RejectApplication(applicationStatusId int,eventname string) error {
	var eventName string
	var userName string

	query := `SELECT app.event_name,status.pending
				FROM application_forms AS app INNER JOIN appllication_statuses AS status 
				ON app.user_name = status.pending WHERE status.application_status_id = $1;`
	err := c.db.QueryRow(query, applicationStatusId).Scan(&eventName, &userName)
	if err != nil {
		return err
	}

	query2 := `UPDATE appllication_statuses SET pending = null, rejected = $1 WHERE application_status_id = $2;`
	err = c.db.QueryRow(query2, userName, applicationStatusId).Err()
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
func (c *eventRepository) DeletePoster(name string, eventid int) error {
	var id int
	query := `DELETE FROM posters WHERE name = $1 AND event_id = $2 RETURNING poster_id;`

	err := c.db.QueryRow(query, name, eventid).Scan(&id)
	fmt.Println("id deleted:", id)
	if err != nil {
		return err
	}
	return nil
}

// FindPoster implements interfaces.EventRepository
func (c *eventRepository) FindPoster(title string, eventid int) (domain.PosterResponse, error) {
	var poster domain.PosterResponse

	query := `SELECT poster_id,
				name,
				image,
				discription,
				date,
				colour FROM posters
				WHERE name  = $1 AND event_id = $2;`

	err := c.db.QueryRow(query, title, eventid).Scan(
		&poster.PosterId,
		&poster.Name,
		&poster.Image,
		&poster.Discription,
		&poster.Date,
		&poster.Colour,
	)

	fmt.Println("poster from find poster :", poster)
	return poster, err
}

// PostersByEvent implements interfaces.EventRepository
func (c *eventRepository) PostersByEvent(eventid int) ([]domain.PosterResponse, error) {
	fmt.Println("all posters called from repo")
	var posters []domain.PosterResponse

	query := `SELECT 
				COUNT(*) OVER(),
				poster_id,
				name,
				image,
				discription,
				date,
				colour FROM posters
				WHERE event_id = $1
				ORDER BY date DESC;`

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
			organizer_name,
			event_pic,
			short_discription,
			long_discription,
			event_date,
			location,
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
			AND concat(event_id::text, title, organizer_name, short_discription, long_discription, location) LIKE '%' || $2 || '%' 
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
			&event.OrganizerName,
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
func (c *eventRepository) DeleteEvent(title string) error {

	var id int
	query := `DELETE FROM events WHERE title = $1 RETURNING event_id;`

	err := c.db.QueryRow(query, title).Scan(&id)
	fmt.Println("id deleted:", id)
	if err != nil {
		return err
	}
	return nil

}

// UpdateEvent implements interfaces.EventRepository
func (c *eventRepository) UpdateEvent(event domain.Events, title string) (int, error) {
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
								website_link = $15 WHERE title = $16;`

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
		event.WebsiteLink, title).Err()

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
					organizer_name,
					event_pic,
					short_discription,
					long_discription,
					event_date,
					location,
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
					website_link FROM events 
					WHERE event_date > $1 AND approved = true 
					AND cusat_only = $2 AND sex = $3 AND online = $4 ORDER BY event_date DESC  
					LIMIT $5 OFFSET $6;`
	query2 := `SELECT 
					COUNT(*) OVER(),
					event_id,
					title,
					organizer_name,
					event_pic,
					short_discription,
					long_discription,
					event_date,
					location,
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
					website_link FROM events 
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
			&event.OrganizerName,
			&event.EventPic,
			&event.ShortDiscription,
			&event.LongDiscription,
			&event.EventDate,
			&event.Location,
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
func (c *eventRepository) FindEvent(title string) (domain.EventResponse, error) {
	var event domain.EventResponse

	query := `SELECT event_id,
	title,
	organizer_name,
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
	website_link FROM events 
	WHERE title = $1;`

	err := c.db.QueryRow(query, title).Scan(

		&event.EventId,
		&event.Title,
		&event.OrganizerName,
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

	fmt.Println("event from find evnet :", event)
	return event, err
}

// InsertUser implements interfaces.UserRepository
func (c *eventRepository) CreateEvent(event domain.Events) (int, error) {
	var id int

	userName := event.OrganizerName

	query := `INSERT INTO events(title,organizer_name,
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
								website_link,application_left)VALUES($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11, $12, $13,$14,$15, $16, $17, $18, $19,$20)
								RETURNING event_id;`

	err := c.db.QueryRow(query, event.Title,
		userName,
		event.EventPic,
		event.ShortDiscription,
		event.LongDiscription,
		event.EventDate,
		event.Location,
		event.CreatedAt,
		event.Approved,
		event.Paid,
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
