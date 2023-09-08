package repository

import "github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"

type IVersionRepo interface {
	GetByID(int, Filter) (dto.Versions, error)
	CreateVersion(dto.CreateVersion) (dto.Version, error)
	UpdateVersion(dto.Version) (dto.Version, error)
	DeleteByID(int) error
	DeleteByProjectID(int) error
	GetLatest(int) (dto.Version, error)
	GetVersionByID(int) (dto.Version, error)
}
