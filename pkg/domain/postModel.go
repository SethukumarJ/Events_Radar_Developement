package domain

import (
	"time"
)

type Users struct {
	UserId       uint   `json:"user_id" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	UserName     string `json:"user_name" gorm:"unique" validate:"required,min=2,max=50"`
	FirstName    string `json:"first_name" validate:"required,min=2,max=50"`
	LastName     string `json:"last_name" validate:"required,min=1,max=50"`
	Email        string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Password     string `json:"password" validate:"required,min=,max=20"`
	Verification bool   `json:"verification" gorm:"default:false"`
	Vip          bool   `json:"vip" gorm:"default:false" swaggerignore:"true"`
	PhoneNumber  string `json:"phonenumber" validate:"required,min=10,max=20,numeric"`
	Profile      string `json:"profile" validate:"omitempty,max=255,url"`
	EventId      uint   `json:"event_id" swaggerignore:"true"`
}

type Login struct {
	Email    string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type Bios struct {
	BioId         uint   `json:"bio_id" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	UserId        int    `json:"user_id" validate:"required,min=2,max=50"`
	Users         Users  `gorm:"foreignKey:UserId;references:UserId"  swaggerignore:"true"`
	UserName      string `json:"user_name"`
	About         string `json:"about"`
	Twitter       string `json:"twitter_link" validate:"omitempty,max=255,url"`
	Github        string `json:"github_link" validate:"omitempty,max=255,url"`
	LinkedIn      string `json:"linked_in" validate:"omitempty,max=255,url"`
	Skills        string `json:"skills"`
	Qualification string `json:"qualification"`
	DevFolio      string `json:"devfolio" validate:"omitempty,max=255,url"`
	WebsiteLink   string `json:"website_link" validate:"omitempty,max=255,url"`
}

type Verification struct {
	Email string `json:"email" validate:"email"`
	Code  string `json:"code"`
}

type Admins struct {
	AdminId      uint   `json:"admin_id" gorm:"primary key;autoIncrement:true;unique"`
	AdminName    string `json:"admin_name" gorm:"unique" validate:"required,min=2,max=50"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" gorm:"notnull;unique" validate:"email,required"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=10,max=20,numeric"`
	Verification bool   `json:"verification" gorm:"default:false" swaggerignore:"true"`
}

type Events struct {
	EventId                uint          `json:"event_id" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	Title                  string        `json:"title" gorm:"unique" validate:"required,min=2,max=50"`
	OrganizationId         int           `json:"organization_id"`
	Organizations          Organizations `gorm:"foreignKey:OrganizationId;references:OrganizationId"`
	User_id                int           `json:"user_id"`
	Users                  Users         `gorm:"foreignKey:UserId;references:UserId"  swaggerignore:"true"`
	CreatedBy              string        `json:"created_by`
	EventPic               string        `json:"event_pic" validate:"required,url"`
	ShortDiscription       string        `json:"short_discription"`
	LongDiscription        string        `json:"long_discription"`
	EventDate              string        `json:"event_date" validate:"required"`
	Location               string        `json:"location"`
	CreatedAt              time.Time     `json:"created_at" swaggerignore:"true"`
	Approved               bool          `json:"approved" gorm:"default:false" swaggerignore:"true"`
	Paid                   bool          `json:"paid" gorm:"default:false"`
	Amount                 string        `json:"amount" validate:"numeric"`
	Sex                    string        `json:"sex" default:"any"`
	CusatOnly              bool          `json:"cusat_only" gorm:"default:false"`
	Archived               bool          `json:"archived" swaggerignore:"true"`
	SubEvents              string        `json:"sub_events"`
	Online                 bool          `json:"online" gorm:"default:false"`
	MaxApplications        int           `json:"max_applications"`
	ApplicationClosingDate string        `json:"application_closing_date"`
	ApplicationLink        string        `json:"application_link" validate:"omitempty,max=255,url"`
	WebsiteLink            string        `json:"website_link" validate:"omitempty,max=255,url"`
	ApplicationLeft        int           `json:"application_left" validate:"numeric"`
	Featured               bool          `json:"featured" gorm:"default:false" swaggerignore:"true"`
}

type Posters struct {
	PosterId    uint   `json:"poster_id" ggorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	Name        string `json:"name"`
	Events      Events `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId     int    `json:"event_id" swaggerignore:"true" `
	Image       string `json:"image" validate:"omitempty,max=255,url"`
	Discription string `json:"discription"`
	Date        string `json:"date" validate:"required" swaggerignore:"true" `
	Colour      string `json:"colour"`
}

type Faqas struct {
	FaqaId         uint          `json:"faqaid" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	Question       string        `json:"question" validate:"required,min=2,max=50"`
	AnswerId       int           `json:"answer_id" gorm:"default:0" swaggerignore:"true"`
	Answers        Answers       `gorm:"foreignKey:AnswerId;references:AnswerId" swaggerignore:"true"`
	Events         Events        `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId        int           `json:"event_id" swaggerignore:"true" `
	CreatedAt      time.Time     `json:"createdat" swaggerignore:"true"`
	UserId         int           `json:"user_id" validate:"required,min=2,max=50" swaggerignore:"true"`
	Users          Users         `gorm:"foreignKey:UserId;references:UserId" swaggerignore:"true"`
	OrganizationId int           `json:"organization_id" validate:"required" swaggerignore:"true"`
	Organizations  Organizations `gorm:"foreignKey:OrganizationId;references:OrganizationId" swaggerignore:"true"`
	Public         bool          `json:"public" gorm:"default:false" swaggerignore:"true"`
}

