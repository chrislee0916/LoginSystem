package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var KEY = os.Getenv("KEY")

func EnableRedisSession() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userID")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("userID", sessionValue.(uint))
		c.Next()

	}
}
