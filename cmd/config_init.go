package cmd

import (
	"github.com/spf13/cobra"
	"guayavita.dev/gvc/internal/logging"
	"guayavita.dev/gvc/pkg/config"
)

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application configuration",
	Long:  `This command initializes the application configuration by creating a new configuration file with default values.`,
	Run: func(cmd *cobra.Command, args []string) {
		logging.Log.Debug().Msg("Initializing configuration")
		config.DumpDefaultConfig()
	},
}
