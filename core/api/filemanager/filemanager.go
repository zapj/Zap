package filemanager

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/protect"
	"github.com/zapj/zap/core/utils/time_utils"
)

func FileManager_List(c *gin.Context) {
	basePath := filepath.Clean(c.PostForm("path"))
	page, err := strconv.Atoi(c.DefaultPostForm("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pagesize, err := strconv.Atoi(c.DefaultPostForm("page_size", "50"))
	if err != nil || pagesize < 1 {
		pagesize = 50
	}

	if basePath == "" {
		basePath = "/"
	}
	dirFs, err := os.ReadDir(basePath)
	if err != nil {
		slog.Info("readdir", "err", err)
	}
	total := len(dirFs)
	totalpage := total % pagesize
	if totalpage == 0 {
		totalpage = total / pagesize
	} else {
		totalpage = (total / pagesize) + 1
	}
	if totalpage > 0 && page > totalpage {
		page = totalpage
	}
	offset := (page - 1) * pagesize
	limit := offset + pagesize
	if limit > total {
		limit = total
	}
	// slog.Info("basePath", basePath, "    total:", total, "  page", page, "   Limit:", limit, "   Offset:", offset)
	fileList := []gin.H{}
	for _, dirEntry := range dirFs[offset:limit] {
		fileInfo, _ := dirEntry.Info()
		is_symlink := fileInfo.Mode()&os.ModeSymlink != 0
		symlink := ""
		is_dir := fileInfo.IsDir()
		is_file := dirEntry.Type().IsRegular()
		if is_symlink {
			symLinkInfo, _ := os.Stat(path.Join(basePath, dirEntry.Name()))
			symlink, _ = os.Readlink(path.Join(basePath, dirEntry.Name()))
			// symlink = symLinkInfo.Name()
			is_dir = symLinkInfo.IsDir()
			is_file = symLinkInfo.Mode().IsRegular()
		}
		fileList = append(fileList, gin.H{
			"name":       dirEntry.Name(),
			"is_dir":     is_dir,
			"is_file":    is_file,
			"is_symlink": is_symlink,
			"symlink_to": symlink,
			"mod_time":   fileInfo.ModTime().Format(time_utils.DATE_TIME_FORMAT),
			"perm":       fmt.Sprintf("%#o", fileInfo.Mode().Perm()),
			"uid":        fileInfo.Sys().(*syscall.Stat_t).Uid,
			"gid":        fileInfo.Sys().(*syscall.Stat_t).Gid,
			"filesize":   humanize.Bytes(uint64(fileInfo.Size())),
		})
	}
	c.JSON(200, gin.H{
		"code":       0,
		"data":       fileList,
		"path":       basePath,
		"total":      total,
		"page":       page,
		"page_size":  pagesize,
		"total_page": totalpage,
	})
}

func FileManager_Fetch(c *gin.Context) {
	filename := c.PostForm("filename")
	mimeType, _ := mimetype.DetectFile(filename)
	if !protect.IsTextFile(mimeType.String()) {
		c.JSON(200, gin.H{"code": 1, "msg": filename + "不是文本类型"})
		return
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": filename + "无法读取文件内容" + err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "OK", "data": string(data), "mimetype": mimeType.String()})
}

func FileManager_PutFile(c *gin.Context) {
	filename := c.PostForm("filename")
	content := c.PostForm("content")
	finfo, err := os.Stat(filename)
	fmode := os.FileMode(0644)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		c.JSON(200, gin.H{"code": 1, "msg": filename + "无法保存" + err.Error()})
		return
	} else if err != nil && errors.Is(err, os.ErrNotExist) {
		fmode = finfo.Mode()
	}
	err = os.WriteFile(filename, []byte(content), fmode)
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": filename + "无法写入文件内容" + err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "文件保存成功"})
}
