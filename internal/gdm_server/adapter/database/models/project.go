package models

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/project"
	"github.com/google/uuid"
)

type Project struct {
	ID            []byte `db:"ID"`
	Link          string `db:"link"`
	ContainerName string `db:"container_name"`
	UpToDate      bool   `db:"up_to_date"`
	Running       bool   `db:"running"`
	Deleted       bool   `db:"deleted"`
}

type Projects struct {
	Projects []Project
}

func (r Project) ConvertToDomain() (project.Project, error) {

	converted, err := uuid.FromBytes(r.ID)
	if err != nil {
		return project.Project{}, err
	}

	return project.Project{
		ID:            converted,
		Link:          r.Link,
		ContainerName: r.ContainerName,
		UpToDate:      r.UpToDate,
		Running:       r.Running,
		Deleted:       r.Deleted,
	}, nil
}
