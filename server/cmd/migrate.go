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
	Short: "Migrate the database, only first time initialize system data",
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
	Short: "Dangerous! Migrate the database data, will delete all data and insert new data",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateData()
	},
}
