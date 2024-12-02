package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	ID            uint
	BorrowingDate *time.Time
	Deadline      *time.Time
	RetrievalDate *time.Time
	Status        int
	Units         []*Unit `gorm:"many2many:borrow_unit;"`
}

func BorrowCreate() *Borrow {
	borrow := Borrow{Status: -1}
	//borrow := Borrow{BorrowingDate: &borrowdate, Deadline: borrowdate.AddDate(0, 0, 7), Status: -1}
	DB.Create(&borrow)

	/*for i := range un {
		un[i].Borrowed = true
	}

	DB.Model(&borrow).Association("Units").Append(un)*/
	return &borrow
}

func BorowUpdateCart(id uint64, un Unit) *Borrow {
	var borrow Borrow
	DB.Where("id = ?", id).First(&borrow)
	DB.Model(&borrow).Association("Units").Append(un)
	return &borrow
}

func BorowUpdateOnGoing(id uint64) *Borrow {
	var borrow Borrow
	DB.Where("id = ?", id).First(&borrow)
	borrowdate := time.Now()
	deadline := borrowdate.AddDate(0, 0, 7)
	DB.Model(&borrow).Updates(Borrow{BorrowingDate: &borrowdate, Deadline: &deadline, Status: 0})
	for i := range borrow.Units {
		borrow.Units[i].Borrowed = true
		DB.Save(&borrow.Units[i])
	}
	return &borrow
}

func BorrowFinish(id uint64) *Borrow {
	var borrow Borrow
	DB.Where("id = ?", id).Preload("Units").First(&borrow)
	retrievedate := time.Now()
	for i := range borrow.Units {
		borrow.Units[i].Borrowed = false
		DB.Save(&borrow.Units[i])
	}
	DB.Model(&borrow).Updates(Borrow{Status: 1, RetrievalDate: &retrievedate})
	return &borrow
}
