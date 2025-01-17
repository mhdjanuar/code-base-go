package main

import (
	"net/http"

	adapterHttp "code-base-go/internal/adapter/http" // Alias untuk package internal HTTP
	"code-base-go/internal/bootstrap"
	"code-base-go/internal/repository"
	"code-base-go/internal/usecase"
	"code-base-go/pkg/config"
	"log"
)

func main() {
	// Load file .env
	err := config.LoadEnvFile("../.env")
	if err != nil {
		log.Println("Peringatan: file .env tidak ditemukan, menggunakan environment default.")
	}

	// Inisialisasi database
	db := bootstrap.InitializeDatabase()

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
