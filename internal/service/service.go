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

type Services struct {
	ApplicationSettings ApplicationSettings
	Projects            Projects
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		ApplicationSettings: NewApplicationSettingsService(deps.Repos.ApplicationSettings),
		Projects:            NewProjectsService(deps.Repos.Projects),
	}
}
