package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	Get(r)
	Post(r)

	fmt.Println("Starting server on :8080")
	r.Run(":8080") // listen and serve on
}
