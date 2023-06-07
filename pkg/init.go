package pkg

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:  "init",
	Long: `create database, create tables `,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initClickhouse(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	root.AddCommand(initCommand)
}

func initClickhouse() error {
	// Read ClickHouse connection information from environment variables
	clickHouseURL := os.Getenv("CLICKHOUSE_URL")
	if clickHouseURL == "" {
		return fmt.Errorf("CLICKHOUSE_URL environment variable is not set")
	}

	// Connect to ClickHouse database
	db, err := sql.Open("clickhouse", clickHouseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to ClickHouse: %v", err)
	}
	defer db.Close()

	// Create the database
	if err := executeSQLFile(db, "scripts/test_database.sql"); err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	// Create the tables
	if err := executeSQLFile(db, "scripts/test_tables.sql"); err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	fmt.Println("Database and tables created successfully")

	return nil
}

func executeSQLFile(db *sql.DB, filePath string) error {
	// Read the content of the SQL file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)
	}

	// Split the file content into individual SQL statements
	statements := strings.Split(string(fileContent), ";")

	// Execute each SQL statement
	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)
		if trimmedStatement == "" {
			continue // Skip empty statements
		}

		_, err := db.Exec(trimmedStatement)
		if err != nil {
			return fmt.Errorf("failed to execute SQL statement: %v", err)
		}
	}

	return nil
}