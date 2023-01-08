package repository

import (
	"database/sql"
	"fmt"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
)

type eventRepository struct {
	db *sql.DB
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
						event.WebsiteLink,).Scan(&id)

	fmt.Println("id", id)
	return id, err
}
