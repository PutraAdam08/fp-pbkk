package main

import (
	"fp_pbkk/controllers"
	m "fp_pbkk/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/*")

	m.DBconnect()
	m.DBMigrate()
	m.DBSeed()

	r.GET("/", controllers.MainPage)
	r.GET("/books", controllers.BookIndex)
	r.GET("/books/:id", controllers.BookShow)

	r.Run() // Default Port 8080
}
