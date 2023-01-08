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

	//userroutes
	user := engine.Group("user")
	{
		user.POST("/signup", authHandler.UserSignup)
		user.POST("/login", authHandler.UserLogin)
		user.POST("/send/verification", userHandler.SendVerificationMail)
		user.PATCH("/verify/account", userHandler.VerifyAccount)

		user.Use(middleware.AuthorizeJwt())
		{
			user.GET("/token/refresh", authHandler.UserRefreshToken)
		}
	}


		//admin routes
	admin := engine.Group("admin")
	{
		admin.POST("/signup", authHandler.AdminSignup)
		admin.POST("/login", authHandler.AdminLogin)
		admin.GET("/listUsers",adminHandler.ViewAllUsers)
		admin.PATCH("/vipuser",adminHandler.VipUser)
		admin.GET("/listEvents", adminHandler.ViewAllEvents)
		admin.PATCH("/approveevent",adminHandler.ApproveEvent)
		
		admin.Use(middleware.AuthorizeJwt())
		{
			admin.GET("/token/refresh", authHandler.AdminRefreshToken)
		}
	}


		//userroutes
		event := engine.Group("event")
		{
			event.POST("/createevent", eventHandler.CreateEvent)
			event.GET("/getApprovedEvents", eventHandler.ViewAllApprovedEvents)
			event.GET("/getEventByTitle", eventHandler.GetEventByTitle)
			event.PATCH("/updateEvent",eventHandler.UpdateEvent)
			event.DELETE("/deleteEvent",eventHandler.DeleteEvent)
			
			
		}
		
			return &ServerHTTP{engine: engine}
	}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
