package liquibase

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunLiquibase(command string) error {
	// Construct the Liquibase command arguments
	liquibaseCmd := []string{
		"--driver=" + os.Getenv("LIQUIBASE_DRIVER"),
		"--url=" + os.Getenv("LIQUIBASE_URL"),
		"--username=" + os.Getenv("DB_USER"),
		"--password=" + os.Getenv("DB_PASSWORD"),
		"--changeLogFile=" + os.Getenv("LIQUIBASE_CHANGELOG_FILE"),
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