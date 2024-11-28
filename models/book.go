package models

import (
	p "fp_pbkk/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
)

type Book struct {
	gorm.Model
	ID    uint
	Title string
	Units []Unit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func BooksAll(ctx *gin.Context) *[]Book {
	var books []Book
	DB.Where("deleted_at is NULL").Scopes(p.Paginate(ctx)).Order("updated_at desc").Find(&books)
	fmt.Println(books)
	return &books
}

func BookFind(id uint64) *Book {
	var book Book
	DB.Where("id = ?", id).First(&book)
	return &book
}

func BookAdd(title string) {
	book := Book{Title: title}
	DB.Create(&book)
}

func BookEdit(id uint, title string) {
	var book Book
	DB.Model(&book).Where("id = ?", id).Updates(Book{Title: title})
}
