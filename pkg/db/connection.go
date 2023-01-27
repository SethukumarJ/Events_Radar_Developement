package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
)

func ConnectDatabase(cfg config.Config) *sql.DB {

	databaseName := cfg.DBName

	dbURI := cfg.DBSOURCE

	//Opens database
	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	// verifies connection to the database is still alive
	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging")
		log.Fatal(err)

	}

	log.Println("\nConnected to database:", databaseName)

	return db

}
