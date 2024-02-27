package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
)

func DeleteWebsite(c *gin.Context) {
	var website models.ZapWebSite
	website.ID = zaputil.MustConvertStringToUint(c.PostForm("id"))
	website.Uid = zaputil.MustConvertStringToUint(c.GetString("JWT_ID"))
	// website.Status = "active"
	global.DB.Delete(&website)
	//删除目录
	mgr := NewWebSiteMgr(c.GetString("JWT_USERNAME"), website.Domain)
	mgr.RemoveWebsite()

	c.JSON(200, commons.SuccessMsg("网站删除成功"))
}
