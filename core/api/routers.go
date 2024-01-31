package api

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zapj/zap/core/api/appstore"
	"github.com/zapj/zap/core/api/dashboard"
	"github.com/zapj/zap/core/api/filemanager"
	"github.com/zapj/zap/core/api/server"
	"github.com/zapj/zap/core/api/users"
	"github.com/zapj/zap/core/auth/jwtauth"
	"github.com/zapj/zap/core/webterm"
)

// 无需授权的路由
func RegisterRouter(router *gin.Engine) {
	router.POST("/api/login", LoginAuthHandler)
	router.POST("/api/logout", LogoutAuthHandler)

	// router.POST("/api/refresh_token", LoginAuthHandler)

}

// api v1
func RegisterAPIV1Router(c *gin.RouterGroup) {
	c.Use(jwtTokenCheck)
	//websocket
	c.GET("/local/ws", webterm.HandlerLocalWS)
	c.GET("/statistics/dashboard", dashboard.DashBoardStats)
	c.GET("/sync/user/settings", users.SyncUserSettings)

	c.POST("/filemanager/list", filemanager.FileManager_List)
	c.POST("/filemanager/fetch", filemanager.FileManager_Fetch)
	c.POST("/filemanager/putfile", filemanager.FileManager_PutFile)

	c.GET("/server/info", server.ServerInfo)
	c.GET("/server/processlist", server.ServerProcessList)
	c.GET("/server/netinterface_list", server.ServerNetInterfaces)
	c.GET("/server/top_info", server.ServerTopInfo)

	c.GET("/app/appstore", appstore.ListApp)

	c.GET("/upgrade/check", server.UpgradeCheck)

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

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return jwtauth.SigningKey, nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
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

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":  "bad jwt token",
			"code": 1,
		})
		slog.Info("bad jwt token")
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":  "unable to parse claims",
			"code": 1,
		})
		slog.Info("unable to parse claims")
		return
	}
	subject, err := claims.GetSubject()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":  "jwt invalid",
			"code": 1,
		})
		slog.Info("jwt invalid")
		return
	}
	c.Set("JWT_SUBJECT", subject)
	c.Set("JWT_TOKEN", jwtToken)
	c.Next()
}
