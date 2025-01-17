package main

import (
	"net/http"

	"code-base-go/internal/bootstrap"
	adapterHttp "code-base-go/internal/delivery/http" // Alias untuk package internal HTTP
	"code-base-go/internal/repository"
	"code-base-go/internal/usecase"
	"code-base-go/pkg/config"
	"log"
)

func main() {
	// Load file .env
	err := config.LoadEnvFile("../../.env")
	if err != nil {
		log.Println("Warning: file .env not found, using environment default.")
	}

	// Inisialisasi database
	db := bootstrap.InitializeDatabase()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan koneksi database: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Gagal menutup koneksi database: %v", err)
		}
	}()

    // Inisialisasi repository, usecase, dan handler
	userRepo := repository.NewSQLUserRepo(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := adapterHttp.NewUserHandler(userUseCase)

    // Atur route HTTP
	http.HandleFunc("/register", userHandler.RegisterUserHandler)

    // Jalankan server
	log.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
