package models

import (
	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	ID       uint64
	Code     string
	Borrowed bool
	BookID   uint64
	Borrow   []*Borrow `gorm:"many2many:user_languages;"`
}

func UnitCreate(code string) *Unit {
	unit := Unit{Code: code, Borrowed: false}
	DB.Create(&unit)
	return &unit
}

func UnitStatusOnBorrow(id uint64) *Unit {
	unit := Unit{Borrowed: true}
	DB.Model(&unit).Where("id = ?", id).Updates(Unit{Borrowed: true})
	return &unit
}

func UnitStatusOnFree(id uint64) *Unit {
	unit := Unit{Borrowed: false}
	DB.Model(&unit).Where("id = ?", id).Updates(Unit{Borrowed: true})
	return &unit
}
