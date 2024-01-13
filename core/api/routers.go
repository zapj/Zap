package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zapj/zap/core/auth/jwtauth"
	"github.com/zapj/zap/core/webterm"
)

// 无需授权的路由
func RegisterRouter(router *gin.Engine) {
	router.POST("/api/login", LoginAuthHandler)
	router.POST("/api/logout", LogoutAuthHandler)

	router.POST("/api/refresh_token", LoginAuthHandler)
	router.GET("/api/ws", webterm.HandlerLocalWS)

}

// api v1
func RegisterAPIV1Router(c *gin.RouterGroup) {
	c.Use(jwtTokenCheck)

	c.GET("/meminfo", MemoryUsageInfo)
	c.GET("/loadavg", func(c *gin.Context) {
		st, _ := load.Avg()
		c.String(200, "Loadavg  %s", st.String())
	})
	c.GET("/mem", func(c *gin.Context) {
		v, _ := mem.VirtualMemory()
		c.String(200, v.String())
	})
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
	// jwtToken, err := c.Cookie("access_token")
	// if err != nil {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	// if err != nil {
	// 	jwtToken, _ = c.Cookie("access_token")
	// }
	// }
	if err != nil {
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
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"msg":  "unable to parse claims",
			"code": 1,
		})
		return
	}
	c.Next()
}
