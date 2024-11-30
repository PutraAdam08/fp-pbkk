package middlewares

import (
	"fp_pbkk/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		session := sessions.Default(c)
		sessionID := session.Get("userID")
		if sessionID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userId := sessionID.(uint)
		// Check if the user exists
		user := models.UserFromId(userId)
		if user.ID == 0 || user.IsAdmin {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("userID", user.ID)

		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		session := sessions.Default(c)
		sessionID := session.Get("userID")
		if sessionID == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userId := sessionID.(uint)
		// Check if the user exists
		user := models.UserFromId(userId)
		if user.ID == 0 || !user.IsAdmin {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("userID", user.ID)

		c.Next()
	}
}
