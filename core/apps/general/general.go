package general

import (
	"github.com/zapj/zap/core/apps/types"
	"github.com/zapj/zap/core/utils/cmdutil"
)

type App struct {
	Name string // app name （service name）
}

func (a *App) Stop() error {
	_, err := cmdutil.ExecCmd("systemctl", "stop", a.Name)
	if err != nil {
		return err
	}
	return nil
}
func (a *App) Start() error {
	_, err := cmdutil.ExecCmd("systemctl", "start", a.Name)
	if err != nil {
		return err
	}
	return nil
}
func (a *App) Restart() error {
	_, err := cmdutil.ExecCmd("systemctl", "restart", a.Name)
	if err != nil {
		return err
	}
	return nil
}
func (a *App) Reload() error {
	_, err := cmdutil.ExecCmd("systemctl", "reload", a.Name)
	if err != nil {
		return err
	}
	return nil
}

func NewAppWithName(name string) types.IApp {
	return &App{Name: name}
}
