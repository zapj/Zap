package websites

import (
	"log/slog"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/services"
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

	mgr := NewWebSiteMgr(w.Username, req.Domain)
	err := mgr.CreateWebsite(req)
	slog.Info("创建网站成功", "err", err, "mgr", mgr)

	if err != nil {
		// w.c.JSON(200, commons.Error(1, err.Error(), nil))
		return err
	}

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
	}
	appService := services.NewAppsService()
	appInfo := appService.GetAppListById(website.ApplicationId)
	if p, err := appInfo.ParseExpose(); err == nil {
		if p.MustGetString("socket") != website.ApplicationExpose {
			website.ApplicationExpose = p.MustGetString("socket")
		}
	}
	slog.Info("appInfo", "info", appInfo)
	website.Application = appInfo.Name
	website.ApplicationVersion = appInfo.Version
	// global.DB.Model(&models.ZapWebSite{}).Create(&website)
	website.WwwRoot = mgr.WwwRoot

	if website.Title == "" {
		website.Title = website.Domain
	}
	slog.Info("website", "data", website)
	return global.DB.Save(&website).Error
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

	website, err := GetWebsiteById(req.WebsiteId)
	if err != nil {
		return err
	}

	mgr := NewWebSiteMgr(w.Username, req.Domain)
	//删除旧网站
	if website.Domain != req.Domain {
		mgr.RemoveWebsiteConfig(website.Domain)
	}
	website.Domain = req.Domain
	website.DomainNames = req.DomainNames
	website.ApplicationId = req.Application
	if req.Title != "" {
		website.Title = req.Title
	}
	if website.Title == "" {
		website.Title = website.Domain
	}
	if website.WwwRoot != mgr.WwwRoot {
		website.WwwRoot = mgr.WwwRoot
		mgr.RenameWwwRoot(website.WwwRoot, mgr.WwwRoot)
	}

	appService := services.NewAppsService()
	appInfo := appService.GetAppListById(website.ApplicationId)
	if appInfo.Id != 0 {
		website.Application = appInfo.Name
		website.ApplicationVersion = appInfo.Version
		if p, err := appInfo.ParseExpose(); err == nil {
			if p.MustGetString("socket") != website.ApplicationExpose {
				website.ApplicationExpose = p.MustGetString("socket")
			}
		}
	}
	slog.Info("appInfo", "info", appInfo)
	slog.Info("website", "data", website)
	if err := global.DB.Save(&website).Error; err != nil {
		return err
	}
	//修改配置文件
	err = mgr.UpdateWebsite(website)
	slog.Info("网站配置修改完成", "err", err)
	return err
}
