package bootstrap

import (
	"code-base-go/pkg/database"
	"log"

	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	// Initialize the database configuration and connect
	config := database.NewConfig()
	db := config.Connect()

	// Ensure the underlying SQL connection is properly closed
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB connection: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Tes koneksi database gagal: %v", err)
	}

	log.Println("Koneksi database berhasil diinisialisasi")
	return db
}