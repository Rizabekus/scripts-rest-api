package storage

import (
	"database/sql"

	"github.com/Rizabekus/scripts-rest-api/internal/models"
)

type Storage struct {
	ScriptStorage models.ScriptStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{ScriptStorage: CreateScriptStorage(db)}
}
