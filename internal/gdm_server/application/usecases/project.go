package usecases

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/repository"
)

func GetProjectByID(repo repository.IProjectRepo, id int) (dto.Project, error) {
	res, err := repo.GetByID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GetProjectByName(repo repository.IProjectRepo, name string) (dto.Project, error) {
	res, err := repo.GetByContainerName(name)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GetProjectsByFilter(repo repository.IProjectRepo, filter repository.Filter) (dto.Projects, error) {
	res, err := repo.GetProjects(filter)
	if err != nil {
		return res, err
	}
	return res, err
}

func CreateProject(repo repository.IProjectRepo, proj dto.CreateProject) (dto.Project, error) {
	created, err := repo.CreateProject(proj)
	if err != nil {
		return dto.Project{}, err
	}
	return created, nil
}

func UpdateProject(repo repository.IProjectRepo, proj dto.Project) (dto.Project, error) {
	updated, err := repo.UpdateProject(proj)
	if err != nil {
		return dto.Project{}, err
	}
	return updated, nil
}

func DeleteByID(repo repository.IProjectRepo, id int) error {
	if err := repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
