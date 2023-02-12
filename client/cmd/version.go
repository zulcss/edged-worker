package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zulcss/edged/pkg/constants"
)

var versionCommand = &cobra.Command{
	Use:	"version",
	Short:	"Display the CLI version",
	Run: func(cmd *cobra.Command, args[] string) {
		fmt.Printf("Version: %s\n", constants.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
