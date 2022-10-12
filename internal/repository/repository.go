package repository

import (
	"3dPrintCalc/internal/domain"
	"database/sql"
)

type Repositories struct {
	ApplicationSettings ApplicationSettings
}

type ApplicationSettings interface {
	Save(applicationSetting domain.ApplicationSetting) error
	GetAll() ([]domain.ApplicationSetting, error)
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		ApplicationSettings: NewApplicationSettingsRepo(db),
	}
}
