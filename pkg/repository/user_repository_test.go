package repository

// import (
// 	"database/sql"
// 	"testing"

// 	"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // mockDB implements the sql.DB interface and serves as a mock database
// type mockDB struct {
// 	mock.Mock
// }

// func (m *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
// 	args = append([]interface{}{query}, args...)
// 	ret := m.Called(args...)

// 	return ret.Get(0).(*sql.Row)
// }

// func TestInsertUser(t *testing.T) {
// 	user := domain.Users{
// 		UserName:    "testuser",
// 		FirstName:   "Test",
// 		LastName:    "User",
// 		Email:       "testuser@example.com",
// 		PhoneNumber: "1234567890",
// 		Password:    "password",
// 		Profile:     "default",
// 	}

// 	// Create a new mockDB
// 	db := new(mockDB)

// 	// Set expectations on the mockDB
// 	db.On("QueryRow", `INSERT INTO users(user_name,first_name,last_name,email,phone_number,password,profile)VALUES($1, $2, $3, $4, $5, $6,$7)RETURNING user_id;`,
// 		user.UserName,
// 		user.FirstName,
// 		user.LastName,
// 		user.Email,
// 		user.PhoneNumber,
// 		user.Password,
// 		user.Profile).Return(&sql.Row{})

// 	db.On("QueryRow", `INSERT INTO bios(user_name)VALUES($1);`, user.UserName).Return(&sql.Row{})

// 	// Create a new userRepository with the mockDB
// 	r := &userRepository{db: db}

// 	// Call the InsertUser function
// 	id, err := r.InsertUser(user)

// 	// Assert that the returned id and error are as expected
// 	assert.Equal(t, 1, id)
// 	assert.Nil(t, err)

// 	// Assert that all expectations were met
// 	db.AssertExpectations(t)
// }
