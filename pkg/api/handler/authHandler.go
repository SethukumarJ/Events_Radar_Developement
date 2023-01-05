package handler

import (
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type AuthHandler struct {
	jwtUserUsecase usecase.JWTUsecase
	authUsecase usecase.AuthUsecase
	userUsecase usecase.UserUseCase
}

func NewAuthHandler(
	jwtUserUsecase usecase.JWTUsecase,
	userUsecase usecase.UserUseCase,
	authUsecase usecase.AuthUsecase,

) *AuthHandler{
	return &AuthHandler{
		jwtUserUsecase: jwtUserUsecase,
		authUsecase: authUsecase,
		userUsecase: userUsecase,
	}
}
