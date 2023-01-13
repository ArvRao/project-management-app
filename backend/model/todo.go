package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	TodoName    string
	Priority    int
	Status      int
	ProjectIdFk uuid.UUID
	Project     Project `gorm:"foreignKey:ProjectIdFk"`
}
