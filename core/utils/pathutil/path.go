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

func InSecurePaths(path string) bool {
	var paths = []string{
		"",
		"/",
		"/sbin",
		"/bin",
		"/run",
		"/opt",
		"/var",
		"/usr",
		"/usr/sbin",
		"/usr/bin",
		"/usr/local",
		"/usr/local/sbin",
		"/usr/local/bin",
		"/usr/local/zap",
		"/usr/local/zap/sbin",
		"/usr/local/zap/bin",
		"/usr/local/zap/sbin",
		"/usr/local/zap/bin",
	}
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}
