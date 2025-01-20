package liquibase

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func CreateMigrationFile(migrationName, author, tableName string) error {
	timestamp := time.Now().Format("20060102150405")

	// Set migration directory
	migrationDir := "db/changelog"
	fileName := fmt.Sprintf("%s/%s_%s.xml", migrationDir, timestamp, migrationName)

	// Ensure the directory exists
	if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		return fmt.Errorf("Error creating directory: %v\n", err)
	}

	// Get the latest ID from existing changelogs
	latestID := getLatestChangeSetID(migrationDir)

	// Increment ID for the new changeSet
	newID := latestID + 1

	// Template for XML migration
	content := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<databaseChangeLog
    xmlns="http://www.liquibase.org/xml/ns/dbchangelog"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://www.liquibase.org/xml/ns/dbchangelog
        http://www.liquibase.org/xml/ns/dbchangelog/dbchangelog-4.17.xsd">

    <changeSet id="%d" author="%s">
        <createTable tableName="%s">
            <column name="id" type="BIGINT">
                <constraints primaryKey="true" />
            </column>
            <column name="created_at" type="TIMESTAMP" defaultValueComputed="CURRENT_TIMESTAMP" />
            <column name="updated_at" type="TIMESTAMP" defaultValueComputed="CURRENT_TIMESTAMP" />
        </createTable>
    </changeSet>

</databaseChangeLog>
`, newID, author, tableName)

	// Write the file
	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return fmt.Errorf("Error writing file: %v\n", err)
	}

	fmt.Printf("Migration file created successfully: %s\n", fileName)

	// Update master.xml
	masterFile := filepath.Join(migrationDir, "master.xml")
	if err := updateMasterFile(masterFile, fileName); err != nil {
		fmt.Printf("Error updating master.xml: %v\n", err)
	}

	return nil
}

// getLatestChangeSetID finds the highest changeSet ID in the existing changelogs
func getLatestChangeSetID(dir string) int {
	files, err := filepath.Glob(filepath.Join(dir, "*.xml"))
	if err != nil {
		fmt.Printf("Error reading changelog files: %v\n", err)
		return 0
	}

	var latestID int
	for _, file := range files {
		id := extractHighestChangeSetID(file)
		if id > latestID {
			latestID = id
		}
	}

	return latestID
}

func extractHighestChangeSetID(file string) int {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", file, err)
		return 0
	}

	// Search for <changeSet id="...">
	var highestID int
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, `<changeSet id="`) {
			id := parseChangeSetID(line)
			if id > highestID {
				highestID = id
			}
		}
	}

	return highestID
}

func parseChangeSetID(line string) int {
	parts := strings.Split(line, `"`)
	if len(parts) > 1 {
		if id, err := strconv.Atoi(parts[1]); err == nil {
			return id
		}
	}
	return 0
}

func updateMasterFile(masterFile, newFile string) error {
	// Ensure master.xml exists
	if _, err := os.Stat(masterFile); os.IsNotExist(err) {
		if err := os.WriteFile(masterFile, []byte(`<?xml version="1.0" encoding="UTF-8"?>
<databaseChangeLog
    xmlns="http://www.liquibase.org/xml/ns/dbchangelog"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://www.liquibase.org/xml/ns/dbchangelog
        http://www.liquibase.org/xml/ns/dbchangelog/dbchangelog-4.17.xsd">
</databaseChangeLog>`), 0644); err != nil {
			return fmt.Errorf("error creating master.xml: %v", err)
		}
	}

	// Read current content of master.xml
	content, err := os.ReadFile(masterFile)
	if err != nil {
		return fmt.Errorf("error reading master.xml: %v", err)
	}

	// Ensure the path in the <include> tag includes the directory
	includeTag := fmt.Sprintf(`<include file="db/changelog/%s" />`, filepath.Base(newFile))
	updatedContent := strings.Replace(string(content), "</databaseChangeLog>", "    "+includeTag+"\n</databaseChangeLog>", 1)

	// Write back to master.xml
	if err := os.WriteFile(masterFile, []byte(updatedContent), 0644); err != nil {
		return fmt.Errorf("error writing to master.xml: %v", err)
	}

	fmt.Println("master.xml updated successfully.")
	return nil
}
