package repository

import (
	"menusPlanningApp/api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DishRepository struct {
	db *gorm.DB
}

func NewDishRepository(db *gorm.DB) *DishRepository {
	return &DishRepository{db}
}

func (r *DishRepository) FindAll(params model.QueryParams) ([]model.Dish, error) {
	var dishes []model.Dish
	query := r.db

	for key, value := range params.Filters {
		query = query.Where(key+" = ?", value)
	}
	if params.SortBy != "" {
		order := params.SortBy
		if params.SortOrder == "desc" {
			order += " DESC"
		} else {
			order += " ASC"
		}
		query = query.Order(order)
	}

	if err := query.Limit(params.Limit).Offset(params.Offset).Find(&dishes).Error; err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *DishRepository) Create(dish *model.Dish) (*model.Dish, error) {
	if err := r.db.Create(dish).Error; err != nil {
		return nil, err
	}
	return dish, nil
}

func (r *DishRepository) FindByID(id uuid.UUID) (*model.Dish, error) {
	var dish model.Dish
	if err := r.db.Where("id = ?", id).First(&dish).Error; err != nil {
		return nil, err
	}
	return &dish, nil
}

func (r *DishRepository) Update(dish *model.Dish) (*model.Dish, error) {
	if err := r.db.Save(dish).Error; err != nil {
		return nil, err
	}
	return dish, nil
}

func (r *DishRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Dish{}).Error; err != nil {
		return err
	}
	return nil
}
