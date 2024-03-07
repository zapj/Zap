package nginx

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/apps/general"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/cmdutil"
)

type App struct {
	general.App
}

func (a *App) Reload() error {
	retry := 10
	for i := 0; i < retry; i++ {
		err := nginxConfigTest()
		if err != nil && i == retry-1 {
			return fmt.Errorf("%s:%s %s:%s", err.ErrLevel, err.ErrMsg, err.ErrFile, err.ErrNum)
		} else if err != nil && fileutils.IsFile(err.ErrFile) {
			confName := filepath.Base(err.ErrFile)
			if !strings.HasSuffix(confName, ".conf") {
				continue
			}
			servername, ok := strings.CutSuffix(confName, ".conf")
			slog.Info("nginx reload ", "servername", servername)

			if ok {
				setWebsiteStatusByServerName(servername, "stop")
				os.Rename(err.ErrFile, err.ErrFile+".error")
			}

		}
	}

	_, err := cmdutil.ExecCmd("systemctl", "reload", a.Name)
	if err != nil {
		return err
	}
	return nil
}

func setWebsiteStatusByServerName(servername, status string) error {
	site := models.ZapWebSite{}
	global.DB.Model(&models.ZapWebSite{}).Where("domain = ?", servername).First(&site)
	slog.Info(" website info", "site", site)
	// r := global.DB.Model(&models.ZapWebSite{}).Where("domain = ?", servername).Update("status = ?", status)
	if site.ID != 0 {
		site.Status = status
		r := global.DB.Save(&site)
		return r.Error
	}
	return nil
}
