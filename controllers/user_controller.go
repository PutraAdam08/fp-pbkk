package controllers

import (
	"fp_pbkk/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type userData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func SignupPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"auth/signup.tpl",
		gin.H{})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK,
		"auth/login.tpl",
		gin.H{})
}

func SignUp(c *gin.Context) {
	var data userData
	c.Bind(&data)

	if !models.CheckUserAvailability(data.Email) {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	user := models.UserCreate(data.Email, data.Password)
	if user == nil || user.ID == 0 {
		c.Render(http.StatusBadRequest, render.Data{})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()
	//c.Redirect(http.StatusFound, "/blogs")
}

func Login(c *gin.Context) {
	var data userData
	c.Bind(&data)

	user := models.UserMatchPassword(data.Email, data.Password)
	if user.ID == 0 {
		c.Render(http.StatusUnauthorized, render.Data{})
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	if user.IsAdmin {
		c.Redirect(http.StatusFound, "/admin")
	} else {
		c.Redirect(http.StatusFound, "/blogs")
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusAccepted)
}
