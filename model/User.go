package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID  uint `gorm:"primary_key"`
	Name string `gorm: "type:varchar(255);unique_index`
	Email string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm: "size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

