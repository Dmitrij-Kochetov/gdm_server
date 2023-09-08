package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/models"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/repository"
	"github.com/jmoiron/sqlx"
)

type VersionRepository struct {
	db *sqlx.DB
}

func NewVersionRepo(db *sqlx.DB) *VersionRepository {
	return &VersionRepository{db: db}
}

func (v *VersionRepository) GetByID(projectId int, filter repository.Filter) (dto.Versions, error) {
	var ver models.Versions
	if err := v.db.Select(&ver,
		`SELECT * FROM versions 
         WHERE project_id=$1 AND deleted=$2
         ORDER BY version_id
         LIMIT $3 OFFSET $4`,
		projectId,
		filter.Deleted,
		filter.Limit,
		filter.Offset,
	); err != nil {
		return dto.Versions{}, err
	}
	return ver.ConvertToDomain(), nil
}

func (v *VersionRepository) CreateVersion(version dto.CreateVersion) (dto.Version, error) {
	tx, err := v.db.Begin()
	if err != nil {
		return dto.Version{}, err
	}

	_, err = tx.Exec(
		`INSERT INTO versions (project_id, version, config) VALUES ($1, $2, $3)`,
		version.ProjectID,
		version.Version,
		version.Config,
	)
	if err != nil {
		return dto.Version{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dto.Version{}, err
	}
	return v.GetLatest(version.ProjectID)
}

func (v *VersionRepository) UpdateVersion(version dto.Version) (dto.Version, error) {
	tx, err := v.db.Begin()
	if err != nil {
		return dto.Version{}, err
	}

	_, err = tx.Exec(
		`UPDATE versions SET config=$1, version=$2 WHERE version_id=$3`,
		version.Config,
		version.Version,
		version.ID,
	)
	if err != nil {
		return dto.Version{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dto.Version{}, err
	}

	ver, err := v.GetVersionByID(version.ID)
	if err != nil {
		return dto.Version{}, err
	}
	return ver, nil
}

func (v *VersionRepository) DeleteByID(versionID int) error {
	tx, err := v.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		`UPDATE versions SET deleted=TRUE WHERE version_id=$1`, versionID,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (v *VersionRepository) DeleteByProjectID(projectID int) error {
	tx, err := v.db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(
		`UPDATE versions SET deleted=TRUE WHERE project_id=$1`,
		projectID,
	); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (v *VersionRepository) GetLatest(projectID int) (dto.Version, error) {
	var ver models.Version
	if err := v.db.Get(&ver,
		`SELECT * FROM versions WHERE project_id=$1 ORDER BY version_id DESC LIMIT 1`,
		projectID,
	); err != nil {
		return dto.Version{}, err
	}
	return dto.Version(ver), nil
}

func (v *VersionRepository) GetVersionByID(versionID int) (dto.Version, error) {
	var ver models.Version
	if err := v.db.Get(&ver,
		`SELECT * FROM versions WHERE version_id=$1`, versionID,
	); err != nil {
		return dto.Version{}, err
	}
	return dto.Version(ver), nil
}
