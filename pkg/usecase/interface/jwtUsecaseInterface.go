package interfaces

import (
	"github.com/golang-jwt/jwt"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

type JWTUsecase interface {
	GenerateToken(userid uint, email string, role string) string
	VerifyToken(token string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error)
	GenerateRefreshToken(token string) (string, error)
}
