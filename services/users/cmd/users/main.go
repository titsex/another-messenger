package main

import (
	"github.com/rs/zerolog"
	"github.com/titsex/another-messenger/services/users/internal/config"
	"github.com/titsex/another-messenger/services/users/internal/database"
	"github.com/titsex/another-messenger/services/users/internal/logger"
)

func main() {
	logger.Init(zerolog.InfoLevel)
	config.Load()

	appConfig := config.Loaded.App
	logger.Init(appConfig.LogLevel)

	database.Load()
	database.AutoMigrate()
}
