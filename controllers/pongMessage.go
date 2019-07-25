package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PongMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
