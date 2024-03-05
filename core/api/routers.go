package api

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zapj/zap/core/api/account"
	"github.com/zapj/zap/core/api/appstore"
	"github.com/zapj/zap/core/api/dashboard"
	"github.com/zapj/zap/core/api/filemanager"
	"github.com/zapj/zap/core/api/server"
	"github.com/zapj/zap/core/api/websites"
	"github.com/zapj/zap/core/auth/jwtauth"
	"github.com/zapj/zap/core/terminal"
)

// 无需授权的路由
func RegisterRouter(router *gin.Engine) {
	router.POST("/api/login", account.LoginAuthHandler)
	router.POST("/api/logout", account.LogoutAuthHandler)

	// router.POST("/api/refresh_token", LoginAuthHandler)

}

// api v1
func RegisterAPIV1Router(c *gin.RouterGroup) {
	c.Use(jwtTokenCheck)
	//websocket
	c.GET("/local/ws", terminal.HandlerLocalWS)
	c.GET("/statistics/dashboard", dashboard.DashBoardStats)
	c.GET("/sync/user/settings", account.SyncUserSettings)

	// websites
	c.GET("/website/list", websites.ListWebsite)
	c.POST("/website/create", websites.CreateWebsite)
	c.POST("/website/update", websites.UpdateWebsite)
	c.POST("/website/delete", websites.DeleteWebsite)
	c.GET("/website/config", websites.CreateWebSiteConfig)
	c.GET("/website/settings", websites.WebSiteSettings)
	c.POST("/website/check_domain", websites.CheckDomainIsExists)

	// filemanager
	c.POST("/filemanager/list", filemanager.FileManager_List)
	c.POST("/filemanager/fetch", filemanager.FileManager_Fetch)
	c.POST("/filemanager/putfile", filemanager.FileManager_PutFile)
	c.POST("/filemanager/delete", filemanager.FileManager_RemoveFile)

	// server
	c.GET("/server/info", server.ServerInfo)
	c.GET("/server/processlist", server.ServerProcessList)
	c.GET("/server/netinterface_list", server.ServerNetInterfaces)
	c.GET("/server/top_info", server.ServerTopInfo)

	// appstore
	c.GET("/app/appstore", appstore.ListApp)
	c.GET("/app/appstore/already_install", appstore.ListInstalledApp)
	c.POST("/app/appstore/uninstall", appstore.UninstallApp)
	c.POST("/app/appstore/install", appstore.AppInstall)

	//task queue
	c.GET("/task/appinstall/tasklist", appstore.TaskList)
	c.POST("/task/appinstall/removetask", appstore.RemoveTask)
	c.POST("/task/appinstall/gentask", appstore.GenTask)
	c.POST("/task/appinstall/canceltask", appstore.CancelTask)

	// upgrade
	c.GET("/upgrade/check", server.UpgradeCheck)

	// WSS
	c.GET("/wss/logviewer", terminal.LogViewerHandler)

}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func jwtTokenCheck(c *gin.Context) {
	// if c.Request.RequestURI == "/api/v1/local/ws" {
	// 	c.Next()
	// 	return
	// }
	// jwtToken, err := c.Cookie("access_token")
	// if err != nil {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	// if err != nil {
	// 	jwtToken, _ = c.Cookie("access_token")
	// }
	// }
	if token := c.GetHeader("Sec-Websocket-Protocol"); token != "" {
		slog.Info("sec websocket protocol", "Sec-Websocket-Protocol", c.GetHeader("Sec-Websocket-Protocol"))
		jwtToken = token
		err = nil
	}
	if err != nil {
		slog.Info("StatusBadRequest", "err", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  err.Error(),
			"code": 1,
		})
		return
	}

	claims, err := jwtauth.CheckJwtToken(jwtToken)
	if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "Token 已过期",
			"code": 1,
		})
		return
	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "认证失败" + err.Error(),
			"code": 1,
		})
		return
	}
	ID := claims.ID
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "jwt invalid", "code": 1})
		return
	}
	c.Set("JWT_ID", claims.ID)
	c.Set("JWT_USERNAME", claims.Username)
	c.Set("JWT_TOKEN", jwtToken)
	c.Next()
}