type Answers struct {
	AnswerId uint   `json:"answerid" gorm:"autoIncrement:true;unique" swaggerignore:"true"`
	Answer   string `json:"answer" validate:"required,min=2,max=255"`
}

type Organizations struct {
	OrganizationId   uint      `json:"organization_id" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	OrganizationName string    `json:"organization_name" gorm:"unique" validate:"required,min=2,max=50"`
	CreatedBy        int       `json:"created_by" swaggerignore:"true"`
	Logo             string    `json:"logo"`
	About            string    `json:"about"`
	CreatedAt        time.Time `json:"created_at"`
	LinkedIn         string    `json:"linkedin"`
	WebsiteLink      string    `json:"website_link"`
	Verified         bool      `json:"verified" gorm:"default:false" swaggerignore:"true"`
}

type Org_Status struct {
	OrgStatusId uint `json:"org_status_id" gorm:"primary key;autoIncrement:true;unique"`
	Registered  int  `json:"registered"`
	Pending     int  `json:"pending"`
	Rejected    int  `json:"rejected"`
}

type Join_Status struct {
	JoinStatusId   uint          `json:"join_status_id" gorm:"primary key;autoIncrement:true;unique"`
	OrganizationId int           `json:"organization_id" validate:"required" swaggerignore:"true"`
	Organizations  Organizations `gorm:"foreignKey:OrganizationId;references:OrganizationId" swaggerignore:"true"`
	Joined         string        `json:"joined"`
	Pending        string        `json:"pending"`
	Rejected       string        `json:"rejected"`
}

type Appllication_Statuses struct {
	ApplicationStatusId uint   `json:"application_status_id" gorm:"primary key;autoIncrement:true;unique"`
	Events              Events `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId             int    `json:"event_id" swaggerignore:"true" `
	Accepted            string `json:"accepted"`
	Pending             string `json:"pending"`
	Rejected            string `json:"rejected"`
}

type User_Organization_Connections struct {
	UserOrganizationConnectionsId uint          `json:"organizationid" gorm:"primary key;autoIncrement:true;unique"`
	OrganizationId                int           `json:"organization_id" validate:"required" swaggerignore:"true"`
	Organizations                 Organizations `gorm:"foreignKey:OrganizationId;references:OrganizationId" swaggerignore:"true"`
	UserId                        int           `json:"user_id" validate:"required,min=2,max=50" swaggerignore:"true"`
	Users                         Users         `gorm:"foreignKey:UserId;references:UserId" swaggerignore:"true"`
	Role                          string        `json:"role" gorm:"not null"`
	JoinedAt                      time.Time     `json:"joinedat"`
}

type Notificaiton struct {
	NotificaitonId uint          `json:"notification_id" gorm:"primary key;autoIncrement:true;unique"`
	UserId         int           `json:"user_id" validate:"required,min=2,max=50" swaggerignore:"true"`
	Users          Users         `gorm:"foreignKey:UserId;references:UserId" swaggerignore:"true"`
	OrganizationId int           `json:"organization_id" validate:"required" swaggerignore:"true"`
	Organizations  Organizations `gorm:"foreignKey:OrganizationId;references:OrganizationId" swaggerignore:"true"`
	Events         Events        `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId        int           `json:"event_id" swaggerignore:"true" `
	Message        string        `json:"message"`
	Time           time.Time     `json:"time"`
}

type ApplicationForm struct {
	ApplicationId uint      `json:"applicationid" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true" `
	UserId        int       `json:"user_id" validate:"required,min=2,max=50" swaggerignore:"true"`
	Users         Users     `gorm:"foreignKey:UserId;references:UserId" swaggerignore:"true"`
	AppliedAt     time.Time `json:"appliedat" swaggerignore:"true"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	Events        Events    `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId       int       `json:"event_id" swaggerignore:"true" `
	Proffession   string    `json:"proffession"`
	College       string    `json:"college"`
	Company       string    `json:"company"`
	About         string    `json:"about"`
	Email         string    `json:"email" gorm:"unique" validate:"email,required"`
	Github        string    `json:"github" validate:"omitempty,max=255,url"`
	Linkedin      string    `json:"linkedin" validate:"omitempty,max=255,url"`
}

type PageVariables struct {
	OrderId string
	Email   string
	Name    string
	Amount  string
	Contact string
}

type Promotion struct {
	PromotionId   uint          `json:"promotion_id" gorm:"autoIncrement:true;unique"`
	Events        Events        `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId       int           `json:"event_id" swaggerignore:"true" `
	OrderId       string        `json:"order_id"`
	PromotedBy    int           `json:"organization_id" validate:"required" swaggerignore:"true"`
	Organizations Organizations `gorm:"foreignKey:PromotedBy;references:OrganizationId" swaggerignore:"true"`
	PaymentId     string        `json:"payment_id"`
	Amount        string        `json:"amount"`
	Plan          string        `json:"plan"`
	Status        bool          `json:"status" gorm:"default:false"`
}

type Packages struct {
	PackagesId uint   `json:"packagesid" gorm:"primary key;autoIncrement:true;unique" swaggerignore:"true"`
	Events     Events `gorm:"foreignKey:EventId;references:EventId" swaggerignore:"true"`
	EventId    int    `json:"event_id" swaggerignore:"true" `
	Basic      bool   `json:"basic" gorm:"default: false"`
	Standard   bool   `json:"standard" gorm:"default: false"`
	Premium    bool   `json:"premium" gorm:"default: false"`
}

type AddMembers struct {
	Members string
}
