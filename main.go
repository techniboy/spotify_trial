package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techniboy/spotify_trial/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.HomeMessage)
	r.GET("/ping", controllers.PongMessage)
	r.Run()
}