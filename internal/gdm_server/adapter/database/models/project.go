package models

import "github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"

type Project struct {
	ID            int    `db:"project_id"`
	Link          string `db:"link"`
	ContainerName string `db:"container_name"`
	Description   string `db:"description"`
	UpToDate      bool   `db:"up_to_date"`
	Deleted       bool   `db:"deleted"`
}

type Projects []Project

func (p Projects) ConvertToDomain() dto.Projects {
	proj := dto.Projects{}
	for _, val := range p {
		val := dto.Project(val)
		proj = append(proj, val)
	}
	return proj
}
