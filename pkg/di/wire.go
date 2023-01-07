//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/thnkrn/go-gin-clean-arch/pkg/api"
	handler "github.com/thnkrn/go-gin-clean-arch/pkg/api/handler"
	middleware "github.com/thnkrn/go-gin-clean-arch/pkg/api/middleware"
	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	db "github.com/thnkrn/go-gin-clean-arch/pkg/db"
	repository "github.com/thnkrn/go-gin-clean-arch/pkg/repository"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, 
			repository.NewUserRepository, 
			repository.NewAdminRespository,
			config.NewMailConfig,
			usecase.NewJWTUserUsecase,
			usecase.NewAuthUsecase,
			usecase.NewAdminUsecase,
			usecase.NewUserUseCase, 
			handler.NewUserHandler,
			handler.NewAuthHandler,
			handler.NewAdminHandler,
			middleware.NewMiddlewareUser,
			http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
