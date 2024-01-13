package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/auth/jwtauth"
)

type refreshTokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

func LoginAuthHandler(c *gin.Context) {

	access_token, err := jwtauth.GenerateAccessToken()
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "生成Token 失败"})
	}
	// c.SetCookie("access_token", access_token, 3600, "/", "", false, true)
	c.JSON(200, gin.H{"code": 0, "msg": "Authorization successful", "access_token": access_token})
}

func LogoutAuthHandler(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"code": 0, "msg": "Logout successful"})
}

func RefreshTokenHandler(c *gin.Context) {
	refreshBody := refreshTokenReqBody{}
	if err := c.BindJSON(&refreshBody); err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "refresh_token 刷新失败"})
	}
	//claims
	_, err := jwtauth.CheckJwtToken(refreshBody.RefreshToken)
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "Token 无效"})
	}

	h, err := jwtauth.GenerateAccessToken()
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "Token 无效"})
	}
	c.JSON(200, h)
}
