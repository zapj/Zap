package task

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/utils/cmdutil"
)

func SystemdAction(c *gin.Context) {
	// action (reload start stop restart)
	action := c.PostForm("action")
	//service name
	name := c.PostForm("name")
	if err := SystemdActionHandler(action, name); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
		})
	}
}

func SystemdActionHandler(action, name string) error {
	if action == "" || name == "" {
		return nil
	}
	result, err := cmdutil.ExecCmd("systemctl", action, name)
	slog.Info("systemctl", "result", result, "action", action, "name", name, "err", err)
	if err != nil {
		return err
	}
	return nil
}
