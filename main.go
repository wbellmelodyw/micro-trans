package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	g := gin.Default()
	g.GET("/a", func(c *gin.Context) {
		c.JSON(200, "hi")
	})
	service := web.NewService(
		web.Name("go.micro.hello"),
		web.Handler(g),
		web.Address(":8080"),
	)

	service.Init(web.Name("go.micro.hello"))
	service.Run()
}
