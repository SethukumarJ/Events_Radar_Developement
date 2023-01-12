package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"


	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	domain "github.com/thnkrn/go-gin-clean-arch/pkg/domain"
)

func ConnectGormDB(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	fmt.Println("Connect gormdb called!")
	psqlInfo := cfg.DBSOURCE
	fmt.Printf("\n\nsql : %v\n\n", psqlInfo)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Verification{})
	db.AutoMigrate(&domain.Admins{})
	db.AutoMigrate(&domain.Events{})
	db.AutoMigrate(&domain.Bios{})
	db.AutoMigrate(&domain.Faqas{})
	db.AutoMigrate(&domain.Answers{})

	return db, dbErr
}
