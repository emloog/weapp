package controllers

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
	"html/template"
	"github.com/GeertJohan/go.rice"
)

func UserIndex(c *gin.Context) {

	// 这里写相对于的执行文件的地址
	box, err := rice.FindBox("views")
	if err != nil {
		println(err.Error())
		return
	}
	// 从目录 Box 读取文件
	str, err := box.String("users/index.html")
	if err != nil {
		println(err.Error())
		return
	}
	t, err := template.New("tpl").Parse(str)
	fmt.Println(t, err)

}
