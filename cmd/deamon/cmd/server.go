package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os/exec"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/api"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/webterm"
	"github.com/zapj/zap/web"
)

func init() {

	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().BoolP("debug", "d", false, "启动debug")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动ZAP服务",
	Run: func(cmd *cobra.Command, args []string) {
		router := gin.Default()
		assetsFS, _ := fs.Sub(web.ASSETS_FS, "static/assets")

		router.StaticFS("/assets", http.FS(assetsFS))
		router.GET("/", func(c *gin.Context) {
			c.Data(200, "text/html", web.INDEX_BYTES)
		})
		router.NoRoute(func(c *gin.Context) {
			accept := c.Request.Header.Get("Accept")
			flag := strings.Contains(accept, "text/html")
			if flag {
				c.Data(200, "text/html", web.INDEX_BYTES)
			}
		})
		router.GET("/favicon.ico", func(c *gin.Context) {
			c.Data(200, "image/x-icon", web.FAVICON_BYTES)
		})

		router.GET("/start", func(c *gin.Context) {
			go func() {
				cmd := exec.Command("./zap_cli", "version")
				cmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
				cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 65534, Gid: 65534, NoSetGroups: true}
				if err := cmd.Run(); err != nil {
					global.LOG.Error(err)
				}

			}()
		})

		v1 := router.Group("/api/v1")
		{
			api.RegisterRouter(v1)

		}
		router.GET("/ws", webterm.HandlerLocalWS)

		srv := &http.Server{
			Addr:    ":2828",
			Handler: router,
		}

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Errorf("listen: %s\n", err)
		} else {
			global.LOG.Debugf("Listen Port : %s", ":2828")
		}
		debugServ, _ := cmd.Flags().GetBool("debug")
		if debugServ {
			fmt.Println("debug")
		}
		// if len(args) >= 1 && args[0] == "" {
		// 	fmt.Println(args[0])
		// }
	},
}
