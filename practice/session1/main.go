/*
// Following the session from the below youtube video : Golang / Go Gin Framework Crash Course 01
//https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w
//FIrst we have to install go gin

//folder path  Microservices/my-project/practice/session1/main.go
*/

package main

import (
	"microservices/my-project/practice/session1/controller"
	"microservices/my-project/practice/session1/middlewares"
	"microservices/my-project/practice/session1/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()

	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	//setupLogOutput()

	// defining the default server
	//	server := gin.Default()

	//if you want to use default middleware you can mentiond, we are doing the same with the default function
	//server := gin.New()
	//server.Use(gin.Recovery())
	//server.use(gin.Logger())
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())

	//creating a get handler
	//an anonayous function to handel it
	//written JSON with status code 200 and  "message" : "OK"
	/*server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})

	})
	*/

	server.GET("/video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	//start the server in release mode
	gin.SetMode(gin.ReleaseMode)
	server.Run(":8080")
}
