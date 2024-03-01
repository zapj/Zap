package websites

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/services"
	"github.com/zapj/zap/core/utils/zaputil"
)

type WebSiteService struct {
	c        *gin.Context
	Uid      uint
	Username string
}

func NewWebSiteService(c *gin.Context) *WebSiteService {
	return &WebSiteService{
		c:        c,
		Username: c.GetString("JWT_USERNAME"),
		Uid:      zaputil.MustConvertStringToUint(c.GetString("JWT_ID")),
	}
}

func (w *WebSiteService) CreateWebsite(req webSiteRequest) error {

	mgr := NewWebSiteMgr(w.c.GetString("JWT_USERNAME"), req.Domain)
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
	mgr := NewWebSiteMgrWithWebSiteId(w.c.GetString("JWT_USERNAME"), req.Domain, req.WebsiteId)
	err := mgr.UpdateWebsite(req)
	slog.Info("创建修改完成", "err", err, "mgr", mgr)

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
