package websites

import "github.com/gin-gonic/gin"

func CheckDomainIsExists(c *gin.Context) {
	domain := c.PostForm("domain")
	if domain == "" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "域名不能为空",
		})
		return
	}

	if err := CheckDomain(domain); err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}
