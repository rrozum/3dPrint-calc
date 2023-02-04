package service

import (
	"3dPrintCalc/internal/domain"
	"3dPrintCalc/internal/repository"
)

type ProjectsService struct {
	projectsRepo repository.Projects
}

func NewProjectsService(projectsRepo repository.Projects) *ProjectsService {
	return &ProjectsService{projectsRepo: projectsRepo}
}

func (s *ProjectsService) GetAll() ([]domain.Project, error) {
	return s.projectsRepo.GetAll()
}
