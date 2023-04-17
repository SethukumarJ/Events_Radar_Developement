package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/SethukumarJ/Events_Radar_Developement/pkg/response"
// 	"github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
// 	mock "github.com/SethukumarJ/Events_Radar_Developement/pkg/mock/usecaseMock"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )


// func TestVerifyAccount(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	c := mock.NewMockUserUseCase(ctrl)
// 	authHandler := NewAuthHandler(nil, c, nil, nil,  config.Config{})

// 	testData := []struct {
// 		name           string
// 		email          string
// 		code           string
// 		beforeTest     func(userUsecase mock.MockUser)
// 		expectCode     int
// 		expectResponse response.Response
// 		expectErr      error
// 	}{
// 		{
// 			name:  "test sucsess response",
// 			email: "jon",
// 			code:  "12345",
// 			beforeTest: func(userUsecase mock.MockAuthUseCase) {
// 				userUsecase.EXPECT().WorkerVerifyAccount("jon", 12345).Return(nil)
// 			},
// 			expectCode: 200,
// 			expectResponse: response.Response{
// 				Status:  true,
// 				Message: "SUCCESS",
// 				Errors:  nil,
// 				Data:    "jon",
// 			},
// 			expectErr: nil,
// 		},
// 		{
// 			name:  "test sucsess response",
// 			email: "ali",
// 			code:  "54321",
// 			beforeTest: func(userUsecase mock.MockAuthUseCase) {
// 				userUsecase.EXPECT().WorkerVerifyAccount("ali", 54321).Return(errors.New("usecase error"))
// 			},
// 			expectCode: 422,
// 			expectResponse: response.Response{
// 				Status:  false,
// 				Message: "Error while verifing worker mail",
// 				Errors:  []interface{}{"usecase error"},
// 				Data:    nil,
// 			},
// 			expectErr: errors.New("usecase error"),
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.beforeTest(*c)

// 			gin := gin.New()
// 			rec := httptest.NewRecorder()

// 			gin.GET("/verify/account", authHandler.WorkerVerifyAccount)

// 			var body []byte
// 			req := httptest.NewRequest("GET", "/verify/account", bytes.NewBuffer(body))
// 			req.Header.Set("Content-Type", "application/json")

// 			// Set a query parameter named "gury" with a value of "param"
// 			q := req.URL.Query()
// 			q.Add("email", tt.email)
// 			q.Add("code", tt.code)
// 			req.URL.RawQuery = q.Encode()

// 			gin.ServeHTTP(rec, req)

// 			var actual response.Response
// 			err := json.Unmarshal(rec.Body.Bytes(), &actual)
// 			assert.NoError(t, err)

// 			assert.Equal(t, tt.expectCode, rec.Code)
// 			assert.Equal(t, tt.expectResponse.Status,actual.Status)
// 			assert.Equal(t,tt.expectResponse.Message,actual.Message)
// 			assert.Equal(t,tt.expectResponse.Errors,actual.Errors)
// 			assert.Equal(t,tt.expectResponse.Data,actual.Data)

// 		})
// 	}
// }



