package main

import (
	"code-base-go/pkg/config"
	"code-base-go/pkg/liquibase"
	"log"

	"fmt"
	"os"
)

func main() {
	// Load file .env
	err := config.LoadEnvFile("../../.env")
	if err != nil {
		log.Println("Warning: file .env not found, using environment default.")
	}


	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a Liquibase command. Example: go run main.go update")
		return
	}

	command := os.Args[2]
	if err := liquibase.RunLiquibase(command); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Liquibase command executed successfully!")
}