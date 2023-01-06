// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/thnkrn/go-gin-clean-arch/pkg/api"
	"github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
	"github.com/thnkrn/go-gin-clean-arch/pkg/api/middleware"
	"github.com/thnkrn/go-gin-clean-arch/pkg/config"
	"github.com/thnkrn/go-gin-clean-arch/pkg/db"
	"github.com/thnkrn/go-gin-clean-arch/pkg/repository"
	"github.com/thnkrn/go-gin-clean-arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	sqlDB := db.ConnectDatabase(cfg)
	userRepository := repository.NewUserRepository(sqlDB)
	mailConfig := config.NewMailConfig()
	userUseCase := usecase.NewUserUseCase(userRepository, mailConfig, cfg)
	userHandler := handler.NewUserHandler(userUseCase)
	jwtUsecase := usecase.NewJWTUsecase()
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHandler := handler.NewAuthHandler(jwtUsecase, userUseCase, authUsecase)
	middlewareMiddleware := middleware.NewMiddlewareUser(jwtUsecase)
	serverHTTP := http.NewServerHTTP(userHandler, authHandler, middlewareMiddleware)
	return serverHTTP, nil
}
