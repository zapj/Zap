package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zaputil"
)

func init() {
	// updateCmd.PersistentFlags().StringP("data", "d", "", "json data: {}")
	updateCmd.PersistentFlags().StringP("keypair", "k", "", " -k \"key=value&a=1\" ")
	updateCmd.PersistentFlags().StringP("jsondata", "j", "", " -j {\"key=value\"} ")
	updateCmd.PersistentFlags().StringP("where", "w", "", " -w \"name=value\" ")
	updateCmd.PersistentFlags().StringP("settings", "s", "", " -s \"name=value\" ")
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "update zap database",
	Example: "zapctl update [table]  -j {\"key=value\"} -k a=1&b=2 -s k=1\\nv=2 -w \"id=123\" ",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// data, _ := cmd.PersistentFlags().GetString("data")
		where, _ := cmd.PersistentFlags().GetString("where")
		jsondata, _ := cmd.PersistentFlags().GetString("data")
		keypair, err := cmd.PersistentFlags().GetString("keypair")
		if err != nil {
			fmt.Println(err)
			return
		}
		settings, _ := cmd.PersistentFlags().GetString("settings")
		if jsondata != "" {
			if ok := updateTable(args[0], jsondata, where); ok {
				fmt.Println("更新成功")
			} else {
				fmt.Println("更新失败")
			}
		}
		if keypair != "" {
			keyPairs := zaputil.ParseKeyPair(keypair)
			if len(keyPairs) > 0 {
				updateTableWithMap(where, keyPairs)
			}
		}
		if settings != "" {
			err := global.DB.Model(&models.ZapApps{}).Where(where).Updates(map[string]any{"settings": settings}).Error
			if err != nil {
				fmt.Println("更新失败", err)
				return
			}
			fmt.Println("更新成功")
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

func updateTableWithMap(where string, changes map[string]any) {
	fmt.Println(changes)
	err := global.DB.Model(&models.ZapApps{}).Where(where).Updates(changes).Error
	if err != nil {
		fmt.Println(err)
		return
	}
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

func transcode(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
