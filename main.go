package main

import (
  "github.com/gin-gonic/gin"
	"chatting-api/controllers"
)

func main() {

	r := gin.Default()

	r.GET("/order", controllers.Getting)
	r.POST("/order", controllers.Posting)
	r.DELETE("/order", controllers.Deleting)
	r.PATCH("/order", controllers.Patching)
	r.OPTIONS("/order", controllers.Options)

	r.Run(":8080")
}
