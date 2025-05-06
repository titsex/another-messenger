package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/titsex/another-messenger/services/users/internal/config"
)

var Connection *pgx.Conn

func Connect(databaseConfig *config.DatabaseConfig) {
	connectionUrl := buildConnectionURL(databaseConfig)\

	connection, err := pgx.Connect(context.Background(), connectionUrl)
	if err != nil {
		panic(err)
	}
	
	Connection = connection

}

func buildConnectionURL(databaseConfig *config.DatabaseConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database)
}