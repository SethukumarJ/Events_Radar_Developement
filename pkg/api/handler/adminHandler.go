package handler

import (
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface")



type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
	userUsecase usecase.UserUseCase
}


func NewAdminHandler(
	adminUsecase usecase.AdminUsecase,
	userUsecase usecase.UserUseCase,
) AdminHandler {
	return AdminHandler{
		adminUsecase: adminUsecase,
		userUsecase: userUsecase,
	}
}


