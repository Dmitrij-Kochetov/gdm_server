package repository

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/models"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/sqlite"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/repository"
)

type ProjectRepository struct {
	storage *sqlite.Storage
}

func NewProjectRepo(storage *sqlite.Storage) *ProjectRepository {
	return &ProjectRepository{
		storage: storage,
	}
}

func (p *ProjectRepository) GetByID(id int) (dto.Project, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Project
	if err := p.storage.DB.Get(&proj, `SELECT * FROM "projects" WHERE ID=$1`, id); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}

func (p *ProjectRepository) GetByContainerName(name string) (dto.Project, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Project
	if err := p.storage.DB.Get(&proj, `SELECT * FROM "projects" WHERE container_name=$1`, name); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}

func (p *ProjectRepository) GetProjects(filter repository.Filter) (dto.Projects, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Projects
	if err := p.storage.DB.Select(&proj,
		`SELECT * FROM projects 
         	WHERE deleted=$1
         	ORDER BY ID
         	LIMIT $2 OFFSET $3`,
		filter.Deleted, filter.Limit, filter.Offset); err != nil {
		return dto.Projects{}, err
	}
	return proj.ConvertToDomain(), nil
}

func (p *ProjectRepository) CreateProject(proj dto.CreateProject) (dto.Project, error) {
	p.storage.Lock()
	defer p.storage.Unlock()

	tx, err := p.storage.DB.Begin()
	if err != nil {
		return dto.Project{}, err
	}

	_, err = tx.Exec(
		`INSERT INTO "projects" (link, container_name, up_to_date, running, deleted)
			VALUES ($1, $2, $3, $4, $5)`,
		proj.Link,
		proj.ContainerName,
		proj.UpToDate,
		proj.Running,
		proj.Deleted,
	)
	if err != nil {
		return dto.Project{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dto.Project{}, err
	}

	return p.getLatest()
}

func (p *ProjectRepository) UpdateProject(proj dto.Project) (dto.Project, error) {
	p.storage.Lock()
	defer p.storage.Unlock()

	tx, err := p.storage.DB.Begin()
	if err != nil {
		return dto.Project{}, err
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
		return dto.Project{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dto.Project{}, err
	}

	proj, err = p.GetByID(proj.ID)
	if err != nil {
		return dto.Project{}, err
	}

	return proj, nil
}

func (p *ProjectRepository) DeleteByID(id int) error {
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

func (p *ProjectRepository) getLatest() (dto.Project, error) {
	p.storage.RLock()
	defer p.storage.RUnlock()

	var proj models.Project
	if err := p.storage.DB.Get(&proj, `SELECT * FROM "projects" ORDER BY ID DESC LIMIT 1`); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}
