// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/api"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/api/handler"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/api/middleware"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/db"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/repository"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	sqlDB := db.ConnectDatabase(cfg)
	userRepository := repository.NewUserRepository(sqlDB)
	adminRepository := repository.NewAdminRespository(sqlDB)
	mailConfig := config.NewMailConfig()
	userUseCase := usecase.NewUserUseCase(userRepository, adminRepository, mailConfig, cfg)
	userHandler := handler.NewUserHandler(userUseCase)
	jwtUsecase := usecase.NewJWTUsecase()
	adminUsecase := usecase.NewAdminUsecase(adminRepository, mailConfig, cfg)
	authUsecase := usecase.NewAuthUsecase(userRepository, adminRepository)
	authHandler := handler.NewAuthHandler(jwtUsecase, userUseCase, adminUsecase, authUsecase, cfg)
	eventRepository := repository.NewEventRepository(sqlDB)
	eventUsecase := usecase.NewEventUseCase(eventRepository)
	adminHandler := handler.NewAdminHandler(adminUsecase, userUseCase, eventUsecase)
	eventHandler := handler.NewEventHandler(adminUsecase, userUseCase, eventUsecase)
	middlewareMiddleware := middleware.NewMiddlewareUser(jwtUsecase, userUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, authHandler, adminHandler, eventHandler, middlewareMiddleware)
	return serverHTTP, nil
}
