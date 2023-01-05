package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID       int
	TodoName string
	Priority int
	Status   int
}