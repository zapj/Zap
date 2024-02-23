package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func CreateWebsite(c *gin.Context) {
	var website models.ZapWebSite
	if err := c.ShouldBind(&website); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	website.Uid = uint(c.GetInt("JWT_ID"))
	website.Status = "active"
	global.DB.Save(&website)
}
