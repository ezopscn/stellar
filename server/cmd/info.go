package cmd

import (
	"fmt"
	"stellar/common"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

// 项目信息
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about the server",
	Run: func(cmd *cobra.Command, args []string) {
		tb := table.NewWriter()
		tb.SetTitle("Basic Information")
		header := table.Row{"NAME", "VALUE"}
		tb.AppendHeader(header)
		rows := []table.Row{}
		rows = append(rows, table.Row{"Project Name", common.SystemProjectName})
		rows = append(rows, table.Row{"Project Description", common.SystemProjectDescription})
		rows = append(rows, table.Row{"System Version", common.SystemVersion})
		rows = append(rows, table.Row{"System Go Version", common.SystemGoVersion})
		rows = append(rows, table.Row{"Developer Name", common.SystemDeveloperName})
		rows = append(rows, table.Row{"Developer Email", common.SystemDeveloperEmail})
		tb.AppendRows(rows)
		fmt.Println(tb.Render())
	},
}
