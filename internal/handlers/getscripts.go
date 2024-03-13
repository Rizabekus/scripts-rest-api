package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/scripts-rest-api/pkg/loggers"
)

func (handler *Handlers) GetScripts(w http.ResponseWriter, r *http.Request) {
	scripts, err := handler.Service.ScriptService.GetScripts()
	if err != nil {
		handler.Service.ScriptService.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		loggers.InfoLog.Println("Internal Server Error")
		return
	}
	jsonData, err := json.Marshal(scripts)
	if err != nil {
		handler.Service.ScriptService.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		loggers.InfoLog.Println("Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	loggers.DebugLog.Println("GetScripts completed successfully")
}
