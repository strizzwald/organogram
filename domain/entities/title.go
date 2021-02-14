package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Title struct {
	gorm.Model
	Id   int32
	Guid uuid.UUID
	Name string
}
