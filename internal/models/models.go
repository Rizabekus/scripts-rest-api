package models

import "net/http"

type ScriptService interface {
	AddScript(script string) error
	SendResponse(response string, w http.ResponseWriter, statusCode int)
	GetScripts() ([]Command, error)
	GetScript(id string) (Command, error)
}

type ScriptStorage interface {
	AddScript(script string) error
	GetScripts() ([]Command, error)
	GetScript(id string) (Command, error)
}

type Command struct {
	Id     int    `json:"id"`
	Script string `json:"script" validate:"bash"`
}

type ErrorResponse struct {
	Details string
}
