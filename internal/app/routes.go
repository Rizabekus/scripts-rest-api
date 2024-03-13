package app

import (
	"log"
	"net/http"

	"github.com/Rizabekus/scripts-rest-api/internal/handlers"
	"github.com/Rizabekus/scripts-rest-api/pkg/loggers"
	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/scripts", h.AddScript).Methods("POST")
	r.HandleFunc("/scripts", h.GetScripts).Methods("GET")
	r.HandleFunc("/scripts/{id}", h.GetOneScript).Methods("GET")
	loggers.InfoLog.Println("Started the server at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
