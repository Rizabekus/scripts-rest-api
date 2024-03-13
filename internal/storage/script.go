package storage

import (
	"database/sql"
	"fmt"

	"github.com/Rizabekus/scripts-rest-api/internal/models"
)

type ScriptDB struct {
	DB *sql.DB
}

func CreateScriptStorage(db *sql.DB) *ScriptDB {
	return &ScriptDB{DB: db}
}

func (sdb *ScriptDB) AddScript(script string) error {
	sqlStatement := `
	INSERT INTO scripts (script)
	VALUES ($1)
`

	_, err := sdb.DB.Exec(sqlStatement, script)
	if err != nil {

		return fmt.Errorf("error executing SQL statement in AddPerson: %v", err)
	}

	return nil
}
func (sdb *ScriptDB) GetScripts() ([]models.Command, error) {
	rows, err := sdb.DB.Query("SELECT id, script FROM scripts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commands []models.Command

	for rows.Next() {
		var command models.Command
		if err := rows.Scan(&command.Id, &command.Script); err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}
func (sdb *ScriptDB) GetScript(id string) (models.Command, error) {
	var command models.Command

	err := sdb.DB.QueryRow("SELECT id, script FROM scripts WHERE id = $1", id).Scan(&command.Id, &command.Script)
	if err != nil {

		return command, err
	}

	return command, nil
}
