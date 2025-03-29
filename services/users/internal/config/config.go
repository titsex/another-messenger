package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/titsex/another-messenger/services/users/internal/logger"
	GormLogger "gorm.io/gorm/logger"
	"os"
)

type DatabaseConfig struct {
	Host     string              `mapstructure:"host"`
	Port     int                 `mapstructure:"port"`
	User     string              `mapstructure:"user"`
	Password string              `mapstructure:"password"`
	Name     string              `mapstructure:"name"`
	LogLevel GormLogger.LogLevel `mapstructure:"log_level"`
}

type AppConfig struct {
	AppName  string        `mapstructure:"app_name"`
	Port     int           `mapstructure:"port"`
	LogLevel zerolog.Level `mapstructure:"log_level"`
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

var Loaded Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		logger.Print.Warn().Msg(".env file not found")
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config"
	}

	appConfig := viper.New()
	appConfig.SetConfigName("app")
	appConfig.SetConfigType("yaml")
	appConfig.AddConfigPath(configPath)

	err = appConfig.ReadInConfig()
	if err != nil {
		logger.Print.Fatal().Err(err).Msg("Error reading app.yaml config file")
	}

	err = appConfig.Unmarshal(&Loaded.App)
	if err != nil {
		logger.Print.Fatal().Err(err).Msg("Failed to parse app.yaml config file")
	}

	databaseConfig := viper.New()
	databaseConfig.SetConfigName("database")
	databaseConfig.SetConfigType("yaml")
	databaseConfig.AddConfigPath(configPath)

	err = databaseConfig.ReadInConfig()
	if err != nil {
		logger.Print.Fatal().Err(err).Msg("Error reading database.yaml config file")
	}

	err = databaseConfig.Unmarshal(&Loaded.Database)
	if err != nil {
		logger.Print.Fatal().Err(err).Msg("Failed to parse database.yaml config file")
	}

	logger.Print.Info().Msg("Configuration loaded successfully")
}
