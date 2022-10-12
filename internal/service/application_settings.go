package service

import (
	"3dPrintCalc/internal/domain"
	"3dPrintCalc/internal/repository"
)

type ApplicationSettingsService struct {
	applicationSettingsRepo repository.ApplicationSettings
}

func NewApplicationSettingsService(applicationSettingsRepo repository.ApplicationSettings) *ApplicationSettingsService {
	return &ApplicationSettingsService{applicationSettingsRepo: applicationSettingsRepo}
}

func (s *ApplicationSettingsService) GetAll() ([]domain.ApplicationSetting, error) {
	return s.applicationSettingsRepo.GetAll()
}
