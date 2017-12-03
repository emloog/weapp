package main

import (

	"github.com/gin-gonic/gin"
	"./controllers"
	//"os"
	//"strings"
)

func main() {
	router := gin.Default()

	router.GET("/user", controllers.UserIndex)


	router.Run(":8080")
}
