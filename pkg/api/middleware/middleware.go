package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thnkrn/go-gin-clean-arch/pkg/response"
	"github.com/thnkrn/go-gin-clean-arch/pkg/utils"
	
	usecase "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type Middleware interface {
	AuthorizeJwt() gin.HandlerFunc
}

type middleware struct {
	jwtUsecase usecase.JWTUsecase
}


func NewMiddlewareUser(jwtUserUsecase usecase.JWTUsecase) Middleware {
	return &middleware{
		jwtUsecase: jwtUserUsecase,
	}

}

// func NewMiddlewareAdmin(jwtAdminUsecase usecase.JWTUsecase) Middleware {
// 	return &middleware{
// 		jwtUsecase: jwtAdminUsecase,
// 	}

// }

func (cr *middleware) AuthorizeJwt() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Writer.Header().Get("Authorization")
		bearerToken := strings.Split(autheader, " ")

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			return
		}

		authtoken := bearerToken[1]
		ok, claims := cr.jwtUsecase.VerifyToken(authtoken)

		if !ok {
			err := errors.New("your token is not valid")
			response := response.ErrorResponse("Error", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			return
		}

		user_email := fmt.Sprint(claims.UserName)
		c.Writer.Header().Set("email", user_email)
		c.Next()

	})
}
