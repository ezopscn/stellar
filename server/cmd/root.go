package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stellar",
	Short: "Stellar is a server for the Stellar network",
}

func Execute() error {
	return rootCmd.Execute()
}
