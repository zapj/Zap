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
	"github.com/zapj/zap/core/models"
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
	WebServerType string // nginx / openresty
}

func NewWebSiteMgr(username, serverName string) *WebsiteMgr {
	return &WebsiteMgr{
		UserHomeDir:   global.SERVER_CONF.GetUserHomeDir(username),
		Username:      strings.ToLower(username),
		WwwRoot:       filepath.Join(global.SERVER_CONF.GetUserHomeDir(username), strings.ToLower(serverName)),
		WwwUser:       global.SERVER_CONF.WwwUser,
		WwwGroup:      global.SERVER_CONF.WwwGroup,
		WwwUserId:     global.SERVER_CONF.WwwUserId,
		WwwGroupId:    global.SERVER_CONF.WwwGroupId,
		ServerName:    strings.ToLower(serverName),
		WebServerType: "nginx",
	}
}

func NewWebSiteMgrWithWebSiteId(username, serverName string, websiteId int) *WebsiteMgr {
	return &WebsiteMgr{
		UserHomeDir:   global.SERVER_CONF.GetUserHomeDir(username),
		Username:      strings.ToLower(username),
		WwwRoot:       filepath.Join(global.SERVER_CONF.GetUserHomeDir(username), strings.ToLower(serverName)),
		WwwUser:       global.SERVER_CONF.WwwUser,
		WwwGroup:      global.SERVER_CONF.WwwGroup,
		WwwUserId:     global.SERVER_CONF.WwwUserId,
		WwwGroupId:    global.SERVER_CONF.WwwGroupId,
		ServerName:    strings.ToLower(serverName),
		WebSiteId:     websiteId,
		WebServerType: "nginx",
	}
}

func (w *WebsiteMgr) SetWwwRoot(www_root string) {
	w.WwwRoot = filepath.Join(global.SERVER_CONF.GetUserHomeDir(w.Username), strings.ToLower(www_root))
}

func (w *WebsiteMgr) CreateWebsite(website *models.ZapWebSite) error {
	if err := w.preCreateDirs(); err != nil {
		return err
	}

	// create index.html
	createWebsiteDefaultPages(w.WwwRoot)

	//write nginx conf
	userConfPath := w.GetUserWebserverConfDir()
	base.TryMkdir(userConfPath)

	ngxServConf := ngx.NewNgxConfServer(w.ServerName)
	ngxServConf.RunDirectory = website.RunDirectory
	ngxServConf.Root = w.WwwRoot

	err := ngxServConf.WriteToFile(filepath.Join(userConfPath, website.Domain+".conf"))
	if err != nil {
		return err
	}
	return ngx.CreateWebsiteIncludeFile(w.Username)
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

func (w *WebsiteMgr) RemoveWebsiteConfig(servername string) error {
	confFile := path.Join(w.GetUserWebserverConfDir(), strings.ToLower(servername)+".conf")
	if fileutils.IsFile(confFile) {
		if err := os.Remove(confFile); err != nil {
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

func (w *WebsiteMgr) RenameWwwRoot(oldpath, newpath string) error {
	if _, err := os.Stat(newpath); err != nil {
		return os.Rename(oldpath, newpath)
	}
	return fmt.Errorf("%s is exist", newpath)
}

// 更新网站
func (w *WebsiteMgr) UpdateWebsite(website *models.ZapWebSite) error {
	if err := w.preCreateDirs(); err != nil {
		return err
	}

	//write nginx conf
	userConfPath := w.GetUserWebserverConfDir()
	base.TryMkdir(userConfPath)
	ngxServConf := ngx.NewNgxConfServer(website.Domain)
	ngxServConf.AddServerName(strings.Fields(website.DomainNames)...)
	ngxServConf.SetUserDataPath(global.SERVER_CONF.GetUserHomeDir(w.Username))
	ngxServConf.RunDirectory = website.RunDirectory
	ngxServConf.Root = website.WwwRoot
	ngxServConf.Index = website.IndexFiles
	ngxServConf.AppName = website.Application
	ngxServConf.AppExposeName = website.ApplicationExposeName
	ngxServConf.AppExpose = website.ApplicationExpose

	err := ngxServConf.WriteToFile(filepath.Join(userConfPath, website.Domain+".conf"))
	if err != nil {
		return err
	}
	website.Status = "running"
	global.DB.Save(website)
	return ngx.CreateWebsiteIncludeFile(w.Username)
}

func (w *WebsiteMgr) GetUserWebserverConfDir() string {
	if w.WebServerType == "openresty" {
		return pathutil.GetPath("data/users", w.Username, "openresty_conf.d")
	} else if w.WebServerType == "httpd" {
		return pathutil.GetPath("data/users", w.Username, "httpd_conf.d")
	}

	return pathutil.GetPath("data/users", w.Username, "nginx_conf.d")
}

func (w *WebsiteMgr) GetUserNginxConfDir() string {
	return pathutil.GetPath("data/users", w.Username, "nginx_conf.d")
}

func (w *WebsiteMgr) GetWebsiteConfFileName() string {
	if w.WebServerType == "openresty" {
		return pathutil.GetPath("data/users", w.Username, "openresty_conf.d", w.ServerName+".conf")
	} else if w.WebServerType == "httpd" {
		return pathutil.GetPath("data/users", w.Username, "httpd_conf.d", w.ServerName+".conf")
	}

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
