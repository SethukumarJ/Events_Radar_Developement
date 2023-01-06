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

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
