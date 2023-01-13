package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 1 project can have many todos
type Project struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;"`
	ProjectName string    `json:"projectname"`
	About       string
}

// Projects struct
type Projects struct {
	Projects []Project `json:"projects"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	project.ID = uuid.New()
	return
}
