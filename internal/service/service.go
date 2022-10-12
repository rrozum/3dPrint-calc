package service

import (
	"3dPrintCalc/internal/domain"
	"3dPrintCalc/internal/repository"
)

type ApplicationSettings interface {
	GetAll() ([]domain.ApplicationSetting, error)
}

type Services struct {
	ApplicationSettings ApplicationSettings
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		ApplicationSettings: NewApplicationSettingsService(deps.Repos.ApplicationSettings),
	}
}
