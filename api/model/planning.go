package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Slot string

const (
	MONDAY_LUNCH     Slot = "MONDAY_LUNCH"
	MONDAY_DINNER    Slot = "MONDAY_DINNER"
	TUESDAY_LUNCH    Slot = "TUESDAY_LUNCH"
	TUESDAY_DINNER   Slot = "TUESDAY_DINNER"
	WEDNESDAY_LUNCH  Slot = "WEDNESDAY_LUNCH"
	WEDNESDAY_DINNER Slot = "WEDNESDAY_DINNER"
	THURSDAY_LUNCH   Slot = "THURSDAY_LUNCH"
	THURSDAY_DINNER  Slot = "THURSDAY_DINNER"
	FRIDAY_LUNCH     Slot = "FRIDAY_LUNCH"
	FRIDAY_DINNER    Slot = "FRIDAY_DINNER"
	SATURDAY_LUNCH   Slot = "SATURDAY_LUNCH"
	SATURDAY_DINNER  Slot = "SATURDAY_DINNER"
	SUNDAY_LUNCH     Slot = "SUNDAY_LUNCH"
	SUNDAY_DINNER    Slot = "SUNDAY_DINNER"
)

func (s Slot) IsValid() bool {
	switch s {
	case
		MONDAY_LUNCH, MONDAY_DINNER,
		TUESDAY_LUNCH, TUESDAY_DINNER,
		WEDNESDAY_LUNCH, WEDNESDAY_DINNER,
		THURSDAY_LUNCH, THURSDAY_DINNER,
		FRIDAY_LUNCH, FRIDAY_DINNER,
		SATURDAY_LUNCH, SATURDAY_DINNER,
		SUNDAY_LUNCH, SUNDAY_DINNER:
		return true
	default:
		return false
	}
}

type Planning struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Slot   Slot      `gorm:"index; not null" json:"slot"`
	DishID uuid.UUID `gorm:"type:uuid;index" json:"dishId"`
	Dish   Dish      `gorm:"foreignKey:DishID;references:ID;constraint:OnDelete:CASCADE" json:"dish"`
}

func (d *Planning) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
