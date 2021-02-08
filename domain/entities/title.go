package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Title struct {
	gorm.Model
	Id   uuid.UUID
	Name string
}
