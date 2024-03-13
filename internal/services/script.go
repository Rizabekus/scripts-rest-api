package services

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/scripts-rest-api/internal/models"
)

type ScriptService struct {
	storage models.ScriptStorage
}

func CreateScriptService(storage models.ScriptStorage) *ScriptService {
	return &ScriptService{storage: storage}
}

func (ScriptService *ScriptService) AddScript(script string) error {
	return ScriptService.storage.AddScript(script)
}
func (ScriptService *ScriptService) SendResponse(response string, w http.ResponseWriter, statusCode int) {
	responseInstance := models.ErrorResponse{Details: response}
	responseJSON, err := json.Marshal(responseInstance)
	if err != nil {

		internalError := models.ErrorResponse{Details: "Internal Server Error"}
		internalErrorJSON, _ := json.Marshal(internalError)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
func (ScriptService *ScriptService) GetScripts() ([]models.Command, error) {
	return ScriptService.storage.GetScripts()
}
func (ScriptService *ScriptService) GetScript(id string) (models.Command, error) {
	return ScriptService.storage.GetScript(id)
}
