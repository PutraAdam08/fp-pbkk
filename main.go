package main

import (
	"fp_pbkk/controllers"
	"fp_pbkk/middlewares"
	m "fp_pbkk/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(gin.Logger())

	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("users", store))
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/*")

	m.DBconnect()
	//m.DBMigrate()
	//m.DBSeed()

	books := r.Group("/books", middlewares.UserAuthMiddleware())
	{
		books.GET("/", controllers.BookIndex)
		books.GET("/:id", controllers.BookShow)

	} /**/

	admins := r.Group("/admin/", middlewares.AdminAuthMiddleware())
	{
		admins.GET("dashboard/", controllers.DashboardPage)
		admins.GET("books/", controllers.BookIndex)
		admins.GET("books/:id", controllers.BookShow)
		admins.POST("books/add", controllers.BookAdd)
		admins.POST("books/edit", controllers.BookEdit)
		admins.GET("remove/:id", controllers.BookRemove)

	} /**/

	//for frontend development
	/*
		r.GET("books/", controllers.BookIndex)
		r.GET("books/:id", controllers.BookShow)
	/**/

	r.GET("/", controllers.MainPage)
	r.GET("/signup", controllers.SignupPage)
	r.GET("/login", controllers.LoginPage)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.DELETE("/logout", controllers.Logout)

	r.Run() // Default Port 8080
}
