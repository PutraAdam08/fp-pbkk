package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fmt"
)

type Category struct {
	gorm.Model
	ID   uint64
	Name string
}

func CategoriesAll(ctx *gin.Context) *[]Category {
	var categories []Category
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&categories)
	fmt.Println(categories)
	return &categories
}

func CategoryFind(id uint64) *Category {
	var category Category
	DB.Where("id = ?", id).First(&category)
	return &category
}

func CategoryCreate(name string) *Category {
	category := Category{Name: name}
	DB.Create(&category)
	return &category
}

func CategoryUpdate(id uint64, name string) *Category {
	var category Category
	DB.Model(&category).Where("id = ?", id).Updates(Category{Name: name})
	return &category
}

func CategoryDelete(id uint64) *gorm.DB {
	result := DB.Unscoped().Delete(&id)
	return result
}
