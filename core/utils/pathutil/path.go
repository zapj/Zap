package pathutil

import (
	"path/filepath"

	"github.com/zapj/zap/core/global"
)

func GetPath(filename string) string {
	return filepath.Join(global.ZAP_BASE_DIR, filename)
}
