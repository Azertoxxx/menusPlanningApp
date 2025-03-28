package service

import (
	"menusPlanningApp/api/model"
	"menusPlanningApp/api/repository"

	"github.com/google/uuid"
)

type PlanningService struct {
	repo *repository.PlanningRepository
}

func NewPlanningService(repo *repository.PlanningRepository) *PlanningService {
	return &PlanningService{repo}
}

func (s *PlanningService) FindAll() ([]model.Planning, error) {
	return s.repo.FindAll()
}

func (s *PlanningService) Create(planning *model.Planning) (*model.Planning, error) {
	return s.repo.Create(planning)
}

func (s *PlanningService) FindByID(id uuid.UUID) (*model.Planning, error) {
	return s.repo.FindByID(id)
}

func (s *PlanningService) Update(planning *model.Planning) (*model.Planning, error) {
	return s.repo.Update(planning)
}

func (s *PlanningService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
