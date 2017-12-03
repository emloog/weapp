package main

import (

	"github.com/gin-gonic/gin"
	"./controllers"
	//"os"
	//"strings"
	//"html/template"
	//"github.com/GeertJohan/go.rice"

)

func main() {
	router := gin.Default()

	router.GET("/user", controllers.UserIndex)

	router.Run(":8080")
}
