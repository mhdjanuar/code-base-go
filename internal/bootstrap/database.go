package bootstrap

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitializeDatabase() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "P@ssw0rd"
	dbname := "godb"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	log.Println("Koneksi ke database berhasil")
	return db
}