package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	Get(r)
	Post(r)

	fmt.Println("Starting server on :8080")
	r.Run(":8080") // listen and serve on
}
