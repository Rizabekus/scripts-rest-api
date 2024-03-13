package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rizabekus/scripts-rest-api/internal/models"
	"github.com/Rizabekus/scripts-rest-api/pkg/loggers"
	"github.com/Rizabekus/scripts-rest-api/pkg/validators"
	"github.com/go-playground/validator"
)

func (handler *Handlers) AddScript(w http.ResponseWriter, r *http.Request) {
	var newScript models.Command
	err := json.NewDecoder(r.Body).Decode(&newScript)
	if err != nil {

		handler.Service.ScriptService.SendResponse("Failed To decode JSON", w, http.StatusBadRequest)
		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}
	loggers.DebugLog.Println("Received data in JSON format")
	validate := validator.New()
	validate.RegisterValidation("bash", validators.IsValidBashScript)
	err = validate.Struct(newScript)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {

			handler.Service.ScriptService.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
			loggers.InfoLog.Println("Internal Server Error")
			return
		}
		firstValidationError := validationErrors[0]

		handler.Service.ScriptService.SendResponse(fmt.Sprintf("Field: %s, Tag: %s\n", firstValidationError.Field(), firstValidationError.Tag()), w, http.StatusBadRequest)
		loggers.InfoLog.Println("Validation Error", err)
		return
	}
	loggers.DebugLog.Println("Validated the data", newScript)
	err = handler.Service.ScriptService.AddScript(newScript.Script)
	if err != nil {
		handler.Service.ScriptService.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		loggers.InfoLog.Println("Internal Server Error")
		return
	}
	loggers.DebugLog.Println("Script successfully was added to database")
	w.WriteHeader(http.StatusCreated)
}
