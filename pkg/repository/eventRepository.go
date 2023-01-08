package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
)

type eventRepository struct {
	db *sql.DB
}

// AllApprovedEvents implements interfaces.EventRepository
func (c *eventRepository) AllApprovedEvents(pagenation utils.Filter) ([]domain.EventResponse, utils.Metadata, error) {
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
					website_link FROM events WHERE event_date > $1 AND approved = true
					LIMIT $2 OFFSET $3;`

	rows, err := c.db.Query(query, dateString,pagenation.Limit(), pagenation.Offset())
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

func NewEventRepository(db *sql.DB) interfaces.EventRepository {
	return &eventRepository{
		db: db,
	}
}

// FindUser implements interfaces.UserRepository
func (c *eventRepository) FindEvent(email string) (domain.EventResponse, error) {

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

	err := c.db.QueryRow(query, email).Scan(
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

	fmt.Println("event from find event :", event)
	return event, err
}

// InsertUser implements interfaces.UserRepository
func (c *eventRepository) CreateEvent(event domain.Events) (int, error) {
	var id int

	query := `INSERT INTO events(
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
								website_link)VALUES($1, $2, $3, $4, $5, $6,$7,$8, $9, $10, $11, $12, $13,$14,$15, $16, $17, $18, $19)
								RETURNING event_id;`

	err := c.db.QueryRow(query,
		event.Title,
		event.OrganizerName,
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
		event.WebsiteLink).Scan(&id)

	fmt.Println("id", id)
	return id, err
}
