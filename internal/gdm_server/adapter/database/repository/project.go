package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/models"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/sqlite"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/project"
	"github.com/google/uuid"
)

type ProjectRepository struct {
	storage *sqlite.Storage
}

func New(storage *sqlite.Storage) *ProjectRepository {
	return &ProjectRepository{
		storage: storage,
	}
}

func (p *ProjectRepository) GetByID(id uuid.UUID) (project.Project, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Project
	if err := p.storage.DB.Get(&proj, `SELECT * FROM "projects" WHERE ID=$1`, id); err != nil {
		return project.Project{}, err
	}

	return proj.ConvertToDomain()
}

func (p *ProjectRepository) GetByContainerName(name string) (project.Project, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Project
	if err := p.storage.DB.Get(`SELECT * FROM "projects" WHERE container_name=$1`, name); err != nil {
		return project.Project{}, err
	}

	return proj.ConvertToDomain()
}

func (p *ProjectRepository) CreateProject(proj project.Project) error {
	p.storage.Lock()
	defer p.storage.Unlock()

	tx, err := p.storage.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`INSERT INTO "projects" (ID, link, container_name, up_to_date, running, deleted)
			VALUES ($1, $2, $3, $4, $5, $6)`,
		proj.ID,
		proj.Link,
		proj.ContainerName,
		proj.UpToDate,
		proj.Running,
		proj.Deleted,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectRepository) UpdateProject(proj project.Project) error {
	p.storage.Lock()
	defer p.storage.Unlock()

	tx, err := p.storage.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		`UPDATE "projects"
			SET link=$1, container_name=$2, up_to_date=$3, running=$4, deleted=$5
			WHERE ID=$6`,
		proj.Link,
		proj.ContainerName,
		proj.UpToDate,
		proj.Running,
		proj.Deleted,
		proj.ID,
	)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectRepository) DeleteProject(id uuid.UUID) error {
	p.storage.Lock()
	defer p.storage.Unlock()

	tx, err := p.storage.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM "projects" WHERE ID=$1`, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
