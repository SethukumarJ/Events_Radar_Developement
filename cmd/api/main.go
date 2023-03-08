package main

import (
	"fmt"
	"log"

	_"github.com/SethukumarJ/Events_Radar_Developement/pkg/domain"
	config "github.com/SethukumarJ/Events_Radar_Developement/pkg/config"
	"github.com/SethukumarJ/Events_Radar_Developement/pkg/db"
	di "github.com/SethukumarJ/Events_Radar_Developement/pkg/di"
)

// @title Events-Radar API
// @version 1.0
// @description This is an Event Management project. You can visit the GitHub repository at https://github.com/SethukumarJ/Events_Radar_Developement

// @contact.name API Support
// @contact.url sethukumarj.com
// @contact.email sethukumarj.76@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	// db.ConnectDB(config)
	gorm, _ := db.ConnectGormDB(config)
	fmt.Printf("\ngorm : %v\n\n", gorm)
	trigger,_ := db.Triggers(config)
	fmt.Printf("\tTrigger : %v\n\n", trigger)

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
