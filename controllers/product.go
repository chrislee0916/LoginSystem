package controllers

import (
	"chris_project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	CategoryID uint   `json:"category_id" form:"category_id"`
	PName      string `json:"p_name" form:"p_name"`
}

func GetProduct(c *gin.Context) {
	pid := c.Param("pid")

	res, err := models.GetProductByID(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})

}

func CreateProduct(c *gin.Context) {
	input := ProductInput{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(uint)
	err := models.CreateProduct(userID, input.CategoryID, input.PName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "product save successfully",
	})

}

func UpdateProduct(c *gin.Context) {
	input := ProductInput{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	pid := c.Param("pid")
	err := models.UpdateProductByID(pid, input.CategoryID, input.PName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Update successfully",
		"pid":     pid,
	})
}

func DeleteProduct(c *gin.Context) {
	pid := c.Param("pid")
	err := models.DeleteProductByID(pid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "delete successfully",
		"pid":     pid,
	})
}
