package usecase

import (
	"errors"
	"fmt"
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

	claims.ExpiresAt = time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}
	return refreshToken, err
}

// GenerateToken implements interfaces.JWTUsecase
func (j *jwtUsecase) GenerateToken(userid uint, username string, role string) string {

	claims := &domain.SignedDetails{
		UserId:   userid,
		UserName: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(24)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		log.Println(err)
	}

	return signedToken
}

// GetTokenFromString implements interfaces.JWTUsecase
func (j *jwtUsecase) GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error){
	return jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.SecretKey), nil
	})

}

// VerifyToken implements interfaces.JWTUsecase
func (j *jwtUsecase) VerifyToken(signedToken string) (bool, *domain.SignedDetails) {
	claims := &domain.SignedDetails{}
	token, _ := j.GetTokenFromString(signedToken, claims)

	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
}
}
return false , claims
}
	




func NewJWTUsecase() usecase.JWTUsecase {
	return &jwtUsecase{
		SecretKey: os.Getenv("USER_KEY"),
	}
}

