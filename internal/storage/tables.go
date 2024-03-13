package storage

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateTables(db *sql.DB) error {
	migrationPath := os.Getenv("MIGRATION_PATH")
	file, err := os.ReadFile(migrationPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(file))
	if err != nil {
		fmt.Println("WHATWHY")
		return err
	}

	return nil
}
