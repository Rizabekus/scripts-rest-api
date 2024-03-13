package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Rizabekus/scripts-rest-api/pkg/loggers"
	"github.com/gorilla/mux"
)

func (handler *Handlers) GetOneScript(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to GetOneScript")
	vars := mux.Vars(r)
	ScriptID := vars["id"]
	userData, err := handler.Service.ScriptService.GetScript(ScriptID)
	if err != nil {
		if err == sql.ErrNoRows {
			handler.Service.ScriptService.SendResponse("Not Found", w, http.StatusNotFound)
			loggers.InfoLog.Println("ScriptID Not Found")
			return
		}
		handler.Service.ScriptService.SendResponse("Internal Server Error", w, http.StatusInternalServerError)
		loggers.InfoLog.Println("Internal Server Error")
		return
	}
	jsonData, err := json.Marshal(userData)
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
