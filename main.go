package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:	"edged-worker",
	Short:  "StarlingX edged api worker",
	Run: func(cmd *cobra.Command, args []string) {},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to execute worker: %v", err)
		os.Exit(-1)
	}
}

func init() {}
