package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func DeleteWebsite(c *gin.Context) {
	var website models.ZapWebSite
	err := global.DB.Where("id = ?", c.PostForm("id")).Where("uid = ?", c.GetString("JWT_ID")).First(&website).Error

	if err != nil {
		c.JSON(200, commons.Error(1, "网站不存在", nil))
		return
	}

	//删除目录
	mgr := NewWebSiteMgr(c.GetString("JWT_USERNAME"), website.Domain)
	if err := mgr.RemoveWebsite(); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}

	global.DB.Delete(&website)

	c.JSON(200, commons.SuccessMsg("网站删除成功"))
}
