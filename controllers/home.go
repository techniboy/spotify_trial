package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Gin Works!",
	})
}
