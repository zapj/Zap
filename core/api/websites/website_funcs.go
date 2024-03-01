package websites

import (
	"errors"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

// CountWebsite count website
func CountWebsite(uid int) int {
	var count int64 = 0
	global.DB.Model(&models.ZapWebSite{}).Where("uid = ?", uid).Count(&count)
	return int(count)
}

func GetWebsiteById(websiteId int) (*models.ZapWebSite, error) {
	website := models.ZapWebSite{}
	err := global.DB.Model(&models.ZapWebSite{}).Where("id = ?", websiteId).First(&website).Error
	return &website, err
}

func CheckDomain(domain string) error {
	var count int64
	global.DB.Model(&models.ZapWebSite{}).Where("status = ?", "active").Or("domain = ?", domain).Or("domain_names LIKE ?", "%"+domain+"%").Count(&count)
	if count > 0 {
		return errors.New("域名已存在")
	}
	return nil
}
