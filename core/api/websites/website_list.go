package websites

import (
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
)

func ListWebsite(c *gin.Context) {

	page := zaputil.MustConvertStringToInt(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	pagesize, err := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	if err != nil || pagesize < 1 {
		pagesize = 50
	}
	var totalCount int64
	global.DB.Model(&models.ZapWebSite{}).Where("uid = ?", c.GetString("JWT_ID")).Count(&totalCount)
	total := int(totalCount)
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

	websites := []models.ZapWebSite{}
	slog.Info("list website", "offset", offset, "limit", limit)
	global.DB.Where("uid = ?", c.GetString("JWT_ID")).Offset(offset).Limit(limit).Find(&websites)
	c.JSON(200, gin.H{
		"code":       0,
		"msg":        "success",
		"total":      total,
		"page":       page,
		"page_size":  pagesize,
		"total_page": totalpage,
		"data":       websites,
	})
}
