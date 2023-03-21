package repository_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	repository "github.com/SethukumarJ/Events_Radar_Developement/pkg/repository"
)


func TestCreateAdmin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repo := repository.NewAdminRespository(db)

	// Test Case 1: Successful insert
	admin := domain.Admins{
		AdminName:    "testuser",
		Email:       "johndoe@example.com",
		PhoneNumber: "555-555-5555",
		Password:    "password",
	}
	
	mock.ExpectQuery(`INSERT INTO admins\(admin_name,email,phone_number,password\)VALUES\(\$1, \$2, \$3, \$4\)RETURNING admin_id;`).
		WithArgs(admin.AdminName,admin.Email, admin.PhoneNumber, admin.Password).
		WillReturnRows(sqlmock.NewRows([]string{"admin_id"}).AddRow(1))



	adminID, err := repo.CreateAdmin(admin)
	assert.NoError(t, err)
	assert.Equal(t, 1, adminID)

	// Test Case 2: Duplicate username
	mock.ExpectQuery(`INSERT INTO admins\(admin_name,email,phone_number,password\)\VALUES(\$1, \$2, \$3, \$4\)RETURNING admin_id;`).
		WithArgs(admin.AdminName,admin.Email, admin.PhoneNumber, admin.Password).
		WillReturnError(errors.New("duplicate key value violates unique constraint"))

	adminID, err = repo.CreateAdmin(admin)
	assert.Error(t, err)
	assert.Equal(t, 0, adminID)

}


func TestFindAdminById(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %s", err)
	}
	defer db.Close()

	// Create a new userRepository instance with the mock database connection
	adminRepo := repository.NewAdminRespository(db)

	// Define a test case
	testCase := struct {
		adminID       int
		expectedUser domain.AdminResponse
		expectedErr  error
	}{
		adminID: 1,
		expectedUser: domain.AdminResponse{
			AdminId:       1,
			AdminName:     "john.doe",
			Email:        "john.doe@example.com",
			Password:     "passw0rd",
			Verification: false,
			PhoneNumber:  "1234567890",
			
		},
		expectedErr: nil,
	}

	// Define the expected SQL query and result
	// query := `SELECT admin_id,admin_name,email,password,phone_number FROM admins WHERE admin_id = \$1;`
	mock.ExpectQuery(`SELECT admin_id,admin_name,email,password,phone_number FROM admins WHERE admin_id = \$1;`).WithArgs(testCase.adminID).WillReturnRows(
		sqlmock.NewRows([]string{"admin_id", "admin_name", "email", "password", "phone_number"}).
			AddRow(testCase.expectedUser.AdminId, testCase.expectedUser.AdminName,testCase.expectedUser.Email, testCase.expectedUser.Password, testCase.expectedUser.PhoneNumber))

	// Call the FindUserById method with the test case user ID
	admin, err := adminRepo.FindAdminById(testCase.adminID)

	// Check the result and error against the expected values
	assert.Equal(t, testCase.expectedErr, err)
	assert.Equal(t, testCase.expectedUser, admin)
}

func TestFindAdminByName(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %s", err)
	}
	defer db.Close()

	// Create a new userRepository instance with the mock database connection
	adminRepo := repository.NewAdminRespository(db)

	// Define a test case
	testCase := struct {
		adminName       string
		expectedUser domain.AdminResponse
		expectedErr  error
	}{
		adminName: "john.doe",
		expectedUser: domain.AdminResponse{
			AdminId:       1,
			AdminName:     "john.doe",
			Email:        "john.doe@example.com",
			Password:     "passw0rd",
			Verification: false,
			PhoneNumber:  "1234567890",
			
		},
		expectedErr: nil,
	}

	// Define the expected SQL query and result
	// query := `SELECT admin_id,admin_name,email,password,phone_number FROM admins WHERE admin_id = \$1;`
	mock.ExpectQuery(`SELECT admin_id,admin_name,email,password,phone_number FROM admins WHERE email = \$1 OR admin_name = \$2;`).WithArgs(testCase.adminName,testCase.adminName).WillReturnRows(
		sqlmock.NewRows([]string{"admin_id", "admin_name", "email", "password", "phone_number"}).
			AddRow(testCase.expectedUser.AdminId, testCase.expectedUser.AdminName,testCase.expectedUser.Email, testCase.expectedUser.Password, testCase.expectedUser.PhoneNumber))

	// Call the FindUserById method with the test case user ID
	admin, err := adminRepo.FindAdminByName(testCase.adminName)

	// Check the result and error against the expected values
	assert.Equal(t, testCase.expectedErr, err)
	assert.Equal(t, testCase.expectedUser, admin)
}
