package websites

import (
	"log/slog"
	"os"
	"strings"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/services"
	"github.com/zapj/zap/core/utils/domainutil"
	"github.com/zapj/zap/core/utils/pathutil"
)

type WebSiteService struct {
	Uid      uint
	Username string
}

func NewWebSiteService(username string, id uint) *WebSiteService {
	return &WebSiteService{
		Username: username,
		Uid:      id,
	}
}

func (w *WebSiteService) CreateWebsite(req webSiteRequest) error {
	domain, domainNames := domainutil.ParseHostingDomain(req.Domain, req.DomainNames)
	req.Domain = domain
	req.DomainNames = strings.Join(domainNames, " ")
	mgr := NewWebSiteMgr(w.Username, req.Domain)
	mgr.SetWwwRoot(req.WwwRoot)
	website := models.ZapWebSite{
		Uid:               w.Uid,
		Domain:            req.Domain,
		Title:             req.Title,
		Description:       req.Description,
		DomainNames:       req.DomainNames,
		WwwRoot:           req.WwwRoot,
		RunDirectory:      req.RunDirectory,
		ApplicationId:     req.Application,
		ApplicationExpose: req.ExposePort,
		IndexFiles:        req.IndexFiles,
	}
	appService := services.NewAppsService()
	appInfo := appService.GetAppListById(website.ApplicationId)
	if p, err := appInfo.ParseExpose(); err == nil {
		for k, v := range p.Map() {
			if v == website.ApplicationExpose {
				website.ApplicationExpose = v
				website.ApplicationExposeName = k
				break
			}
		}
	}
	website.Application = appInfo.Name
	website.ApplicationVersion = appInfo.Version
	// global.DB.Model(&models.ZapWebSite{}).Create(&website)
	website.WwwRoot = mgr.WwwRoot

	if website.Title == "" {
		website.Title = website.Domain
	}
	slog.Info("website", "data", website)
	if err := global.DB.Save(&website).Error; err != nil {
		return err
	}
	err := mgr.CreateWebsite(&website)
	if err != nil {
		global.DB.Delete(&website)
	}
	return err
}

func (w *WebSiteService) CheckDomain(domain string) error {
	website := models.ZapWebSite{}
	err := global.DB.Where("domain = ?", domain).Or("domain_names like ?", "%"+domain+"%").First(&website).Error
	if err != nil {
		return err
	}
	return nil
}

func (w *WebSiteService) UpdateWebsite(req webSiteRequest) error {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("panic", "err", err)
		}
	}()
	website, err := GetWebsiteById(req.WebsiteId)
	if err != nil {
		return err
	}
	//parse domains
	domain, domainNames := domainutil.ParseHostingDomain(req.Domain, req.DomainNames)
	req.Domain = domain
	req.DomainNames = strings.Join(domainNames, " ")
	mgr := NewWebSiteMgr(w.Username, req.Domain)
	mgr.SetWwwRoot(req.WwwRoot)
	//删除旧网站
	if website.Domain != req.Domain {
		mgr.RemoveWebsiteConfig(website.Domain)
	}
	website.Domain = req.Domain
	website.DomainNames = req.DomainNames
	website.IndexFiles = req.GetIndexFiles()
	website.ApplicationId = req.Application
	website.ApplicationExposeName = req.ExposeProto
	website.ApplicationExpose = req.ExposePort
	if req.Title != "" {
		website.Title = req.Title
	}
	if website.Title == "" {
		website.Title = website.Domain
	}
	if website.WwwRoot != mgr.WwwRoot {
		mgr.RenameWwwRoot(website.WwwRoot, mgr.WwwRoot)
		website.WwwRoot = mgr.WwwRoot
	}

	appService := services.NewAppsService()
	appInfo := appService.GetAppListById(website.ApplicationId)
	if appInfo.Id != 0 {
		website.Application = appInfo.Name
		website.ApplicationVersion = appInfo.Version
		if p, err := appInfo.ParseExpose(); err == nil {
			for k, v := range p.Map() {
				if v == website.ApplicationExpose {
					website.ApplicationExpose = v
					website.ApplicationExposeName = k
					break
				}
			}
		}
	}

	if err := global.DB.Save(&website).Error; err != nil {
		return err
	}
	//修改配置文件
	err = mgr.UpdateWebsite(website)
	slog.Info("网站配置修改完成", "err", err)
	return err
}

func (w *WebSiteService) StopWebsite(websiteId int) error {

	website, err := GetWebsiteById(websiteId)
	if err != nil {
		return err
	}
	//remove conf
	if fileutils.IsFile(pathutil.GetPath("data/users", w.Username, "nginx_conf.d", website.Domain+".conf")) {
		if err := os.Remove(pathutil.GetPath("data/users", w.Username, "nginx_conf.d", website.Domain+".conf")); err != nil {
			return err
		}
	}
	website.Status = "suspend"
	return global.DB.Save(&website).Error
}

func (w *WebSiteService) StartWebsite(websiteId int) error {

	website, err := GetWebsiteById(websiteId)
	if err != nil {
		return err
	}

	website.Status = "running"
	global.DB.Save(&website)
	mgr := NewWebSiteMgr(w.Username, website.Domain)
	return mgr.UpdateWebsite(website)

}
