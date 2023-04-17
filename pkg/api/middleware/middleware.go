package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/utils"
	"github.com/gin-gonic/gin"

	usecase "github.com/SethukumarJ/Events_Radar_Developement/pkg/usecase/interface"
)

type Middleware interface {
	AuthorizeJwt() gin.HandlerFunc
	AuthorizeJwtAdmin() gin.HandlerFunc
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
		organization_id, _ := strconv.Atoi(c.Query("Organization_id"))
		fmt.Println("organization id from middleware", organization_id)

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
		ok, claims := cr.jwtUsecase.VerifyTokenUser(authtoken)
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
		userId := claims.UserId
		fmt.Println("username", userName,"user_id",userId,"organization_id",organization_id)
		role, err := cr.userUsecase.VerifyRole(int(userId), int(organization_id))
		fmt.Println(role, "///////////", err, "role and pathrole")
		if err != nil {
			err = errors.New("your role input is invalid")
			response := response.ErrorResponse("Error", err.Error(), role)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		c.Writer.Header().Set("userName", userName)
		c.Writer.Header().Set("user_id", fmt.Sprint(userId))
		c.Writer.Header().Set("organization_id", fmt.Sprint(organization_id))
		c.Writer.Header().Set("role", role)
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
		ok, claims := cr.jwtUsecase.VerifyTokenUser(authtoken)
		source := fmt.Sprint(claims.Source)
		fmt.Println("///////////////token role", claims.Role)
		if claims.Role != "user" {
			err := errors.New("your role of the token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

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
		} else if ok && source == "refreshtoken" {

			c.Writer.Header().Set("Authorization", authtoken)
			c.Next()
		} else {

			userName := fmt.Sprint(claims.UserName)
			userId := fmt.Sprint(claims.UserId)
			c.Writer.Header().Set("userName", userName)
			c.Writer.Header().Set("user_id",userId)
			c.Next()

		}

	})
}

func (cr *middleware) AuthorizeJwtAdmin() gin.HandlerFunc {
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
			ok, claims := cr.jwtUsecase.VerifyTokenAdmin(authtoken)
			fmt.Println("ok",ok)
			source := fmt.Sprint(claims.Source)
			fmt.Println("///////////////token role", claims.Role)
			if claims.Role != "admin" {
				err := errors.New("your role of the token is not valid")
				response := response.ErrorResponse("Error", err.Error(), source)
				c.Writer.Header().Add("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusUnauthorized)
				utils.ResponseJSON(*c, response)
				c.Abort()
				return
			}
	
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
			} else if ok && source == "refreshtoken" {
	
				c.Writer.Header().Set("Authorization", authtoken)
				c.Next()
			} else {
	
				userName := fmt.Sprint(claims.UserName)
				userId := fmt.Sprint(claims.UserId)
				c.Writer.Header().Set("userName", userName)
				c.Writer.Header().Set("user_id",userId)
				c.Next()
	
			}
	
		})
}
