package usecase

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type jwtUsecase struct {
	SecretKey string
}

// GenerateRefreshToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateRefreshToken(accessToken string) (string, error) {
	claims := &domain.SignedDetails{}
	j.GetTokenFromString(accessToken, claims)

	if time.Until(time.Unix(claims.ExpiresAt, 0)) > 30*time.Second {
		return "", errors.New("too early to generate refresh token")
	}

	claims.ExpiresAt = time.Now().Local().Add(time.Minute * time.Duration(5)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}
	return refreshToken, err
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
