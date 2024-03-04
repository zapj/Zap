package websites

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/utils/zaputil"
)

type webSiteRequest struct {
	Domain       string `json:"domain"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	DomainNames  string `json:"domain_names"`
	WwwRoot      string `json:"www_root"` //只包含目录名 Userhome + wwwroot
	RunDirectory string `json:"run_directory"`
	AccessLog    string `json:"access_log"`
	ErrorLog     string `json:"error_log"`
	WebsiteId    int    `json:"website_id"` // models website id
	applicationRequest
}

type applicationRequest struct {
	Application int    `json:"application"`
	ExposePort  string `json:"expose_port"`
	ExposeProto string `json:"expose_proto"`
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

	c.JSON(200, commons.SuccessMsg("网站修改成功"))
}
