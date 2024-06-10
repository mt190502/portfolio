package main

import (
	"github.com/mt190502/portfolio/src/api"
	"github.com/mt190502/portfolio/src/utils/config"
	"github.com/mt190502/portfolio/src/utils/database"
	"github.com/mt190502/portfolio/src/utils/localizator"
	"github.com/mt190502/portfolio/src/utils/logger"
)

func main() {
	// Init logger
	logger.InitLogger()

	// Parse config
	config.InitConfig()

	// Init database
	database.InitDatabase()

	// Init localizator
	localizator.InitLocalizator()

	// Init API service
	api.InitAPI()
} 
