package domain

import "time"

type UserResponse struct {
	UserId       uint   `json:"user_id"`
	UserName     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification" `
	Vip          bool   `json:"vip" `
	PhoneNumber  string `json:"phone_number"`
	Profile      string `json:"profile"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AdminResponse struct {
	AdminId      uint   `json:"admin_id"`
	AdminName    string `json:"user_name"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification"`
	PhoneNumber  string `json:"phone_number"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type EventResponse struct {
	EventId                uint      `json:"eventid"`
	Title                  string    `json:"title"`
	OrganizationId         int       `json:"organization_id" `
	User_id                int       `json:"user_id"`
	CreatedBy              string    `json:"created_by"`
	EventPic               string    `json:"event_pic" `
	ShortDiscription       string    `json:"short_discription"`
	LongDiscription        string    `json:"long_discription"`
	EventDate              string    `json:"event_date"`
	Location               string    `json:"location"`
	CreatedAt              time.Time `json:"created_at"`
	Approved               bool      `json:"approved"`
	Paid                   bool      `json:"paid" `
	Amount                 string    `json:"amount"`
	Sex                    string    `json:"sex" `
	CusatOnly              bool      `json:"cusat_only"`
	Archived               bool      `json:"archived"`
	SubEvents              string    `json:"subevents"`
	Online                 bool      `json:"online"`
	MaxApplications        int       `json:"max_applications"`
	ApplicationLeft        int       `json:"application_left"`
	ApplicationClosingDate string    `json:"application_closing_date"`
	ApplicationLink        string    `json:"application_link"`
	WebsiteLink            string    `json:"website_link"`
}
type PosterResponse struct {
	PosterId    uint   `json:"poster_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Discription string `json:"discription"`
	Date        string `json:"date"`
	Colour      string `json:"colour"`
	EventId     int    `json:"event_id"`
}

type FaqaResponse struct {
	FaqaId          uint      `json:"faqa_id"`
	Question        string    `json:"question"`
	AnswerId        int       `json:"answer_id"`
	EventId         int       `json:"event_id"`
	CreatedAt       time.Time `json:"created_at"`
	UserName        string    `json:"user_name"`
	OrganizaitionId int       `json:"organization_id"`
}

type QAResponse struct {
	FaqaId          uint      `json:"faqa_id"`
	Question        string    `json:"question"`
	AnswerId        int       `json:"answer_id"`
	EventId         int       `json:"event_id"`
	CreatedAt       time.Time `json:"created_at"`
	UserName        string    `json:"user_name"`
	OrganizaitionId int       `json:"organization_id"`
	Answer          string    `json:"answer"`
}

type OrganizationsResponse struct {
	OrganizationId   uint      `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	CreatedBy        int       `json:"created_by"`
	Logo             string    `json:"logo"`
	About            string    `json:"about"`
	CreatedAt        time.Time `json:"created_at"`
	LinkedIn         string    `json:"linkedin"`
	WebsiteLink      string    `json:"website_link"`
	Verified         bool      `json:"verified"`
	OrgStatusId      int       `json:"org_status_id"`
}

type UserOrganizationConnectionResponse struct {
	OrganizationId int    `json:"organization_id"`
	UserId         int    `json:"user_id"`
	Role           string `json:"role"`
}

type Join_StatusResponse struct {
	JoinStatusId   uint   `json:"join_status_id"`
	OrganizationId int    `json:"organization_id"`
	Joined         string `json:"joined"`
	Pending        string `json:"pending"`
	Rejected       string `json:"rejected"`
}

type ApplicationFormResponse struct {
	ApplicationId       uint      `json:"applicationid"`
	UserId              string    `json:"user_id"`
	AppliedAt           time.Time `json:"applied_at"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	EventId             int       `json:"event_id"`
	Proffession         string    `json:"proffession"`
	College             string    `json:"college"`
	Company             string    `json:"company"`
	About               string    `json:"about"`
	Email               string    `json:"email"`
	Github              string    `json:"github"`
	Linkedin            string    `json:"linkedin"`
	ApplicationStatusId int       `json:"application_status_id"`
}

type PromotionResponse struct {
	PromotionId    uint   `json:"promotion_id"`
	EventId        int    `json:"event_id"`
	OrderId        string `json:"order_id"`
	PromotedBy     string `json:"promoted_by"`
	OrganizationId int    `json:"organization_id"`
	PaymentId      string `json:"payment_id"`
	Amount         string `json:"amount"`
	Plan           string `json:"plan"`
}
