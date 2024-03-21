package cmd

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/sevlyar/go-daemon"
	"github.com/soheilhy/cmux"
	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/api"
	"github.com/zapj/zap/core/conf"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/zaputil"
	"github.com/zapj/zap/core/utils/zlog"
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
				zlog.Fatal("Unable to run: ", err)
			}
			if d != nil {
				return
			}
			defer cntxt.Release()
		} else {
			pid := os.Getpid()
			os.WriteFile(pathutil.GetPath("zapd.pid"), []byte(strconv.Itoa(pid)), 0644)
		}

		//初始化缓存
		global.CACHE = cache.New(5*time.Minute, 10*time.Minute)
		conf.InitEnv()

		// {Uid: 104, Gid: 116}

		if global.ZAP_MODE == "PRO" {
			u, err := user.Lookup("zapadm")
			if err != nil {
				slog.Info("zapadm user not found ")
			}
			if syscall.Getegid() == 0 {
				syscall.Setgid(zaputil.MustConvertStringToInt(u.Gid))
			}
			if syscall.Geteuid() == 0 {
				syscall.Setuid(zaputil.MustConvertStringToInt(u.Uid))
			}
		}
		conf.LogInit("zaps")

		conf.DbInit()
		conf.ServerStart_INIT()
		conf.InitCrons()
		gin.SetMode(gin.ReleaseMode)
		if global.ZAP_MODE == "DEV" {
			gin.SetMode(gin.DebugMode)
		}
		sessionStore := gormsessions.NewStore(global.DB, true, []byte(global.SERVER_CONF.SigningKey))
		router := gin.Default()
		router.Use(sessions.Sessions("ZAPSID", sessionStore))
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
			// go func() {
			// 	cmd := exec.Command("./zap_cli", "version")
			// 	cmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
			// 	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 65534, Gid: 65534, NoSetGroups: true}
			// 	if err := cmd.Run(); err != nil {
			// 		slog.Error("", err)
			// 	}

			// }()
		})
		api.RegisterRouter(router)
		api.RegisterAPIV1Router(router.Group("/api/v1"))

		l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", global.SERVER_CONF.Host, global.SERVER_CONF.Port))
		if err != nil {
			zlog.Fatalf("无法监听port:%d , err : %s", global.SERVER_CONF.Port, err.Error())
		}
		zlog.Infof("listen and serve on https://%s:%d", global.SERVER_CONF.Host, global.SERVER_CONF.Port)
		m := cmux.New(l)

		httpl := m.Match(cmux.HTTP1Fast())
		tlsl := m.Match(cmux.Any())

		go serveHTTP(httpl)
		go serveHTTPS(tlsl, router)

		if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
			zlog.Fatal("start server err", "err", err)
		}
	},
}

func serveHTTP(l net.Listener) {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, fmt.Sprintf("https://%s%s", req.Host, req.RequestURI), http.StatusMovedPermanently)
		}),
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func serveHTTPS(l net.Listener, router *gin.Engine) {
	// Load certificates.
	certificate, err := tls.LoadX509KeyPair(global.SERVER_CONF.CertFile, global.SERVER_CONF.KeyFile)
	if err != nil {
		zlog.Fatal("load certificate", "err", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		MinVersion:   tls.VersionTLS10,
		Rand:         rand.Reader,
	}

	// Create TLS listener.
	tlsl := tls.NewListener(l, config)

	s := &http.Server{
		Handler: router,
	}
	if err := s.Serve(tlsl); err != cmux.ErrListenerClosed {
		panic(err)
	}
}
