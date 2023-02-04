package repository

import (
	"3dPrintCalc/internal/domain"
	"database/sql"
)

type Repositories struct {
	ApplicationSettings ApplicationSettings
	Projects            Projects
	Details             Details
}

type ApplicationSettings interface {
	Save(applicationSetting domain.ApplicationSetting) error
	GetAll() ([]domain.ApplicationSetting, error)
}

type Projects interface {
	Save(project domain.Project) error
	GetAll() ([]domain.Project, error)
}

type Details interface {
	Save(detail domain.Detail) error
	GetByProject(projectId int) ([]domain.Detail, error)
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		ApplicationSettings: NewApplicationSettingsRepo(db),
		Projects:            NewProjectsRepo(db),
		Details:             NewDetailsRepo(db),
	}
}
