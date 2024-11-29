package cmd

import (
	"fmt"
	"stellar/common"
	"stellar/initialize"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initSystemCmd)
}

// 初始化命令
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the server",
}

// 初始化系统命令
var initSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Initialize the system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start initialize system: ", time.Now().Format(common.TimeMillisecondFormat))
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateTable()
		initialize.MigrateData()
		fmt.Println("Initialize system completed: ", time.Now().Format(common.TimeMillisecondFormat))
	},
}
