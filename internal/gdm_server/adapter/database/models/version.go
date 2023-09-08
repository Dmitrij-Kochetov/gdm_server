package models

import "github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"

type Version struct {
	ID        int    `db:"version_id"`
	ProjectID int    `db:"project_id"`
	Version   string `db:"version"`
	Config    string `db:"config"`
	Deleted   bool   `db:"deleted"`
}

type Versions []Version

func (v Versions) ConvertToDomain() dto.Versions {
	ver := dto.Versions{}
	for _, val := range v {
		val := dto.Version(val)
		ver = append(ver, val)
	}
	return ver
}
