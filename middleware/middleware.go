package middleware

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")
		username := session.Get("username")
		isLoggedIn := false
		if userID != nil {
			log.Println("User ID found in session:", userID)
			isLoggedIn = true
		} else {
			log.Println("User ID not found in session")
		}
		c.Set("isLoggedIn", isLoggedIn)
		c.Set("username", username)
		c.Next()
	}
}
