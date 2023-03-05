package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/SethukumarJ/Events_Radar_Developement/cmd/api/docs"
	handler "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/handler"
	middleware "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/middleware"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/cors"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handler.UserHandler,
	authHandler handler.AuthHandler,
	adminHandler handler.AdminHandler,
	eventHandler handler.EventHandler,
	middleware middleware.Middleware) *ServerHTTP {
	authHandler.InitializeOAuthGoogle()

	engine := gin.Default()
	// Enable CORS for all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowOrigins = []string{"https://eventsradar.online"}
	config.AllowMethods = []string{"GET", "POST","PUT","PATCH","DELETE"}
	config.AllowHeaders = []string{"Authorization"}
	engine.Use(cors.New(config))
	engine.HTMLRender = gintemplate.Default()
	engine.GET("/r", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "My Webpage",
		})
	})

	// Use logger from Gin
	engine.Use(gin.Logger())
	// engine.Static("/templates", "./api/templates")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	engine.GET("/payment-success", userHandler.PaymentSuccess)
	engine.GET("/payment-faliure", userHandler.PaymentFaliure)
	engine.GET("/organization/event/promote", userHandler.Pay)
	user := engine.Group("user")
	{
		user.GET("/sso-google", authHandler.GoogleAuth)
		user.GET("/callback-gl", authHandler.CallBackFromGoogle)
		user.POST("/signup", authHandler.UserSignup)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send-verification", userHandler.SendVerificationMail)
		user.GET("/verify-account", authHandler.VerifyAccount)
		user.GET("/list-faqas", userHandler.GetPublicFaqas)
		user.GET("/list-organizations", userHandler.ListOrganizations)
		user.GET("/list-approved-events", eventHandler.ViewAllApprovedEvents)
		user.GET("/geteventbytitle", eventHandler.GetEventByTitle)
		user.PATCH("/update-password", userHandler.UpdatePassword)
		user.GET("/search-event", eventHandler.SearchEventUser)
		user.GET("/accept-invitation", userHandler.AcceptJoinInvitation)
		user.Use(middleware.AuthorizeJwt())
		{
			user.POST("/token-refresh", authHandler.UserRefreshToken)
			user.POST("/apply-event", userHandler.ApplyEvent)
			user.POST("/event/post-question", userHandler.PostQuestion)
			user.PATCH("/update-profile", userHandler.UpdateProfile)
			user.POST("/create-organization", userHandler.CreateOrganization)
			user.POST("/create-event", eventHandler.CreateEventUser)
			user.PATCH("/join-organization", userHandler.JoinOrganization)

		}

	}

	//Admin routes
	admin := engine.Group("admin")
	{
		admin.POST("/signup", authHandler.AdminSignup)
		admin.POST("/login", authHandler.AdminLogin)

		admin.Use(middleware.AuthorizeJwt())
		{
			admin.GET("/token-refresh", authHandler.AdminRefreshToken)
			admin.PATCH("/approve-event", adminHandler.ApproveEvent)
			admin.POST("/create-event", eventHandler.CreateEventAdmin)
			admin.GET("/list-users", adminHandler.ViewAllUsers)
			admin.GET("/list-events", adminHandler.ViewAllEvents)
			admin.PATCH("/register-organization", adminHandler.RegisterOrganization)
			admin.PATCH("/reject-organization", adminHandler.RejectOrganization)
			admin.GET("/list-organizations", adminHandler.ListOrgRequests)
			admin.PATCH("/vipfy-user", adminHandler.VipUser)
			admin.GET("/search-event", adminHandler.SearchEvent)
		}
	}

	//Organization routes
	organization := engine.Group("organization")
	{
		organization.GET("/get-organization", userHandler.GetOrganization)
		organization.GET("/list-organizations", userHandler.ListOrganizations)
		organization.GET("/event/get-posters", eventHandler.PostersByEvent)
		organization.GET("/event/get-posterbytitle", eventHandler.GetPosterByTitle)
		organization.Use(middleware.AuthorizeOrg())
		{
			organization.POST("/event/create-poster", eventHandler.CreatePosterOrganization)
			organization.DELETE("/event/delete-poster", eventHandler.DeletePoster)
			organization.PATCH("/event/update-event", eventHandler.UpdateEvent)
			organization.DELETE("/event/delete-event", eventHandler.DeleteEvent)
			organization.GET("/event/list-questions", userHandler.GetQuestions)
			organization.POST("/event/post-answer", userHandler.PostAnswer)
			organization.POST("/create-event", eventHandler.CreateEventOrganization)
			organization.POST("/admin/add-members", userHandler.AddMembers)
			organization.GET("/admin/list-members")
			organization.PATCH("/admin/admit-member", userHandler.AdmitMember)
			organization.GET("/join-requests", userHandler.ListJoinRequests)
			organization.PATCH("/event/accept-application", eventHandler.AcceptApplication)
			organization.PATCH("/reject-application", eventHandler.RejectApplication)
			organization.GET("/event/list-applications", eventHandler.ListApplications)

		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
