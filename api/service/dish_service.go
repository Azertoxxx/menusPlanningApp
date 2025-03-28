package service

import (
	"menusPlanningApp/api/model"
	"menusPlanningApp/api/repository"

	"github.com/google/uuid"
)

type DishService struct {
	repo *repository.DishRepository
}

func NewDishService(repo *repository.DishRepository) *DishService {
	return &DishService{repo}
}

func (s *DishService) FindAll(params model.QueryParams) ([]model.Dish, error) {
	return s.repo.FindAll(params)
}

func (s *DishService) Create(dish *model.Dish) (*model.Dish, error) {
	return s.repo.Create(dish)
}

func (s *DishService) FindByID(id uuid.UUID) (*model.Dish, error) {
	return s.repo.FindByID(id)
}

func (s *DishService) Update(dish *model.Dish) (*model.Dish, error) {
	return s.repo.Update(dish)
}

func (s *DishService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
