package service

import (
	"3dPrintCalc/internal/domain"
	"3dPrintCalc/internal/repository"
)

type DetailsService struct {
	detailsRepo repository.Details
}

func NewDetailsService(detailsRepo repository.Details) *DetailsService {
	return &DetailsService{detailsRepo: detailsRepo}
}

func (s *DetailsService) GetByProject(projectId int) ([]domain.Detail, error) {
	return s.detailsRepo.GetByProject(projectId)
}
