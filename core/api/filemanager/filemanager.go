package filemanager

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/time_utils"
)

func FileManager_List(c *gin.Context) {
	basePath := filepath.Clean(c.PostForm("path"))
	if basePath == "" {
		basePath = "/"
	}
	dirFs, err := os.ReadDir(basePath)
	if err != nil {
		global.LOG.Info(err)
	}

	fileList := []gin.H{}
	for _, dirEntry := range dirFs {
		fileInfo, _ := dirEntry.Info()
		is_symlink := fileInfo.Mode()&os.ModeSymlink != 0
		symlink := ""
		is_dir := fileInfo.IsDir()
		if is_symlink {
			symLinkInfo, _ := os.Stat(path.Join(basePath, dirEntry.Name()))
			symlink, _ = os.Readlink(path.Join(basePath, dirEntry.Name()))
			// symlink = symLinkInfo.Name()
			is_dir = symLinkInfo.IsDir()
		}
		fileList = append(fileList, gin.H{
			"name":       dirEntry.Name(),
			"is_dir":     is_dir,
			"is_file":    dirEntry.Type().IsRegular(),
			"is_symlink": is_symlink,
			"symlink_to": symlink,
			"mod_time":   fileInfo.ModTime().Format(time_utils.DATE_TIME_FORMAT),
			"perm":       fmt.Sprintf("%#o", fileInfo.Mode().Perm()),
			"uid":        fileInfo.Sys().(*syscall.Stat_t).Uid,
			"gid":        fileInfo.Sys().(*syscall.Stat_t).Gid,
			"filesize":   humanize.Bytes(uint64(fileInfo.Size())),
		})
	}
	c.JSON(200, gin.H{"code": 0, "data": fileList})
}
