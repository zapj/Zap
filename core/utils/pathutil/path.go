package pathutil

import (
	"path/filepath"

	"github.com/zapj/zap/core/global"
)

func GetPath(elem ...string) string {
	if len(elem) == 1 {
		return filepath.Join(global.ZAP_BASE_DIR, elem[0])
	}
	elem = append([]string{global.ZAP_BASE_DIR}, elem...)
	return filepath.Join(elem...)
}
