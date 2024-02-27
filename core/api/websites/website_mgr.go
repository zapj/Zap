package websites

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/base/ngx"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
)

type WebsiteMgr struct {
	UserHomeDir string // user home
	Username    string
	WwwRoot     string
	WwwUser     string
	WwwGroup    string
	WwwUserId   int
	WwwGroupId  int
	ServerName  string
}

func NewWebSiteMgr(username, serverName string) *WebsiteMgr {
	return &WebsiteMgr{
		UserHomeDir: global.SERVER_CONF.GetUserHomeDir(username),
		Username:    strings.ToLower(username),
		WwwRoot:     filepath.Join(global.SERVER_CONF.GetUserHomeDir(username), strings.ToLower(serverName)),
		WwwUser:     global.SERVER_CONF.WwwUser,
		WwwGroup:    global.SERVER_CONF.WwwGroup,
		WwwUserId:   global.SERVER_CONF.WwwUserId,
		WwwGroupId:  global.SERVER_CONF.WwwGroupId,
		ServerName:  strings.ToLower(serverName),
	}
}

func (w *WebsiteMgr) CreateWebsite() error {
	if err := w.preCreateDirs(); err != nil {
		return err
	}
	indexFile, err := os.Create(filepath.Join(w.WwwRoot, "index.html"))
	if err != nil {
		return err
	}
	defer indexFile.Close()
	indexData, err := os.ReadFile(pathutil.GetPath("scripts", "zap", "html", "index.html"))
	if err != nil {
		return err
	}
	indexFile.Write(indexData)

	//write nginx conf
	userDataPath := pathutil.GetPath("data/users", w.Username, "nginx_conf.d")
	slog.Info("userDataPath: ", "path", userDataPath)
	if !fileutils.IsDir(userDataPath) {
		os.MkdirAll(userDataPath, 0755)
	}
	slog.Info("ngConfFile", "file", filepath.Join(userDataPath, w.ServerName+".conf"))
	ngConfFile, err := os.Create(filepath.Join(userDataPath, w.ServerName+".conf"))
	if err != nil {
		return err
	}
	defer ngConfFile.Close()
	ngxServConf := ngx.NewNgxConfServer(w.ServerName)
	if conf, err := ngxServConf.GenerateToServer(); err == nil {
		ngConfFile.WriteString(conf)
	} else {
		return err
	}
	return nil
}

func (w *WebsiteMgr) preCreateDirs() error {
	if _, err := os.Stat(w.UserHomeDir); os.IsNotExist(err) {
		err = os.Mkdir(w.UserHomeDir, 0755)
		if err != nil {
			return err
		}
		err = os.Chown(w.UserHomeDir, w.WwwUserId, w.WwwGroupId)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(w.WwwRoot); os.IsNotExist(err) {
		err = os.Mkdir(w.WwwRoot, 0755)
		if err != nil {
			return err
		}
		err = os.Chown(w.WwwRoot, w.WwwUserId, w.WwwGroupId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WebsiteMgr) RemoveWebsite() error {
	if w.ServerName != "" && fileutils.IsDir(w.WwwRoot) && strings.HasPrefix(w.WwwRoot, w.UserHomeDir) {
		if err := os.RemoveAll(w.WwwRoot); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("website not exist")
	}

	//删除nginx配置文件
	srvConfFile := w.GetNgxConfig(w.ServerName + ".conf")
	if fileutils.IsFile(srvConfFile) {
		if err := os.Remove(srvConfFile); err != nil {
			return err
		}
	}

	return nil
}

func (w *WebsiteMgr) GetUserConfig(conf ...string) string {
	if len(conf) > 0 {
		return pathutil.GetPath("data/users", w.Username, conf[0])
	}
	return pathutil.GetPath("data/users", w.Username)
}

func (w *WebsiteMgr) GetNgxConfig(filename string) string {
	return w.GetUserConfig("nginx_conf.d", filename)
}
