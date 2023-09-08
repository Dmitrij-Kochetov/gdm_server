package usecases

import (
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/dto"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/repository"
)

func GetVersionsByProjectID(repo repository.IVersionRepo, id int, filter repository.Filter) (dto.Versions, error) {
	res, err := repo.GetByID(id, filter)
	if err != nil {
		return dto.Versions{}, err
	}
	return res, nil
}

func CreateVersion(repo repository.IVersionRepo, version dto.CreateVersion) (dto.Version, error) {
	res, err := repo.CreateVersion(version)
	if err != nil {
		return dto.Version{}, err
	}
	return res, nil
}

func UpdateVersion(repo repository.IVersionRepo, version dto.Version) (dto.Version, error) {
	res, err := repo.UpdateVersion(version)
	if err != nil {
		return dto.Version{}, err
	}
	return res, nil
}

func DeleteVersionByID(repo repository.IVersionRepo, id int) error {
	if err := repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}

func DeleteVersionsByProjectID(repo repository.IVersionRepo, id int) error {
	if err := repo.DeleteByProjectID(id); err != nil {
		return err
	}
	return nil
}
