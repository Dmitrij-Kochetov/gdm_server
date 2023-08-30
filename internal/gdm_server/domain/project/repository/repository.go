package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/project"
	"github.com/google/uuid"
)

type IProject interface {
	GetByID(uuid.UUID) (project.Project, error)
	GetByContainerName(string) (project.Project, error)
	CreateProject(project.Project) error
	UpdateProject(project.Project) error
	DeleteById(uuid.UUID) error
}
