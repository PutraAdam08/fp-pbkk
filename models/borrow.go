package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	ID            uint
	BorrowingDate time.Time
	Deadline      time.Time
	RetrievalDate *time.Time
	Status        int
	Unit          []Unit
}

func BorrowCreate() *Borrow {
	borrowdate := time.Now()
	borrow := Borrow{BorrowingDate: borrowdate, Deadline: borrowdate.AddDate(0, 0, 7), Status: 0}
	DB.Create(&borrow)
	return &borrow
}

func BorrowFinish(id uint64) *Borrow {
	var borrow Borrow
	retrievedate := time.Now()
	DB.Model(&borrow).Where("id = ?", id).Updates(Borrow{Status: 1, RetrievalDate: &retrievedate})
	return &borrow
}
