package main

import (
	"fmt"
	"log"

	_"github.com/thnkrn/go-gin-clean-arch/pkg/domain"
	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	"github.com/thnkrn/go-gin-clean-arch/pkg/db"
	di "github.com/thnkrn/go-gin-clean-arch/pkg/di"
)

// @title Go + Gin Radar API
// @version 1.0
// @description This is an Events Radar project. You can visit the GitHub repository at https://github.com/SethukumarJ/Events_Radar_Developement

// @contact.name API Support
// @contact.url sethukumarj.com
// @contact.email sethukumarj.76@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey ApiKeyAuth
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

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
