package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionValue := session.Get("userID")
	if sessionValue == nil {
		return false
	} else {
		return true
	}
}

func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userID", id)
	session.Save()
}

func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("dummy", "content")
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.SetCookie("SAMPLE", "", -1, "/", "localhost", false, true)
}

func GetSessionUserID(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userID")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}
