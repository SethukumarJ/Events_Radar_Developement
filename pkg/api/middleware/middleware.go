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
	AuthorizeOrg() gin.HandlerFunc
}

type middleware struct {
	jwtUsecase  usecase.JWTUsecase
	userUsecase usecase.UserUseCase
}



func NewMiddlewareUser(jwtUserUsecase usecase.JWTUsecase, userUsecase usecase.UserUseCase) Middleware {
	return &middleware{
		jwtUsecase:  jwtUserUsecase,
		userUsecase: userUsecase,
	}

}


// AuthorizeOrg implements Middleware
func (cr *middleware) AuthorizeOrg() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Request.Header["Authorization"]
		organizationName := c.Query("organizationName")
		pathRole := c.Param("pathRole")
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntoken : %v\n\n", autheader)

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()

			return
		}

		authtoken := bearerToken[1]
		ok, claims := cr.jwtUsecase.VerifyToken(authtoken)
		source := fmt.Sprint(claims.Source)

		if !ok && source == "accesstoken" {
			err := errors.New("your access token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if !ok && source == "refreshtoken" {
			err := errors.New("your refresh token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		userName := fmt.Sprint(claims.UserName)

		role , err := cr.userUsecase.VerifyRole(userName, organizationName)
		fmt.Println(role, pathRole, "role and pathrole")
		if role != pathRole {
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}


		c.Writer.Header().Set("userName", userName)
		c.Writer.Header().Set("organizationName",organizationName)
		c.Writer.Header().Set("role",role)
		c.Next()

	})
}

func (cr *middleware) AuthorizeJwt() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Request.Header["Authorization"]
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntocen : %v\n\n", autheader)

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()

			return
		}

		authtoken := bearerToken[1]
		ok, claims := cr.jwtUsecase.VerifyToken(authtoken)
		source := fmt.Sprint(claims.Source)

		if !ok && source == "accesstoken" {
			err := errors.New("your access token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if !ok && source == "refreshtoken" {
			err := errors.New("your refresh token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		userName := fmt.Sprint(claims.UserName)
		c.Writer.Header().Set("userName", userName)
		c.Next()

	})
}
