package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/thnkrn/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
	middleware "github.com/thnkrn/go-gin-clean-arch/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handler.UserHandler,
	authHandler handler.AuthHandler,
	adminHandler handler.AdminHandler,
	eventHandler handler.EventHandler,
	middleware middleware.Middleware) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))


	
	user := engine.Group("user")
	{
		user.POST("/signup", authHandler.UserSignup)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send/verification", userHandler.SendVerificationMail)
		user.PATCH("/verify/account",authHandler.VerifyAccount)
		user.PATCH("/password/update",userHandler.UpdatePassword)
		user.GET("/list/faqas",userHandler.GetPublicFaqas)
		user.GET("/list-organizations",userHandler.ListOrganizations)

		user.Use(middleware.AuthorizeJwt())
		{
			user.POST("/token/refresh", authHandler.UserRefreshToken)
			user.POST("/event/create", eventHandler.CreateEventUser)
			user.POST("/event/post/question", userHandler.PostQuestion)
			user.PATCH("/update/profile",userHandler.UpdateProfile)
			user.POST("/event/post/answer",userHandler.PostAnswer)
			user.GET("/list/questions",userHandler.GetQuestions)
			user.POST("/organization/create", userHandler.CreateOrganization)
			user.PATCH("/organization/join",userHandler.JoinOrganization)
		}
	}


		//admin routes
	admin := engine.Group("admin")
	{
		admin.POST("/signup", authHandler.AdminSignup)
		admin.POST("/login", authHandler.AdminLogin)
		
	admin.Use(middleware.AuthorizeJwt())
		{
			admin.GET("/token/refresh", authHandler.AdminRefreshToken)
			admin.PATCH("/approveevent",adminHandler.ApproveEvent)
			admin.POST("/event/create", eventHandler.CreateEventAdmin)
			admin.GET("/listUsers",adminHandler.ViewAllUsers)
			admin.PATCH("/vipuser",adminHandler.VipUser)
			admin.GET("/listEvents", adminHandler.ViewAllEvents)
			admin.PATCH("/organization/register",adminHandler.RegisterOrganization)
			admin.PATCH("/organization/reject",adminHandler.RejectOrganization)
			admin.GET("/list-organizations",adminHandler.ListOrgRequests)
		}
	}


		//eventroutes
		event := engine.Group("event")
		{
			
			event.GET("/approved", eventHandler.ViewAllApprovedEvents)
			event.GET("/geteventbytitle", eventHandler.GetEventByTitle)
			event.PATCH("/update",eventHandler.UpdateEvent)
			event.DELETE("/delete",eventHandler.DeleteEvent)
			
			
		}

		

		engine.Use(middleware.AuthorizeOrg()) 
	{
		engine.GET("/get-organization/",userHandler.GetOrganization)

	}
	

		
			return &ServerHTTP{engine: engine}
	}



func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
