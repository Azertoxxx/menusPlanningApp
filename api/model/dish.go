package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dish struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id"`
	Name        string         `gorm:"index" json:"name"`
	ImageURL    string         `json:"imageUrl"`
	Ingredients string         `json:"ingredients"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func (d *Dish) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
