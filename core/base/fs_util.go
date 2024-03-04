package base

import (
	"os"
)

func TryMkdir(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, 0755)
	}
	return false
}
