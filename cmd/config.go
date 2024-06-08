package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Handle the gvc configuration",
	Long:  `This command initializes gvc configuration by creating a new configuration file with default values.`,
}

func init() {
	configCmd.AddCommand(initConfigCmd)
}
