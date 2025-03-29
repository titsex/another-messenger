package database

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/titsex/another-messenger/services/users/internal/config"
	"github.com/titsex/another-messenger/services/users/internal/logger"
	"github.com/titsex/another-messenger/services/users/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	GormLogger "gorm.io/gorm/logger"
	"time"
)

var Loaded *gorm.DB

type CustomGormLogger struct {
	GormLogger.Interface
}

func (l *CustomGormLogger) LogMode(level GormLogger.LogLevel) GormLogger.Interface {
	return &CustomGormLogger{
		Interface: l.Interface.LogMode(level),
	}
}

func (l *CustomGormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if logger.Print.GetLevel() <= zerolog.InfoLevel {
		logger.Print.Info().Msgf(msg, data...)
	}
}

func (l *CustomGormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if logger.Print.GetLevel() <= zerolog.WarnLevel {
		logger.Print.Warn().Msgf(msg, data...)
	}
}

func (l *CustomGormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	logger.Print.Error().Msgf(msg, data...)
}

func (l *CustomGormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	duration := time.Since(begin)

	event := logger.Print.Debug()
	if err != nil {
		event = logger.Print.Error().Err(err)
	}

	if logger.Print.GetLevel() <= zerolog.DebugLevel {
		event.Str("sql", sql).Int64("rows", rows).Dur("duration", duration).Msg("SQL executed")
	}
}

func Load() {
	databaseConfig := config.Loaded.Database

	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Password, databaseConfig.Name)

	database, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		Logger: &CustomGormLogger{},
	})
	if err != nil {
		logger.Print.Fatal().Msgf("Failed to connect to database, %v", err)
	}

	Loaded = database
	logger.Print.Info().Msg("Database loaded successfully")
}

func AutoMigrate() {
	err := Loaded.AutoMigrate(&models.User{})
	if err != nil {
		logger.Print.Fatal().Msgf("Failed to migrate database: %v", err)
	}

	logger.Print.Info().Msg("Database migrated successfully")
}
