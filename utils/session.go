package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//判斷context裡有沒有session
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	sessionValue := session.Get("userID")
	if sessionValue == nil {
		return false
	} else {
		return true
	}
}

//登入後儲存userID到session裡
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userID", id)
	session.Save()
}

//登出時需要把session清除，包括redis裡的跟web的cookie
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("dummy", "content")
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.SetCookie("SAMPLE", "", -1, "/", "localhost", false, true)
}
