package domain

type JWTClamis struct {
	UserId uint `json:"userid"`
	Email string `json:"email"`
	Role string `json:"role"`
}

