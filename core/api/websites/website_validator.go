package websites

import "github.com/gin-gonic/gin"

func CheckDomainIsExists(c *gin.Context) {
	domain := c.PostForm("domain")
	websiteId := c.PostForm("website_id")
	if domain == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "域名不能为空",
		})
		return
	}

	if !CheckDomain(domain, websiteId, c.GetString("JWT_ID")) {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "域名已存在",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}
