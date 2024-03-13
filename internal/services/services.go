package services

import (
	"github.com/Rizabekus/scripts-rest-api/internal/models"
	"github.com/Rizabekus/scripts-rest-api/internal/storage"
)

type Services struct {
	ScriptService models.ScriptService
}

func ServiceInstance(storage *storage.Storage) *Services {
	return &Services{
		ScriptService: CreateScriptService(storage.ScriptStorage),
	}
}
