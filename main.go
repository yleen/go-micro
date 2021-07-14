package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //返回默认路由引擎

	r.GET("/hello", sayHello)
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "post",
		})
	})
	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "put",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "delete",
		})
	})
	r.Run()
	// r.Run("9090") 默认8080 可自定义端口

}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}
