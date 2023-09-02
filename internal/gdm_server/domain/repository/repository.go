package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
)

type Filter struct {
	Limit   int
	Offset  int
	Deleted bool
}

type IProjectRepo interface {
	GetByID(int) (dto.Project, error)
	GetByContainerName(string) (dto.Project, error)
	GetProjects(Filter) (dto.Projects, error)
	CreateProject(dto.Project) (dto.Project, error)
	UpdateProject(dto.Project) (dto.Project, error)
	DeleteByID(int) error
	getLatest() (dto.Project, error)
}
