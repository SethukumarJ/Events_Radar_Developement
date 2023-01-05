package usecase

import (
	"context"

	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/thnkrn/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// CreateUser implements interfaces.UserUseCase
func (*userUseCase) CreateUser(user domain.Users) error {
	panic("unimplemented")
}

// FindUser implements interfaces.UserUseCase
func (*userUseCase) FindUser(email string) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// SendVerificationEmail implements interfaces.UserUseCase
func (*userUseCase) SendVerificationEmail(email string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.UserUseCase
func (*userUseCase) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

// Delete implements interfaces.UserUseCase
func (*userUseCase) Delete(ctx context.Context, user domain.Users) error {
	panic("unimplemented")
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

// func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
// 	err := c.userRepo.Delete(ctx, user)

// 	return err
// }
