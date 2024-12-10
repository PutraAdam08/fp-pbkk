package models

import (
	p "fp_pbkk/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
)

type Book struct {
	gorm.Model
	ID          uint64
	Title       string
	PublishYear string
	ISBN        string
	Category    string
	Units       []Unit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func BooksAll(ctx *gin.Context) *[]Book {
	var books []Book
	DB.Where("deleted_at is NULL").Scopes(p.Paginate(ctx)).Order("updated_at desc").Find(&books)
	fmt.Println(books)
	return &books
}

func BooksCategoryFilter(ctx *gin.Context, find string) *[]Book {
	var books []Book
	category := "%" + find + "%"
	DB.Where("deleted_at is NULL AND category LIKE ?", category).Scopes(p.Paginate(ctx)).Order("updated_at desc").Find(&books)
	fmt.Println(books)
	return &books
}

func BookFind(id uint64) *Book {
	var book Book
	DB.Where("id = ?", id).First(&book)
	return &book
}

func BookCreate(title string, publishYear string, ISBN string, cat string) *Book {
	book := Book{Title: title, PublishYear: publishYear, ISBN: ISBN, Category: cat}
	DB.Create(&book)
	return &book
}

func BookUpdate(id uint64, title string, publishYear string, ISBN string, cat string) *Book {
	var book Book
	DB.Model(&book).Where("id = ?", id).Updates(Book{Title: title, PublishYear: publishYear, ISBN: ISBN, Category: cat})
	return &book
}

func BookDelete(id uint64) *gorm.DB {
	var book Book
	result := DB.Where("id = ?", id).Delete(&book)
	return result
}
