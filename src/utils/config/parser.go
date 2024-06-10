package config

import (
	"flag"
	"os"

	"github.com/mt190502/portfolio/src/utils/logger"
	"github.com/spf13/viper"
)

var Config ConfigStruct

type ConfigStruct struct {
	Database struct {
		Engine string
		DBHost string
		DBPort string
		DBName string
		DBPath string
		DBUser string
		DBPass string
	}
	Interface struct {
		Language string
	}
	WebServer struct {
		Address string
		Port    int
	}
}

func InitConfig() {
	configFile := "config.yaml"

	version := flag.Bool("version", false, "Print version information")
	serverport := flag.Int("port", 8080, "Port to run the server on")
	configpath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	if *version {
		logger.Log.Infof("Version: 1.0.0")
		os.Exit(0)
	}

	if *configpath != "config.yaml" {
		configFile = *configpath
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		logger.Log.Fatalf("Configuration file: %s does not exist, %v\n", configFile, err)
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Fatalf("Error reading config file, %s\n", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		logger.Log.Fatalf("Unable to decode into struct, %v\n", err)
	}
	logger.Log.Infof("Config file loaded successfully: %s\n", configFile)

	if *serverport != 8080 {
		logger.Log.Infof("Overriding server port to: %d\n", *serverport)
		Config.WebServer.Port = *serverport
	}
}
