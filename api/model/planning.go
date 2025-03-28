package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Planning struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Slot   string    `gorm:"index" json:"slot"`
	DishID uuid.UUID `gorm:"type:uuid;index" json:"dishId"`
	Dish   Dish      `gorm:"foreignKey:DishID;references:ID;constraint:OnDelete:CASCADE"`
}

func (d *Planning) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
