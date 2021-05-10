package main

import (
	"github.com/gin-gonic/gin"
	"log/ginDemo/config"
	"log/ginDemo/middleware"
	"log/ginDemo/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}