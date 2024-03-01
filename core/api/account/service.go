package account

import (
	"fmt"

	"github.com/zapj/go-properties"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

type AccountService struct {
	Username string `json:"username"`
	HomeDir  string `json:"home_dir"`
	User     string `json:"user"`
	Group    string `json:"group"`
	UserId   int    `json:"user_id"`
	GroupId  int    `json:"group_id"`
}

func NewAccountService(username string) *AccountService {
	return &AccountService{
		Username: username,
		HomeDir:  global.SERVER_CONF.GetUserHomeDir(username),
		User:     global.SERVER_CONF.WwwUser,
		Group:    global.SERVER_CONF.WwwGroup,
		UserId:   global.SERVER_CONF.WwwUserId,
		GroupId:  global.SERVER_CONF.WwwGroupId,
	}

}

func (a *AccountService) GetUserHomeDir() string {
	return a.HomeDir
}

func (a *AccountService) GetWebsiteConfig() base.ZapMap {
	return base.ZapMap{
		"home_dir":     a.HomeDir,
		"username":     a.Username,
		"applications": getAllAppList(),
	}
}

// 获取所有可用中间件服务
func getAllAppList() []base.ZapMap {
	appList := []models.ZapApps{}
	global.DB.Model(&models.ZapApps{}).Where("status = ?", "active").Where("app_type = ?", "application").Find(&appList)
	var result []base.ZapMap

	//static website
	result = append(result, base.ZapMap{
		"id":           0,
		"name":         "static",
		"title":        "静态网站",
		"app_status":   "running",
		"expose_proto": "http",
		"expose_port":  "",
	})

	for _, v := range appList {
		if v.Expose == "" {
			continue
		}

		p, err := properties.LoadString(v.Expose)
		if err != nil || p.Len() < 1 {
			continue
		}

		for key, val := range p.Map() {
			result = append(result, base.ZapMap{
				"id":           v.Id,
				"name":         v.Name,
				"title":        fmt.Sprintf("%s %s", v.Title, v.Version),
				"app_status":   v.AppStatus,
				"expose_proto": key,
				"expose_port":  val,
			})
		}

	}
	return result
}
