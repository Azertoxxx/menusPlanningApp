package repository

import (
	"menusPlanningApp/api/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlanningRepository struct {
	db *gorm.DB
}

func NewPlanningRepository(db *gorm.DB) *PlanningRepository {
	return &PlanningRepository{db}
}

func (r *PlanningRepository) FindAll() ([]model.Planning, error) {
	var plannings []model.Planning
	if err := r.db.Preload("Dish").Find(&plannings).Error; err != nil {
		return nil, err
	}
	return plannings, nil
}

func (r *PlanningRepository) Create(planning *model.Planning) (*model.Planning, error) {
	if err := r.db.Create(planning).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Dish").First(planning, "id = ?", planning.ID).Error; err != nil {
		return nil, err
	}
	return planning, nil
}

func (r *PlanningRepository) FindByID(id uuid.UUID) (*model.Planning, error) {
	var planning model.Planning
	err := r.db.Preload("Dish").First(&planning, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &planning, nil
}

func (r *PlanningRepository) Update(planning *model.Planning) (*model.Planning, error) {
	if err := r.db.Save(planning).Error; err != nil {
		return nil, err
	}
	return planning, nil
}

func (r *PlanningRepository) Delete(id uuid.UUID) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Planning{}).Error; err != nil {
		return err
	}
	return nil
}
