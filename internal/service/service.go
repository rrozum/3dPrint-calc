package service

import (
	"3dPrintCalc/internal/domain"
	"3dPrintCalc/internal/repository"
)

type ApplicationSettings interface {
	GetAll() ([]domain.ApplicationSetting, error)
}

type Projects interface {
	GetAll() ([]domain.Project, error)
}

type Details interface {
	GetByProject(projectId int) ([]domain.Detail, error)
}

type Services struct {
	ApplicationSettings ApplicationSettings
	Projects            Projects
	Details             Details
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		ApplicationSettings: NewApplicationSettingsService(deps.Repos.ApplicationSettings),
		Projects:            NewProjectsService(deps.Repos.Projects),
		Details:             NewDetailsService(deps.Repos.Details),
	}
}
