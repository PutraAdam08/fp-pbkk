package controllers

import (
	"fmt"
	"fp_pbkk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookIndex(c *gin.Context) {
	books := models.BooksAll(c)
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
