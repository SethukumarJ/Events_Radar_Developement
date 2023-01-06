package main

import (
	"fmt"
	"log"

	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	"github.com/thnkrn/go-gin-clean-arch/pkg/db"
	di "github.com/thnkrn/go-gin-clean-arch/pkg/di"
)

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
