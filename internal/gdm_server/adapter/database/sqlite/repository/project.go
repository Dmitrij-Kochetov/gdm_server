package repository

import (
	"context"
	"fmt"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/adapter/database/models"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/repository"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepo(path string) (*ProjectRepository, error) {
	const op = "adapter.database.sqlite.repository.project"

	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ProjectRepository{db: db}, nil
}

func (p *ProjectRepository) Close(ctx context.Context) error {
	const op = "adapter.database.sqlite.Close"
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (p *ProjectRepository) GetByID(id int) (dto.Project, error) {
	var proj models.Project
	if err := p.db.Get(&proj, `SELECT * FROM projects WHERE project_id=$1`, id); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}

func (p *ProjectRepository) GetByContainerName(name string) (dto.Project, error) {
	var proj models.Project
	if err := p.db.Get(&proj, `SELECT * FROM "projects" WHERE container_name=$1`, name); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}

func (p *ProjectRepository) GetProjects(filter repository.Filter) (dto.Projects, error) {
	var proj models.Projects
	if err := p.db.Select(&proj,
		`SELECT * FROM projects 
         	WHERE deleted=$1
         	ORDER BY project_id
         	LIMIT $2 OFFSET $3`,
		filter.Deleted, filter.Limit, filter.Offset); err != nil {
		return dto.Projects{}, err
	}
	return proj.ConvertToDomain(), nil
}

func (p *ProjectRepository) CreateProject(proj dto.CreateProject) (dto.Project, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return dto.Project{}, err
	}

	_, err = tx.Exec(
		`INSERT INTO "projects" (link, container_name, description, up_to_date, deleted)
			VALUES ($1, $2, $3, $4, $5)`,
		proj.Link,
		proj.ContainerName,
		proj.Description,
		proj.UpToDate,
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
	tx, err := p.db.Begin()
	if err != nil {
		return dto.Project{}, err
	}

	_, err = tx.Exec(
		`UPDATE "projects"
			SET link=$1, container_name=$2, description=$3, up_to_date=$4, deleted=$5
			WHERE project_id=$6`,
		proj.Link,
		proj.ContainerName,
		proj.Description,
		proj.UpToDate,
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
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE projects SET deleted=TRUE WHERE project_id=$1`, id)
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
	var proj models.Project
	if err := p.db.Get(&proj, `SELECT * FROM "projects" ORDER BY project_id DESC LIMIT 1`); err != nil {
		return dto.Project{}, err
	}

	return dto.Project(proj), nil
}
