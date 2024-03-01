package websites

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/base/ngx"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
)

type WebsiteMgr struct {
	UserHomeDir   string // user home
	Username      string
	WwwRoot       string
	WwwUser       string
	WwwGroup      string
	WwwUserId     int
	WwwGroupId    int
	ServerName    string
	SiteDirectory string //网站目录
	WebSiteId     int    // models website id
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

func NewWebSiteMgrWithWebSiteId(username, serverName string, websiteId int) *WebsiteMgr {
	return &WebsiteMgr{
		UserHomeDir: global.SERVER_CONF.GetUserHomeDir(username),
		Username:    strings.ToLower(username),
		WwwRoot:     filepath.Join(global.SERVER_CONF.GetUserHomeDir(username), strings.ToLower(serverName)),
		WwwUser:     global.SERVER_CONF.WwwUser,
		WwwGroup:    global.SERVER_CONF.WwwGroup,
		WwwUserId:   global.SERVER_CONF.WwwUserId,
		WwwGroupId:  global.SERVER_CONF.WwwGroupId,
		ServerName:  strings.ToLower(serverName),
		WebSiteId:   websiteId,
	}
}

func (w *WebsiteMgr) CreateWebsite(req webSiteRequest) error {
	if w.ServerName != req.WwwRoot {
		w.WwwRoot = filepath.Join(global.SERVER_CONF.GetUserHomeDir(w.Username), strings.ToLower(req.WwwRoot))
	}
	if err := w.preCreateDirs(); err != nil {
		return err
	}

	// create index.html
	createWebsiteDefaultPages(w.WwwRoot)

	//write nginx conf
	userDataPath := w.GetUserNginxConfDir()
	base.TryMkdir(userDataPath)

	ngxServConf := ngx.NewNgxConfServer(w.ServerName)
	ngxServConf.RunDirectory = req.RunDirectory
	ngxServConf.Root = w.WwwRoot

	if conf, err := ngxServConf.GenerateToString(); err == nil {
		writeNgxConfFile(filepath.Join(userDataPath, w.ServerName+".conf"), conf)
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
		slog.Info("remove website", "wwwRoot", w.WwwRoot)
		if err := os.RemoveAll(w.WwwRoot); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("website not exist")
	}

	//删除nginx配置文件

	if fileutils.IsFile("/usr/local/apps/nginx/conf/sites-enabled/" + w.ServerName + ".conf") {
		os.Remove("/usr/local/apps/nginx/conf/sites-enabled/" + w.ServerName + ".conf")
	}

	srvConfFile := w.GetNgxConfig(w.ServerName + ".conf")
	slog.Info("srvConfFile", "file", srvConfFile)
	if fileutils.IsFile(srvConfFile) {
		if err := os.Remove(srvConfFile); err != nil {
			return err
		}
	}

	return nil
}

func (w *WebsiteMgr) GetUserConfig(conf ...string) string {
	if len(conf) > 0 {
		return pathutil.GetPath("data/users", w.Username, path.Join(conf...))
	}
	return pathutil.GetPath("data/users", w.Username)
}

func (w *WebsiteMgr) GetNgxConfig(filename string) string {
	return w.GetUserConfig("nginx_conf.d", filename)
}

// 更新网站
func (w *WebsiteMgr) UpdateWebsite(req webSiteRequest) error {
	website, err := GetWebsiteById(w.WebSiteId)
	if err != nil {
		return err
	}
	// if website.Domain != req.Domain {

	// }

	website.Domain = req.Domain
	website.DomainNames = req.DomainNames
	website.Description = req.Description
	website.Title = req.Title
	website.ApplicationId = req.Application
	if w.ServerName != req.WwwRoot {
		w.WwwRoot = filepath.Join(global.SERVER_CONF.GetUserHomeDir(w.Username), strings.ToLower(req.WwwRoot))
		if !fileutils.IsDir(w.WwwRoot) {
			os.Rename(website.WwwRoot, w.WwwRoot)
		}
		website.WwwRoot = w.WwwRoot
	}
	if err := w.preCreateDirs(); err != nil {
		return err
	}

	//write nginx conf
	userNgxPath := w.GetUserNginxConfDir()
	base.TryMkdir(userNgxPath)
	ngxServConf := ngx.NewNgxConfServer(w.ServerName)
	ngxServConf.RunDirectory = req.RunDirectory
	ngxServConf.Root = w.WwwRoot

	if ngServer, err := ngxServConf.GenerateToString(); err == nil {
		writeNgxConfFile(filepath.Join(userNgxPath, w.ServerName+".conf"), ngServer)
	} else {
		return err
	}
	return nil
}

func (w *WebsiteMgr) GetUserNginxConfDir() string {
	return pathutil.GetPath("data/users", w.Username, "nginx_conf.d")
}

func (w *WebsiteMgr) GetWebsiteConfFileName() string {

	return pathutil.GetPath("data/users", w.Username, "nginx_conf.d", w.ServerName+".conf")
}

func writeNgxConfFile(userNgxPath string, conf string) error {
	return os.WriteFile(userNgxPath, []byte(conf), 0644)
}

func getDefaultIndexHtml() []byte {
	indexData, _ := os.ReadFile(pathutil.GetPath("scripts/zap/html/index.html"))
	return indexData
}

func createWebsiteDefaultPages(www_root string) error {
	indexData := getDefaultIndexHtml()
	return os.WriteFile(filepath.Join(www_root, "index.html"), indexData, 0644)
}
