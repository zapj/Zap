package account

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/auth/jwtauth"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
	"gorm.io/gorm"
)

type refreshTokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}
type loginReqBody struct {
	Username string `json:"username"`
	Passowrd string `json:"password"`
	AuthType string `json:"auth_type"`
}

func LoginAuthHandler(c *gin.Context) {
	userForm := loginReqBody{}
	c.BindJSON(&userForm)
	if userForm.Username == "" || userForm.Passowrd == "" {
		c.JSON(200, gin.H{"code": 1, "msg": "登陆失败, 用户名或密码不正确"})
		return
	}
	user := models.ZapUsers{Username: userForm.Username}
	rs := global.DB.First(&user)
	if errors.Is(rs.Error, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{"code": 1, "msg": "用户名或密码不正确"})
		return
	}
	// slog.Error("userForm.Passowrd, user.Password", "userform", userForm.Passowrd, "userPassword", user.Password)
	if !zaputil.CheckPasswordHash(userForm.Passowrd, user.Password) {
		c.JSON(200, gin.H{"code": 1, "msg": "用户名或密码不正确"})
		return
	}
	if userForm.AuthType == "jwt" {
		access_token, err := jwtauth.GenerateAccessToken(fmt.Sprint(user.ID), user.Username)
		if err != nil {
			c.JSON(200, gin.H{"code": 1, "msg": "生成Token 失败", "err": err.Error()})
			return
		}

		// c.SetCookie("access_token", access_token, 3600, "/", "", false, true)
		c.JSON(200, gin.H{"code": 0, "msg": "Authorization successful", "access_token": access_token})
	} else {
		session := sessions.Default(c)
		session.Set("JWT_ID", user.ID)
		session.Set("JWT_USERNAME", user.Username)
		session.Save()
		c.JSON(200, gin.H{"code": 0, "msg": "Authorization successful"})
	}

}

func LogoutAuthHandler(c *gin.Context) {
	// c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"code": 0, "msg": "Logout successful"})
}

func RefreshTokenHandler(c *gin.Context) {
	refreshBody := refreshTokenReqBody{}
	if err := c.BindJSON(&refreshBody); err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "refresh_token 刷新失败"})
	}
	//claims
	claims, err := jwtauth.CheckJwtToken(refreshBody.RefreshToken)
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "Token 无效"})
	}

	h, err := jwtauth.GenerateAccessToken(claims.ID, claims.Username)
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "Token 无效"})
	}
	c.JSON(200, h)
}
