package task

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/apps"
)

func SystemdAction(c *gin.Context) {
	// action (reload start stop restart)
	action := c.PostForm("action")
	//service name
	name := c.PostForm("name")
	if action == "" || name == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "action or name is empty",
		})
		return
	}
	app := apps.NewApp(name)
	var err error
	switch action {
	case "reload":
		if err = app.Reload(); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		}
	case "start":
		if err = app.Start(); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		}
	case "stop":
		if err = app.Stop(); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
		}
	case "restart":
		if err = app.Restart(); err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  err.Error(),
			})
			return
		}
	default:
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "action is not support",
		})
	}

	if err == nil {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
		})
	}

}
