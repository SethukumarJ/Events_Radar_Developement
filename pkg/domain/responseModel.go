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
