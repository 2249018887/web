package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用默认路由引擎
	r := gin.Default()
	//路由，请求到/ping时会触发后边的回调函数，传入gin.Context参数
	r.GET("/ping", func(c *gin.Context) {
		//返回一个json
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		//返回一个字符
		c.String(200, "hello:%v", "word天")
	})
	//POST请求 上传文件
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 监听并在 0.0.0.0:8081 上启动服务,默认8080端口
	r.Run(":8001")
}
