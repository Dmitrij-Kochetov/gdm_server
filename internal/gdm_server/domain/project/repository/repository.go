package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/project"
	"github.com/google/uuid"
)

type Filter struct {
	Limit   int
	Offset  int
	Deleted bool
}

type IProjectRepo interface {
	GetByID(uuid.UUID) (project.Project, error)
	GetByContainerName(string) (project.Project, error)
	GetProjects(Filter) (project.Projects, error)
	CreateProject(project.Project) error
	UpdateProject(project.Project) error
	DeleteById(uuid.UUID) error
}
