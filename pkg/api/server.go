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
		user.PATCH("/verify/account",authHandler.VerifyAccount)
		user.PATCH("/password/update",userHandler.UpdatePassword)

		user.Use(middleware.AuthorizeJwt())
		{
			user.POST("/token/refresh", authHandler.UserRefreshToken)
			user.POST("/event/create", eventHandler.CreateEventUser)
			user.POST("/event/post/question", userHandler.PostQuestion)
			user.PATCH("/update/profile",userHandler.UpdateProfile)
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
		
			return &ServerHTTP{engine: engine}
	}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
