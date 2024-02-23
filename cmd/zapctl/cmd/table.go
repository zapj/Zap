package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
)

func init() {
	// updateCmd.PersistentFlags().StringP("data", "d", "", "json data: {}")
	tableCmd.PersistentFlags().StringToStringP("data", "d", nil, " -d key=value\\na=b ")
	tableCmd.PersistentFlags().StringP("where", "w", "", " -w \"name=value\" ")
	rootCmd.AddCommand(tableCmd)
}

var tableCmd = &cobra.Command{
	Use:     "table",
	Short:   "update table data",
	Example: "zapctl table [table name]  -d k=1,v=2 -w \"id=123\" ",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := cmd.PersistentFlags().GetStringToString("data")
		where, _ := cmd.PersistentFlags().GetString("where")
		if len(data) > 0 && where != "" {
			if ok := updateTable(args[0], jsonutil.EncodeToString(data), where); ok {
				fmt.Println("更新成功")
			} else {
				fmt.Println("更新失败")
			}
		}

	},
}

func updateTable(table string, data string, where string) bool {
	model := getModels(table)
	var colsData map[string]any
	err := json.Unmarshal([]byte(data), &colsData)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = global.DB.Model(&model).Where(where).Updates(colsData).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func getModels(table string) any {
	switch table {
	case "apps":
		return &models.ZapApps{}
	case "users":
		return &models.ZapUsers{}
	case "crontab":
		return &models.ZapCrontab{}
	case "web_site":
		return &models.ZapWebSite{}
	case "data_base":
		return &models.ZapDataBase{}
	case "app_store":
		return &models.ZapAppStore{}
	case "access_log":
		return &models.ZapAccessLog{}
	}
	return nil
}
