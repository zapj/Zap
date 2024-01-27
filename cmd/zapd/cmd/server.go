package cmd

import (
	"crypto/rand"
	"crypto/tls"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sevlyar/go-daemon"
	"github.com/soheilhy/cmux"
	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/api"
	"github.com/zapj/zap/core/conf"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/web"
)

func init() {

	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().BoolP("daemon", "d", false, "daemon mode")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "zap server",
	Run: func(cmd *cobra.Command, args []string) {
		daemonServ, _ := cmd.Flags().GetBool("daemon")
		if daemonServ {
			cntxt := &daemon.Context{
				PidFileName: "zapd.pid",
				PidFilePerm: 0644,
				LogFileName: "zapd.log",
				LogFilePerm: 0640,
				WorkDir:     "./",
				Umask:       022,
				Args:        []string{},
			}

			d, err := cntxt.Reborn()
			if err != nil {
				global.LOG.Fatal("Unable to run: ", err)
			}
			if d != nil {
				return
			}
			defer cntxt.Release()
		} else {
			pid := os.Getpid()
			os.WriteFile("zapd.pid", []byte(strconv.Itoa(pid)), 0644)
		}
		//初始化缓存
		global.CACHE = cache.New(5*time.Minute, 10*time.Minute)
		conf.InitEnv()
		conf.LogInit()
		conf.DbInit()
		conf.ServerStart_INIT()
		conf.InitCrons()
		gin.SetMode(gin.ReleaseMode)
		if len(args) == 1 && args[0] == "debug" {
			gin.SetMode(gin.DebugMode)
		}
		router := gin.Default()
		router.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
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
		api.RegisterRouter(router)
		api.RegisterAPIV1Router(router.Group("/api/v1"))

		// srv := &http.Server{
		// 	// Addr:    ":2828",
		// 	Handler: router,
		// }

		// if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		// 	global.LOG.Errorf("listen: %s\n", err)
		// } else {
		// 	global.LOG.Debugf("Listen Port : %s", ":2828")
		// }
		// debugServ, _ := cmd.Flags().GetBool("debug")
		// if debugServ {
		// 	fmt.Println("debug")
		// }
		// if len(args) >= 1 && args[0] == "" {
		// 	fmt.Println(args[0])
		// }
		l, err := net.Listen("tcp", ":2828")
		if err != nil {
			global.LOG.Panic(err)
		}
		global.LOG.Println("listen and serve on https://0.0.0.0:2828")
		m := cmux.New(l)

		httpl := m.Match(cmux.HTTP1Fast())
		tlsl := m.Match(cmux.Any())

		go serveHTTP(httpl, router)
		go serveHTTPS(tlsl, router)

		if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
			global.LOG.Panic(err)
		}
	},
}

func serveHTTP(l net.Listener, router *gin.Engine) {
	s := &http.Server{
		Handler: router,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func serveHTTPS(l net.Listener, router *gin.Engine) {
	// Load certificates.
	certificate, err := tls.LoadX509KeyPair("data/server.crt", "data/server.key")
	if err != nil {
		log.Panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		Rand:         rand.Reader,
	}

	// Create TLS listener.
	tlsl := tls.NewListener(l, config)

	// Serve HTTP over TLS.
	serveHTTP(tlsl, router)
}
