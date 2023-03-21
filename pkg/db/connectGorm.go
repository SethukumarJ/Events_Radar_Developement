package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"


	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	domain "github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
)

func ConnectGormDB(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	fmt.Println("Connect gormdb called!")
	psqlInfo := cfg.DBSOURCE
	fmt.Printf("\n\nsql : %v\n\n", psqlInfo)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// Migrating models
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Verification{})
	db.AutoMigrate(&domain.Admins{})
	db.AutoMigrate(&domain.Events{})
	db.AutoMigrate(&domain.Bios{})
	db.AutoMigrate(&domain.Answers{})
	db.AutoMigrate(&domain.Organizations{})
	db.AutoMigrate(&domain.Faqas{})	
	db.AutoMigrate(&domain.Org_Status{})
	db.AutoMigrate(&domain.User_Organization_Connections{})
	db.AutoMigrate(&domain.Join_Status{})
	db.AutoMigrate(&domain.Notificaiton{})
	db.AutoMigrate(&domain.ApplicationForm{})
	db.AutoMigrate(&domain.Appllication_Statuses{})
	db.AutoMigrate(&domain.Posters{})
	db.AutoMigrate(&domain.Packages{})
	db.AutoMigrate(&domain.Promotion{})

	return db, dbErr
}
