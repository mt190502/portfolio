package database

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mt190502/portfolio/src/utils/config"
	"github.com/mt190502/portfolio/src/utils/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

type Profile struct {
	gorm.Model
	Sitetitle string
	About     string
	Name      string
	Email     string
	Phone     string
	Address   string
	Linkedin  string
	Github    string
	Instagram string
	Facebook  string
	Twitter   string
}

type Credential struct {
	gorm.Model
	Email    string
	Password string
}

var InitialProfile = Profile{
	Sitetitle: "Portfolio",
	About:     "This is a simple portfolio website.",
	Name:      "John Doe",
	Email:     "john.doe@example.com",
	Phone:     "+1234567890",
	Address:   "123 Main St, City, Country",
	Linkedin:  "johndoe",
	Github:    "johndoe",
	Instagram: "johndoe",
	Facebook:  "johndoe",
	Twitter:   "johndoe",
}

func InitSQLite3() {
	db, err := gorm.Open(sqlite.Open(config.Config.Database.DBPath), &gorm.Config{
		Logger: gorm_logger.Default.LogMode(gorm_logger.Silent),
	})
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}
	DBConnection = db
	DBConnection.AutoMigrate(&Profile{})
	DBConnection.AutoMigrate(&Credential{})
}

func InitDatabase() {
	switch config.Config.Database.Engine {
	case "sqlite3":
		InitSQLite3()
	default:
		logger.Log.Fatalf("Database engine '%v' not supported yet...", config.Config.Database.Engine)
		os.Exit(1)
	}
	siteTitle := DBConnection.First(&Profile{})
	if siteTitle.Error != nil {
		if siteTitle.Error.Error() == "record not found" {
			DBConnection.Create(&InitialProfile)
		}
	}
}
