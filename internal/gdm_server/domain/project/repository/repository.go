package repository

import (
	"context"
	"github.com/Dmitrij-Kochetov/gdm_server/internal/gdm_server/domain/project"
)

type ProjectRepository interface {
	Get(context.Context, int) (project.Project, error)
}
