package controllers

import (
	"fmt"
	"fp_pbkk/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type categoryData struct {
	Name string `form:"name"`
}

func CategoryShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	category := models.CategoryFind(id)
	c.HTML(
		http.StatusOK,
		"category/show.tpl",
		gin.H{
			"category": category,
		},
	)
}

func CategoryAdd(c *gin.Context) {
	var data categoryData
	c.Bind(&data)

	category := models.CategoryCreate(data.Name)
	if category == nil || category.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	c.Redirect(http.StatusFound, "/books")

}

func categoryEdit(c *gin.Context) {
	var data categoryData
	c.Bind(&data)
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	category := models.CategoryUpdate(id, data.Name)
	if category == nil || category.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	c.Redirect(http.StatusFound, "/category")

}

func categoryRemove(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	res := models.CategoryDelete(id)
	if res != nil {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}
	c.Redirect(http.StatusFound, "/category")
}
