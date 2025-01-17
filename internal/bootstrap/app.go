package bootstrap

import (
	"code-base-go/pkg/database"
	"database/sql"
)

func InitializeDatabase() *sql.DB {
	config := database.NewConfig()
	return config.Connect()
}