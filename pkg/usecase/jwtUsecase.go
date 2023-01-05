package usecase

import (
	"os"

	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type jwtUsecase struct {
	SecretKey string
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (*jwtUsecase) GenerateRefreshToken(token string) (string, error) {
	panic("unimplemented")
}

// GenerateToken implements interfaces.JWTUsecase
func (*jwtUsecase) GenerateToken(userid int, email string, role string) string {
	panic("unimplemented")
}

// GetTokenFromString implements interfaces.JWTUsecase
func (*jwtUsecase) GetTokenFromString(signedToken string, claims *domain.SignedDetails) {
	panic("unimplemented")
}

// VerifyToken implements interfaces.JWTUsecase
func (*jwtUsecase) VerifyToken(token string) (bool, *domain.SignedDetails) {
	panic("unimplemented")
}

func NewJwtUsecase() usecase.JWTUsecase {
	return &jwtUsecase{
		SecretKey: os.Getenv("USER_KEY"),
	}
}
