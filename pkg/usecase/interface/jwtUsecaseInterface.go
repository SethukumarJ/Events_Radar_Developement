package interfaces

import (
	"github.com/golang-jwt/jwt"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
)

type JWTUsecase interface {
	GenerateAccessToken(userid uint, userName string, role string) (string, error)
	VerifyToken(token string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error)
	GenerateRefreshToken(userid uint, userName string, role string) (string, error)

}
