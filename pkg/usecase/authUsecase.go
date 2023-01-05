package usecase

import (
	repository "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

// authUsecase is the struct for the authentication service
type authUsecase struct {
	userRepo repository.UserRepository
}

// VerifyUser implements interfaces.AuthUsecase
func (*authUsecase) VerifyUser(email string, password string) error {
	panic("unimplemented")
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
) usecase.AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}
