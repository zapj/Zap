package services

import (
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

type AppsService struct {
}

func NewAppsService() *AppsService {
	return &AppsService{}
}

func (a *AppsService) GetAppList() []models.ZapApps {
	var appList []models.ZapApps
	global.DB.Where("status = ?", "active").Find(&appList)
	return appList
}

func (a *AppsService) GetAppListById(id int) *models.ZapApps {
	var app models.ZapApps
	global.DB.Where("id = ?", id).First(&app)
	return &app
}

func (a *AppsService) GetAppListByName(name string) []models.ZapApps {
	var appList []models.ZapApps
	global.DB.Where("name = ?", name).Find(&appList)
	return appList
}

func (a *AppsService) GetAppByNameAndVersion(name, version string) (models.ZapApps, error) {
	var app models.ZapApps
	err := global.DB.Where("name = ? and version = ?", name, version).First(&app).Error
	return app, err
}
