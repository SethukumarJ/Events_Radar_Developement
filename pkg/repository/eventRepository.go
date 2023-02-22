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

	
	rows, err := c.db.Query(query, dateString,filter.CusatOnly,filter.Sex,filter.Online, pagenation.Limit(), pagenation.Offset())
	fmt.Println("rows", rows)
	if err != nil {
		return nil, utils.Metadata{}, err
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
