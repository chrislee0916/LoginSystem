package controllers

import (
	"net/http"

	"chris_project/models"
	"chris_project/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err = models.RegisterCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "register successfully!",
	})

}

type LoginInput struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	id, err := models.LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	utils.SaveAuthSession(c, id)

	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"message":   "login successfully",
		"userID":    id,
		"sessionID": sessions.Default(c).ID(),
	})
}

func Logout(c *gin.Context) {
	if !utils.HasSession(c) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": "you haven't login",
		})
		return
	}
	utils.ClearAuthSession(c)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "logout successfully",
	})
}

func CurrentUser(c *gin.Context) {

	userID := c.MustGet("userID").(uint)
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":         200,
		"current_user": user,
	})
}
