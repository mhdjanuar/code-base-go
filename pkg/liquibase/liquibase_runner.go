package liquibase

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunLiquibase(command string) error {
	// Get the database configurations from the environment
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	liquibaseURL := fmt.Sprintf("jdbc:postgresql://%s:%s/%s", dbHost, dbPort, dbName)

	// Construct the Liquibase command arguments
	liquibaseCmd := []string{
		"--driver=" + "org.postgresql.Driver",
		"--url=" + liquibaseURL,
		"--username=" + os.Getenv("DB_USER"),
		"--password=" + os.Getenv("DB_PASSWORD"),
		"--changeLogFile=" + "db/changelog/master.xml",
		command,
	}

	// Create the command
	cmd := exec.Command("liquibase", liquibaseCmd...)
	cmd.Stdout = os.Stdout // Redirect standard output
	cmd.Stderr = os.Stderr // Redirect standard error

	// Print the command for debugging
	fmt.Println("Running Liquibase command:", strings.Join(cmd.Args, " "))

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Liquibase command failed: %w", err)
	}

	return nil
}