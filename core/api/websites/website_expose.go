package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/account"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
	"github.com/zapj/zap/core/zapi"
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

func CreateWebsite(c *gin.Context) {
	service := NewWebSiteService(c.GetString("JWT_USERNAME"), zaputil.MustConvertStringToUint(c.GetString("JWT_ID")))
	var websiteReq webSiteRequest

	if err := c.ShouldBind(&websiteReq); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	if err := service.CreateWebsite(websiteReq); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	resp := zapi.Client.Systemctl("reload", global.SERVER_CONF.WebServer)
	if resp.HasError() || resp.Code != 0 {
		c.JSON(200, commons.Error(1, resp.ErrorMessage(), nil))
	}
	c.JSON(200, commons.SuccessMsg("网站添加成功"))
}

func UpdateWebsite(c *gin.Context) {
	service := NewWebSiteService(c.GetString("JWT_USERNAME"), zaputil.MustConvertStringToUint(c.GetString("JWT_ID")))

	var websiteReq webSiteRequest

	if err := c.ShouldBind(&websiteReq); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	if err := service.UpdateWebsite(websiteReq); err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	resp := zapi.Client.Systemctl("reload", global.SERVER_CONF.WebServer)
	if resp.HasError() || resp.Code != 0 {
		c.JSON(200, commons.Error(1, resp.ErrorMessage(), nil))
		return
	}

	c.JSON(200, commons.SuccessMsg("网站修改成功"))
}

func SetWebsiteStatus(c *gin.Context) {
	id := zaputil.MustConvertStringToInt(c.PostForm("id"))
	status := c.PostForm("status")
	if id == 0 {
		c.JSON(200, commons.Error(1, "网站id 不能为空", nil))
		return
	}

	service := NewWebSiteService(c.GetString("JWT_USERNAME"), zaputil.MustConvertStringToUint(c.GetString("JWT_ID")))
	if status == "suspend" {
		if err := service.StopWebsite(id); err != nil {
			c.JSON(200, commons.Error(1, err.Error(), nil))
			return
		}
	} else if status == "running" {
		if err := service.StartWebsite(id); err != nil {
			c.JSON(200, commons.Error(1, err.Error(), nil))
			return
		}
	}
	resp := zapi.Client.Systemctl("reload", global.SERVER_CONF.WebServer)
	if resp.HasError() || resp.Code != 0 {
		c.JSON(200, commons.Error(1, resp.ErrorMessage(), nil))
		return
	}

	c.JSON(200, commons.SuccessMsg("操作成功"))
}
