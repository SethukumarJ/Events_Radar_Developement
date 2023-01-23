package domain

import "time"

type UserResponse struct {
	UserId       uint   `json:"userid"`
	UserName     string `json:"username"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification" `
	Vip          bool   `json:"vip" `
	PhoneNumber  string `json:"phonenumber"`
	Profile      string `json:"profile"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type AdminResponse struct {
	AdminId      uint   `json:"adminid"`
	AdminName    string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Verification bool   `json:"verification"`
	PhoneNumber  string `json:"phonenumber"`
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreshtoken"`
}

type EventResponse struct {
	EventId                uint      `json:"eventid"`
	Title                  string    `json:"title"`
	OrganizerName          string    `json:"organizername" `
	EventPic               string    `json:"eventpic" `
	ShortDiscription       string    `json:"shortdiscription"`
	LongDiscription        string    `json:"longdiscription"`
	EventDate              string    `json:"eventdate"`
	Location               string    `json:"location"`
	CreatedAt              time.Time `json:"createdat"`
	Approved               bool      `json:"approved"`
	Paid                   bool      `json:"paid" `
	Sex                    string    `json:"sex" `
	CusatOnly              bool      `json:"cusatonly"`
	Archived               bool      `json:"archived"`
	SubEvents              string    `json:"subevents"`
	Online                 bool      `json:"online"`
	MaxApplications        int       `json:"maxapplications"`
	ApplicationClosingDate string    `json:"applicationclosingdate"`
	ApplicationLink        string    `json:"applicationlink"`
	WebsiteLink            string    `json:"websitelink"`
}

type FaqaResponse struct {
	FaqaId        uint      `json:"faqaid"`
	Question      string    `json:"question"`
	AnswerId      int       `json:"answerid"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"createdat"`
	UserName      string    `json:"username"`
	OrganizerName string    `json:"organizername"`
}

type QAResponse struct {
	FaqaId        uint      `json:"faqaid"`
	Question      string    `json:"question"`
	AnswerId      int       `json:"answerid"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"createdat"`
	UserName      string    `json:"username"`
	OrganizerName string    `json:"organizername"`
	Answer        string    `json:"answer"`
}

type OrganizationsResponse struct {
	OrganizationId   uint      `json:"organizationid"`
	OrganizationName string    `json:"Oorganizationname"`
	CreatedBy        string    `json:"createdby"`
	Logo             string    `json:"logo"`
	About            string    `json:"about"`
	CreatedAt        time.Time `json:"createdat"`
	LinkedIn         string    `json:"linkedin"`
	WebsiteLink      string    `json:"websitelink"`
	Verified         bool      `json:"verified"`
	OrgStatusId      int       `json:"orgstatusid"`
}

type UserOrganizationConnectionResponse struct {
	OrganizationName string `json:"organizationname"`
	UserName         string `json:"username"`
	Role             string `json:"role" gorm:"not null"`
}


type Join_StatusResponse struct {
	JoinStatusId     uint   `json:"orgstatusid"`
	OrganizationName string `json:"organizationname"`
	Joined           string `json:"joined"`
	Pending          string `json:"pending"`
	Rejected         string `json:"rejected"`
}


type ApplicationFormResponse struct {
	ApplicationId uint      `json:"applicationid" gorm:"autoIncrement:true;unique"`
	UserName      string    `json:"username"`
	AppliedAt     time.Time `json:"appliedat"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Event_name    string    `json:"event_name"`
	Proffession   string    `json:"proffession"`
	College       string    `json:"college"`
	Company       string    `json:"company"`
	About         string    `json:"about"`
	Email         string    `json:"email"`
	Github        string    `json:"github"`
	Linkedin      string    `json:"linkedin"`
}
