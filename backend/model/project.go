package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project struct
// add other fields like about and description of project
type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	ProjectName string    `json:"projectname"`
}

// Projects struct
type Projects struct {
	Projects []Project `json:"projects"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	project.ID = uuid.New()
	return
}
