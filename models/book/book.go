package book

import (
	p "fp_pbkk/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
)

type Book struct {
	gorm.Model
	ID    int64
	Title string
}

func All(db *gorm.DB, ctx *gin.Context) *[]Book {
	var books []Book
	db.Where("deleted_at is NULL").Scopes(p.Paginate(ctx)).Order("updated_at desc").Find(&books)
	fmt.Println(books)
	return &books
}

func Find(db *gorm.DB, id uint64) *Book {
	var book Book
	db.Where("id = ?", id).First(&book)
	return &book
}
