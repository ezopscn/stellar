package cmd

import (
	"stellar/common"
	"stellar/initialize"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateTableCmd)
	migrateTableCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SystemConfigFilename, "The path to the configuration file")
	migrateCmd.AddCommand(migrateDataCmd)
	migrateDataCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SystemConfigFilename, "The path to the configuration file")
}

// 迁移命令
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
}

// 迁移表命令
var migrateTableCmd = &cobra.Command{
	Use:   "table",
	Short: "Migrate the database table",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateTable()
	},
}

// 迁移数据命令
var migrateDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Migrate the database data",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateData()
	},
}
