package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClamis struct {
	UserId uint   `json:"userid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (claims JWTClamis) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	return fmt.Errorf("Invalid token")
}

type SignedDetails struct {
	UserId uint `json:"userid"`
	UserName string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}