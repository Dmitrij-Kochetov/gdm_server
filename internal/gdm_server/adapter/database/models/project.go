package models

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
)

type Projects []Project

func (p Projects) ConvertToDomain() dto.Projects {
	proj := dto.Projects{}
	for _, val := range p {
		val := dto.Project(val)
		proj = append(proj, val)
	}
	return proj
}

type Project struct {
	ID            int    `db:"ID"`
	Link          string `db:"link"`
	ContainerName string `db:"container_name"`
	UpToDate      bool   `db:"up_to_date"`
	Running       bool   `db:"running"`
	Deleted       bool   `db:"deleted"`
}
