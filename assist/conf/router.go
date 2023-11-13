package conf

import "github.com/gin-gonic/gin"

func RouterInit(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "test")
	})
}
