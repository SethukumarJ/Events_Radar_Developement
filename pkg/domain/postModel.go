package domain

import "time"

type Users struct {
	UserId       uint   `json:"userid" gorm:"autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	FirstName    string `json:"firstname" validate:"required,min=2,max=50"`
	LastName     string `json:"lastname" validate:"required,min=1,max=50"`
	Email        string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Password     string `json:"password" validate:"required"`
	Verification bool   `json:"verification" gorm:"default:false"`
	Vip          bool   `json:"vip" gorm:"default:false"`
	PhoneNumber  string `json:"phonenumber" gorm:"unique"`
	Profile      string `json:"profile"`
	EventId      uint   `json:"eventid"`
}

type Verification struct {
	Email string `json:"email" validate:"email"`
	Code  int    `json:"code"`
}

type Admins struct {
	AdminId      uint   `json:"adminid" gorm:"autoIncrement:true;unique"`
	AdminName    string `json:"adminname" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	PhoneNumber  string `json:"phonenumber" gorm:"unique"`
	Verification bool   `json:"verification" gorm:"default:false"`
}

type Events struct {
	EventId                uint      `json:"eventid" gorm:"autoIncrement:true;unique"`
	Title                  string    `json:"title" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	OrganizerName          string    `json:"organizername" validate:"required"`
	EventPic               string    `json:"eventpic" validate:"required"`
	ShortDiscription       string    `json:"shortdiscription"`
	LongDiscription        string    `json:"longdiscription"`
	EventDate              string    `json:"eventdate" validate:"required"`
	Location               string    `json:"location"`
	CreatedAt              time.Time `json:"createdat"`
	Approved               bool      `json:"approved" gorm:"default:false"`
	Paid                   bool      `json:"paid" gorm:"default:false"`
	Sex                    string    `json:"sex" gorm:"default:any"`
	CusatOnly              bool      `json:"cusatonly" gorm:"default:false"`
	Archived               bool      `json:"archived"`
	SubEvents              string    `json:"subevents"`
	Online                 bool      `json:"offline" gorm:"default:false"`
	MaxApplications        int       `json:"maxapplications"`
	ApplicationClosingDate string    `json:"applicationclosingdate"`
	ApplicationLink        string    `json:"applicationlink"`
	WebsiteLink            string    `json:"websitelink"`
}
