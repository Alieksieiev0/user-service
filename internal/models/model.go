package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}

type User struct {
	Base
	Username string `gorm:"default:null;not null;unique;"`
	Password string `gorm:"default:null;not null;"`
	Email    string `gorm:"default:null;not null;unique;"`
}
