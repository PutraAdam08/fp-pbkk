package models

import (
	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	ID       uint
	Code     string
	Borrowed bool
	BookID   uint
}
