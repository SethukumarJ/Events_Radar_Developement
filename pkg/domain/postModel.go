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

type Bios struct {
	BioId         uint   `json:"bioid" gorm:"autoIncrement:true;unique"`
	UserName      string `json:"username" validate:"required,min=2,max=50"`
	About         string `json:"about"`
	Twitter       string `json:"twitterlink"`
	Github        string `json:"githublink"`
	LinkedIn      string `json:"linkedin"`
	Skills        string `json:"skills"`
	Qualification string `json:"qualification"`
	DevFolio      string `json:"devfolio"`
	WebsiteLink   string `json:"websitelink"`
}

type Verification struct {
	Email string `json:"email" validate:"email"`
	Code  string `json:"code"`
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
	Online                 bool      `json:"online" gorm:"default:false"`
	MaxApplications        int       `json:"maxapplications"`
	ApplicationClosingDate string    `json:"applicationclosingdate"`
	ApplicationLink        string    `json:"applicationlink"`
	WebsiteLink            string    `json:"websitelink"`
}

type Faqas struct {
	FaqaId        uint      `json:"faqaid" gorm:"autoIncrement:true;unique"`
	Question      string    `json:"question" validate:"required,min=2,max=50"`
	AnswerId      int       `json:"answerid" gorm:"default:0"`
	Title         string    `json:"title"`
	CreatedAt     time.Time `json:"createdat"`
	UserName      string    `json:"username"`
	OrganizerName string    `json:"organizername"`
	Public        bool      `json:"public" gorm:"default:false"`
}

type Answers struct {
	AnswerId uint   `json:"answerid" gorm:"autoIncrement:true;unique"`
	Answer   string `json:"answer"`
}

type Organizations struct {
	OrganizationId   uint      `json:"organizationid" gorm:"autoIncrement:true;unique"`
	OrganizationName string    `json:"organizationname" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	CreatedBy        string    `json:"createdby"`
	Logo             string    `json:"logo"`
	About            string    `json:"about"`
	CreatedAt        time.Time `json:"createdat"`
	LinkedIn         string    `json:"linkedin"`
	WebsiteLink      string    `json:"websitelink"`
	Verified         bool      `json:"verified" gorm:"default:false"`
}

type Org_Status struct {
	OrgStatusId uint   `json:"orgstatusid" gorm:"autoIncrement:true;unique"`
	Registered  string `json:"registered"`
	Pending     string `json:"pending"`
	Rejected    string `json:"renected"`
}

type Join_Status struct {
	JoinStatusId uint   `json:"orgstatusid" gorm:"autoIncrement:true;unique"`
	Joined       string `json:"joined"`
	Pending      string `json:"pending"`
	Rejected     string `json:"renected"`
}

type User_Organization_Connections struct {
	UserOrganizationConnectionsId uint          `json:"organizationid" gorm:"autoIncrement:true;unique"`
	OrganizationName              string        `json:"organizationname"`
	Organizations                 Organizations `gorm:"foreignKey:OrganizationName;references:OrganizationName"`
	UserName                      string        `json:"username"`
	Users                         Users         `gorm:"foreignKey:UserName;references:UserName"`
	Role                          string        `json:"role" gorm:"not null"`
}
