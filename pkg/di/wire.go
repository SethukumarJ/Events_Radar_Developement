//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/SethukumarJ/Events_Radar_Developement/pkg/api"
	handler "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/handler"
	middleware "github.com/SethukumarJ/Events_Radar_Developement/pkg/api/middleware"
	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	db "github.com/SethukumarJ/Events_Radar_Developement/pkg/db"
	repository "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository"
	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, 
			repository.NewUserRepository, 
			repository.NewAdminRespository,
			repository.NewEventRepository,
			config.NewMailConfig,
			usecase.NewJWTUsecase,
			usecase.NewAuthUsecase,
			usecase.NewAdminUsecase,
			usecase.NewEventUseCase,
			usecase.NewUserUseCase, 
			handler.NewUserHandler,
			handler.NewAuthHandler,
			handler.NewAdminHandler,
			handler.NewEventHandler,
			middleware.NewMiddlewareUser,		
			http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
