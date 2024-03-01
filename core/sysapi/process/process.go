package process

import (
	"bytes"
	"os"
	"strconv"
)

func GetPidFile(file string) int {
	if pidStr, err := os.ReadFile(file); err == nil {
		pid, _ := strconv.Atoi(string(bytes.Trim(pidStr, "\n")))
		return pid
	}
	return 0
}
