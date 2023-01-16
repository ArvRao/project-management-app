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

type TodoApiStruct struct {
	TodoName string `json:"todos_name"`
	Priority int    `json:"priority"`
	// Status      int        `json:"status"`
	ProjectIdFk uuid.UUID  `json:"project_id_fk"`
	Project     ProjectApi `gorm:"foreignKey:ProjectIdFk"`
}
