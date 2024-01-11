package main

import (
	"fmt"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		fmt.Printf("gid %d uid %d", syscall.Getegid(), syscall.Geteuid())
	})
	go func() {
		syscall.Setegid(65534)
		syscall.Seteuid(65534)
		router.Run(":8080") // data services
	}()

	routerAdmin := gin.Default()
	routerAdmin.GET("/", func(c *gin.Context) {
		fmt.Printf("gid %d uid %d", syscall.Getegid(), syscall.Geteuid())
	})
	routerAdmin.Run(":8090") // admin and monitor services
}
