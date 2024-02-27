package websites

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
)

type WebSiteForm struct {
	Domain string `json:"domain"`
}

func CreateWebsite(c *gin.Context) {
	var website models.ZapWebSite
	if err := c.ShouldBind(&website); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	website.Uid = zaputil.MustConvertStringToUint(c.GetString("JWT_ID"))
	website.Status = "active"

	mgr := NewWebSiteMgr(c.GetString("JWT_USERNAME"), website.Domain)
	err := mgr.CreateWebsite()
	slog.Info("创建网站成功", "err", website)
	if err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	website.WwwRoot = mgr.WwwRoot

	if website.Title == "" {
		website.Title = website.Domain
	}

	global.DB.Save(&website)

	c.JSON(200, commons.SuccessMsg("网站添加成功"))
}
