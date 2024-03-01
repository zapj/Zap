package base

import (
	"os"
)

func TryMkdir(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil && fi.IsDir() {
		os.MkdirAll(path, 0755)
	}
	return false
}
