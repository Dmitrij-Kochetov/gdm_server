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

func CreateProject(repo repository.IProjectRepo, proj dto.Project) error {
	if err := repo.CreateProject(proj); err != nil {
		return err
	}
	return nil
}

func UpdateProject(repo repository.IProjectRepo, proj dto.Project) error {
	if err := repo.UpdateProject(proj); err != nil {
		return err
	}
	return nil
}

func DeleteByID(repo repository.IProjectRepo, id int) error {
	if err := repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
