package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/bcrypt_util"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户管理",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "list":
			listUser()
		case "passwd":
			if len(args) == 3 {
				changePassword(args[1], args[2])
			} else {
				global.LOG.Println("zapctl user passwd [username] [password]")
			}
		case "check":
			if len(args) == 3 {
				checkPassword(args[1], args[2])
			} else {
				global.LOG.Println("zapctl user check [username] [password]")
			}
		default:
			global.LOG.Printf("不支持该参数 %s", args[0])
		}
	},
}

func listUser() {
	var result []models.ZapUsers
	global.DB.Model(&models.ZapUsers{}).Take(&result)
	for _, user := range result {
		fmt.Printf("--------------------------------------\n")
		fmt.Printf("ID: %d \n", user.ID)
		fmt.Printf("Uid: %d \n", user.Uid)
		fmt.Printf("Gid: %d \n", user.Gid)
		fmt.Printf("Username: %s \n", user.Username)
		fmt.Printf("Password: %s \n", user.Password)
		fmt.Printf("Home: %s \n", user.Home)
		fmt.Printf("Shell: %s \n", user.Shell)
		fmt.Println()
	}
}

func changePassword(username, password string) {
	user := models.ZapUsers{}
	err := global.DB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("%s 用户不存在\n", username)
		return
	}
	passwd, _ := bcrypt_util.HashPassword(password)
	user.Password = passwd
	global.DB.Save(&user)
	fmt.Printf("%s 密码修改成功\n", username)
}

func checkPassword(username, password string) {
	user := models.ZapUsers{}
	err := global.DB.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("%s 用户不存在\n", username)
		return
	}
	if bcrypt_util.CheckPasswordHash(password, user.Password) {
		fmt.Println("密码正确")
	} else {
		fmt.Println("密码错误")
	}

}
