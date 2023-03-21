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