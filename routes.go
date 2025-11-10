package main

import "github.com/gin-gonic/gin"

func Get(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/start", start)
	r.GET("/last", Get_Last_Message)
}

func Post(r *gin.Engine) {
	r.POST("/add", AddToChat)
}
