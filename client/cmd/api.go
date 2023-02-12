package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var apiCommand = &cobra.Command{
	Use: 	"api",
	Short:	"Query Edged API worker",
}

var apiShow = &cobra.Command{
	Use:	"show",
	Short:	"Show Edged API information",
	Run: func(cmd *cobra.Command, args[] string) {
		fmt.Println(Server)
	},
}

func init() {
	rootCmd.AddCommand(apiCommand)

	apiCommand.AddCommand(apiShow)

}
