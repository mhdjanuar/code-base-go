package main

import (
	"code-base-go/pkg/config"
	"code-base-go/pkg/liquibase"
	"fmt"
	"log"
	"os"
)

func main() {
	// Load file .env
	err := config.LoadEnvFile("../../.env")
	if err != nil {
		log.Println("Warning: file .env not found, using environment default.")
	}

	// Ensure at least one argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a Liquibase command or action.")
		fmt.Println("Examples:")
		fmt.Println("  Run Liquibase command: go run main.go update")
		fmt.Println("  Create migration file: go run main.go create migration_name author table_name")
		os.Exit(1)
	}

	// Get the first argument as the action
	action := os.Args[1]
	switch action {
		case "update", "rollback", "status":
			command := os.Args[1]
			if err := liquibase.RunLiquibase(command); err != nil {
				fmt.Println("Error running Liquibase:", err)
				os.Exit(1)
			}
	
			fmt.Println("Liquibase command executed successfully!")
		case "create":
			// Create a new migration file
			if len(os.Args) < 5 {
				fmt.Println("Error: Please provide migration name, author, and table name.")
				fmt.Println("Example: go run main.go create migration_name yourname table_name")
				os.Exit(1)
			}

			migrationName := os.Args[2]
			author := os.Args[3]
			tableName := os.Args[4]

			if err := liquibase.CreateMigrationFile(migrationName, author, tableName); err != nil {
				fmt.Println("Error creating migration file:", err)
				os.Exit(1)
			}

			fmt.Println("Migration file created successfully!")
		default:
			fmt.Println("Error: Invalid action. Supported actions are 'update', 'rollback', 'status', and 'create'.")
			os.Exit(1)
	}
}