package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	// 记录到文件。
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/**/*")
	engine.Static("/assets", "./assets")

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index/index.gohtml", gin.H{
			"message": "hello gin",
		})
	})

	engine.GET("/posts/index", func(context *gin.Context) {
		/*posts/index.gohtml这个名字需要在.gohtml里面用define进行定义，否则gin会找不到这个名字*/
		context.HTML(http.StatusOK, "posts/index.gohtml", gin.H{
			"post": "This is test post",
		})
	})

	engine.Run(":8080")
}
