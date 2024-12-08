package controllers

import (
	"fmt"
	"fp_pbkk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type formData struct {
	Title       string `form:"title"`
	PublishYear string `form:"publishYear"`
	ISBN        string `form:"isbn"`
}

func MainPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"books/main.tpl",
		gin.H{},
	)
}

func BookIndex(c *gin.Context) {
	books := models.BooksAll(c)
	//categories := models.CategoriesAll(c)
	c.HTML(
		http.StatusOK,
		"books/index.tpl",
		gin.H{
			"books":      books,
			"page":       c.GetInt("page"),
			"pageSize":   c.GetInt("pageSize"),
			"totalPages": c.GetInt("totalPages"),
		},
	)
}

func BookShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	book := models.BookFind(id)
	c.HTML(
		http.StatusOK,
		"books/show.tpl",
		gin.H{
			"book": book,
		},
	)
}

func BookAdd(c *gin.Context) {
	var data formData
	c.Bind(&data)

	book := models.BookCreate(data.Title, data.PublishYear, data.ISBN)
	if book == nil || book.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	c.Redirect(http.StatusFound, "/books")

}

func BookEdit(c *gin.Context) {
	var data formData
	c.Bind(&data)
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	book := models.BookUpdate(id, data.Title, data.PublishYear, data.ISBN)
	if book == nil || book.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	c.Redirect(http.StatusFound, "/books")

}

func BookRemove(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	res := models.BookDelete(id)
	if res != nil {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}
	c.Redirect(http.StatusFound, "/books")
}
