package commons

import "github.com/gin-gonic/gin"

func JSONError(code int, msg string) gin.H {
	return gin.H{"code": code, "msg": msg}
}
