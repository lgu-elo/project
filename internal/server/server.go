package server

import "github.com/lgu-elo/project/internal/project"

type (
	ProjectHandler project.IHandler

	API struct {
		ProjectHandler
	}
)

func NewAPI(project project.IHandler) *API {
	return &API{project}
}
