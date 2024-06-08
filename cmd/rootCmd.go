package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"guayavita.dev/gvc/pkg/config"
	"tinygo.org/x/go-llvm"
)

var rootCmd = &cobra.Command{
	Use:     "gvc",
	Version: fmt.Sprintf("GVC 0.0.1 LLVM %s %t \n", llvm.Version, config.GVConfig.Core.Verbose),
	Short:   "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
}
