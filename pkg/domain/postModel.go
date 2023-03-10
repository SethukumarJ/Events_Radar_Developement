package domain

import (
	"time"
)

type Users struct {
	UserId       uint   `json:"userid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	UserName     string `json:"username" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	FirstName    string `json:"firstname" validate:"required,min=2,max=50"`
	LastName     string `json:"lastname" validate:"required,min=1,max=50"`
	Email        string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Password     string `json:"password"`
	Verification bool   `json:"verification" gorm:"default:false"`
	Vip          bool   `json:"vip" gorm:"default:false" swaggerignore:"true"`
	PhoneNumber  string `json:"phonenumber"`
	Profile      string `json:"profile"`
	EventId      uint   `json:"eventid" swaggerignore:"true"`
}

type Login struct {
	Email    string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type Bios struct {
	BioId         uint   `json:"bioid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
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
	Verification bool   `json:"verification" gorm:"default:false" swaggerignore:"true"`
}

type Events struct {
	EventId                uint      `json:"eventid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	Title                  string    `json:"title" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	OrganizerName          string    `json:"organizername" validate:"required"`
	EventPic               string    `json:"eventpic" validate:"required"`
	ShortDiscription       string    `json:"shortdiscription"`
	LongDiscription        string    `json:"longdiscription"`
	EventDate              string    `json:"eventdate" validate:"required"`
	Location               string    `json:"location"`
	CreatedAt              time.Time `json:"createdat" swaggerignore:"true"`
	Approved               bool      `json:"approved" gorm:"default:false" swaggerignore:"true"`
	Paid                   bool      `json:"paid" gorm:"default:false"`
	Sex                    string    `json:"sex" default:"any"`
	CusatOnly              bool      `json:"cusatonly" gorm:"default:false"`
	Archived               bool      `json:"archived" swaggerignore:"true"`
	SubEvents              string    `json:"subevents"`
	Online                 bool      `json:"online" gorm:"default:false"`
	MaxApplications        int       `json:"maxapplications"`
	ApplicationClosingDate string    `json:"applicationclosingdate"`
	ApplicationLink        string    `json:"applicationlink"`
	WebsiteLink            string    `json:"websitelink"`
	ApplicationLeft        int       `json:"applicationleft"`
	Featured               bool      `json:"featred" gorm:"default:false" swaggerignore:"true"`
}

type Posters struct {
	PosterId    uint   `json:"posterid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	Name        string `json:"name"`
	Events      Events `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId     int    `json:"eventid" swaggerignore:"true" `
	Image       string `json:"image"`
	Discription string `json:"discription"`
	Date        string `json:"date" swaggerignore:"true" `
	Colour      string `json:"colour"`
}

type Faqas struct {
	FaqaId        uint      `json:"faqaid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	Question      string    `json:"question" validate:"required,min=2,max=50"`
	AnswerId      int       `json:"answerid" gorm:"default:0" swaggerignore:"true"`
	Title         string    `json:"title" swaggerignore:"true"`
	CreatedAt     time.Time `json:"createdat" swaggerignore:"true"`
	UserName      string    `json:"username" swaggerignore:"true"`
	OrganizerName string    `json:"organizername" swaggerignore:"true"`
	Public        bool      `json:"public" gorm:"default:false" swaggerignore:"true"`
}

type Answers struct {
	AnswerId uint   `json:"answerid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	Answer   string `json:"answer"`
}

type Organizations struct {
	OrganizationId   uint      `json:"organizationid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	OrganizationName string    `json:"organization_name" gorm:"primary key;unique" validate:"required,min=2,max=50"`
	CreatedBy        string    `json:"createdby" swaggerignore:"true"`
	Logo             string    `json:"logo"`
	About            string    `json:"about"`
	CreatedAt        time.Time `json:"createdat"`
	LinkedIn         string    `json:"linkedin"`
	WebsiteLink      string    `json:"websitelink"`
	Verified         bool      `json:"verified" gorm:"default:false" swaggerignore:"true"`
}

type Org_Status struct {
	OrgStatusId uint   `json:"orgstatusid" gorm:"autoIncrement:true;unique"`
	Registered  string `json:"registered"`
	Pending     string `json:"pending"`
	Rejected    string `json:"rejected"`
}

type Join_Status struct {
	JoinStatusId     uint   `json:"orgstatusid" gorm:"autoIncrement:true;unique"`
	OrganizationName string `json:"organizationname"`
	Joined           string `json:"joined"`
	Pending          string `json:"pending"`
	Rejected         string `json:"rejected"`
}

type Appllication_Statuses struct {
	ApplicationStatusId uint   `json:"applicationstatusid" gorm:"autoIncrement:true;unique"`
	EventName           string `json:"eventname"`
	Accepted            string `json:"accepted"`
	Pending             string `json:"pending"`
	Rejected            string `json:"rejected"`
}

type User_Organization_Connections struct {
	UserOrganizationConnectionsId uint          `json:"organizationid" gorm:"autoIncrement:true;unique"`
	OrganizationName              string        `json:"organizationname"`
	Organizations                 Organizations `gorm:"foreignKey:OrganizationName;references:OrganizationName"`
	UserName                      string        `json:"username"`
	Users                         Users         `gorm:"foreignKey:UserName;references:UserName"`
	Role                          string        `json:"role" gorm:"not null"`
	JoinedAt                      time.Time     `json:"joinedat"`
}

type Notificaiton struct {
	NotificaitonId   uint      `json:"notification_id" gorm:"autoIncrement:true;unique"`
	UserName         string    `json:"username"`
	OrganizationName string    `json:"organizationname"`
	Event_Title      string    `json:"eventtitle"`
	Message          string    `json:"message"`
	Time             time.Time `json:"time"`
}

type ApplicationForm struct {
	ApplicationId       uint      `json:"applicationid" gorm:"autoIncrement:true;unique" swaggerignore:"true" `
	UserName            string    `json:"username" swaggerignore:"true"`
	AppliedAt           time.Time `json:"appliedat" swaggerignore:"true"`
	FirstName           string    `json:"firstname"`
	LastName            string    `json:"lastname"`
	Event_name          string    `json:"event_name" swaggerignore:"true"`
	Proffession         string    `json:"proffession"`
	College             string    `json:"college"`
	Company             string    `json:"company"`
	About               string    `json:"about"`
	Email               string    `json:"email"`
	Github              string    `json:"github"`
	Linkedin            string    `json:"linkedin"`
}

type PageVariables struct {
	OrderId string
	Email   string
	Name    string
	Amount  string
	Contact string
}

type Promotion struct {
	PromotionId uint   `json:"promotionid" gorm:"autoIncrement:true;unique"`
	EventTitle  string `json:"eventtitle"`
	OrderId     string `json:"orderid"`
	PromotedBy  string `json:"organizationname"`
	PaymentId   string `json:"paymentid"`
	Amount      string `json:"amount"`
	Plan        string `json:"plan"`
	Status      bool   `json:"status" gorm:"default:false"`
}

type Packages struct {
	PackagesId uint   `json:"packagesid" gorm:"autoIncrement:true;unique"`
	EventTitle string `json:"eventtitle"`
	Basic      bool   `json:"basic" gorm:"default: false"`
	Standard   bool   `json:"standard" gorm:"default: false"`
	Premium    bool   `json:"premium" gorm:"default: false"`
}

type AddMembers struct {

	Members string
}