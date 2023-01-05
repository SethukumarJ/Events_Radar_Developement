package interfaces

import domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"

type JWTUsecase interface {
	GenerateToken(userid int, email string, role string) string
	VerifyToken(token string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails)
	GenerateRefreshToken(token string) (string, error)
}
