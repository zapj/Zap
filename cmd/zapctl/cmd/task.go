package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zdate"
	"github.com/zapj/zap/core/zapi"
)

func init() {
	rootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "task manager",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "list":
			listTaskJob()
		case "put":
			putTask()
		case "refresh":
			refreshTask()
		}
	},
}

func listTaskJob() {
	result := []models.ZapTask{}
	global.DB.Find(&result)
	if len(result) == 0 {
		fmt.Println("暂无任务")
	}
	for _, task := range result {
		fmt.Printf("--------------------------------------\n")
		fmt.Printf("ID: %d \n", task.Id)
		fmt.Printf("Title: %s \n", task.Title)
		fmt.Printf("Cmd: %s \n", task.Cmd)
		fmt.Printf("Retry: %d \n", task.Retry)
		fmt.Printf("RetryCount: %d \n", task.RetryCount)
		fmt.Printf("LogFile: %s \n", task.LogFile)
		fmt.Printf("TargetDir: %s \n", task.TargetDir)
		fmt.Printf("Status: %s \n", task.Status)
		fmt.Printf("StartTime: %s \n", time.Unix(task.StartTime, 0).Format(zdate.DATE_TIME_FORMAT))
		fmt.Printf("EndTime: %s \n", time.Unix(task.EndTime, 0).Format(zdate.DATE_TIME_FORMAT))

		fmt.Println()
	}
}

func putTask() {
	// _, err := task.Post("/task/put")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func refreshTask() {
	resp := zapi.NewZapi().RefreshTask()
	if resp.HasError() {
		fmt.Printf("Error Message : %s\n", resp.ErrorMessage())
		return
	}
	if resp.Status == 200 {
		r := resp.ZapMap()
		fmt.Printf("code:%s , msg : %s\n", r.GetString("code"), r.GetString("msg"))
	}
}
