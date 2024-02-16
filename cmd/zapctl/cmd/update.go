package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func init() {
	// updateCmd.PersistentFlags().StringP("data", "d", "", "json data: {}")
	updateCmd.PersistentFlags().StringP("data", "d", "", " -d \"key=value\" ")
	updateCmd.PersistentFlags().StringP("where", "w", "", " -w \"name=value\" ")
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "update zap database",
	Example: "zapctl update [table]  -d \"key=value\" -w \"id=123\" ",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// data, _ := cmd.PersistentFlags().GetString("data")
		where, _ := cmd.PersistentFlags().GetString("where")
		data, _ := cmd.PersistentFlags().GetString("data")
		// json_data := jsonutil.DecodeToZapMap(data)

		if ok := updateTable(args[0], data, where); ok {
			fmt.Println("更新成功")
		} else {
			fmt.Println("更新失败")
		}
	},
}

func updateTable(table string, data string, where string) bool {
	model := getModels(table)
	err := json.Unmarshal([]byte(data), model)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = global.DB.Where(where).Updates(model).Error
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
