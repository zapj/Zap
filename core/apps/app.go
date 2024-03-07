package apps

import (
	"github.com/zapj/zap/core/apps/general"
	"github.com/zapj/zap/core/apps/nginx"
	"github.com/zapj/zap/core/apps/types"
)

func NewApp(name string) types.IApp {
	comm := general.App{Name: name}
	switch name {
	case "nginx":
		return &nginx.App{App: comm}
	}
	return &comm
}
