package websites

import (
	"fmt"

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

func GetWebsiteByServerName(servername string) (*models.ZapWebSite, error) {
	website := models.ZapWebSite{}
	err := global.DB.Model(&models.ZapWebSite{}).Where("domain = ?", servername).First(&website).Error
	return &website, err
}

func SetWebsiteStatusByServerName(servername, status string) error {
	err := global.DB.Model(&models.ZapWebSite{}).Where("domain = ?", servername).Update("status = ?", status).Error
	return err
}

func CheckDomain(domain, websiteId, uid string) bool {
	website := models.ZapWebSite{}
	global.DB.Model(&models.ZapWebSite{}).Where("status = ?", "active").Or("domain = ?", domain).Or("domain_names LIKE ?", "%"+domain+"%").First(&website)
	if (websiteId == "" || websiteId == "0") && website.ID != 0 {
		return false
	} else if (websiteId != "0") && websiteId != fmt.Sprint(website.ID) && fmt.Sprint(website.Uid) != uid {
		return false
	}
	return true
}
