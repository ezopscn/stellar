package cmd

import (
	"fmt"
	"stellar/common"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version of the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("System Version:", common.SystemVersion)
		fmt.Println("System Go Version:", common.SystemGoVersion)
	},
}
