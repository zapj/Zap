package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/account"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func WebSiteSettings(c *gin.Context) {
	websiteId := c.Query("id")
	if websiteId == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "id 不能为空",
		})
		return
	}
	website := models.ZapWebSite{}
	res := global.DB.Find(&website, "id = ?", websiteId)
	if res.Error != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  res.Error.Error(),
		})
		return
	}
	accountService := account.NewAccountService(c.GetString("JWT_USERNAME"))
	websiteConfig := accountService.GetWebsiteConfig()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{
			"website": website,
			"config":  websiteConfig,
		},
	})
}
