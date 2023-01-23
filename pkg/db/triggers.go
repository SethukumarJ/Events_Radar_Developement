package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
)

func Triggers(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	fmt.Println("Connect gormdb called!")
	psqlInfo := cfg.DBSOURCE
	fmt.Printf("\n\nsql : %v\n\n", psqlInfo)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	//Migrating triggers

	
	err := db.Exec(joined_notification)
		if err != nil {
		fmt.Println("joined_notificationerr", err)
		}


	err = db.Exec(admit_member_notification_trigger)
		if err != nil {
		fmt.Println("admit_member_notification_trigger", err)
		}


	err = db.Exec(organization_created_notification)
	if err != nil {
	fmt.Println("organization_created_notification_trigger", err)
	}


err = db.Exec(organization_created_notification_trigger)
	if err != nil {
	fmt.Println("organization_created_notification_trigger", err)
	}

	return db, dbErr
}


const (

	// For giving notification on joined status  to the user
	joined_notification = `CREATE OR REPLACE FUNCTION joined_notification() RETURNS TRIGGER AS $$
	BEGIN
	INSERT INTO notificaitons (user_name, organization_name, time, message)
	VALUES (NEW.user_name, NEW.organization_name, NEW.joined_at, 'You are successfully joined to the organization');
	RETURN NEW;
	END; $$ LANGUAGE plpgsql;`
	admit_member_notification_trigger = `CREATE OR REPLACE TRIGGER admit_member_notification
	AFTER INSERT ON user_organization_connections
	FOR EACH ROW
	EXECUTE FUNCTION joined_notification();`

	// For giving notification on joined status  to the user
	organization_created_notification = `CREATE OR REPLACE FUNCTION org_created_notification() RETURNS TRIGGER AS $$
	BEGIN
	INSERT INTO notificaitons (user_name, organization_name, time, message)
	VALUES ((SELECT created_by from organizations where organization_name = NEW.registered), NEW.registered, 
	(SELECT created_at from organizations where organization_name = NEW.registered), 'Organization have been successfully registered');
	RETURN NEW;
	END; $$ LANGUAGE plpgsql;`
	organization_created_notification_trigger = `CREATE OR REPLACE TRIGGER organization_created_notification
	AFTER UPDATE ON org_statuses
	FOR EACH ROW
	EXECUTE FUNCTION org_created_notification();`



)