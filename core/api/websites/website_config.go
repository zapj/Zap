package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/account"
	"github.com/zapj/zap/core/api/commons"
)

func CreateWebSiteConfig(c *gin.Context) {
	accountService := account.NewAccountService(c.GetString("JWT_USERNAME"))
	websiteConfig := accountService.GetWebsiteConfig()
	c.JSON(200, commons.Success(websiteConfig))
}
