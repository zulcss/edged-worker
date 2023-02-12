package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/zulcss/edged/pkg/utils"
	"github.com/zulcss/edged/pkg/service"
)

var (
	Verbose		bool
	Config		string
)



var rootCmd = &cobra.Command{
	Use:	"edged-worker",
	Short:  "StarlingX edged api worker",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := InitLog(); err != nil {
			return err
		}

		// Run as root
		if !utils.CheckUser() {
			fmt.Println("Please run as root")
			os.Exit(-1)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		service.Run(Config)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to execute worker: %v", err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&Config, "config", "c", "etc/edged/config.toml", "Server config file")
}

func InitLog() error {
        if Verbose {
                logrus.SetLevel(logrus.DebugLevel)
        }

        // general log output
        formatter := &logrus.TextFormatter{
                FullTimestamp:   true,
                TimestampFormat: "2006-01-02 15:04:05",
        }
        logrus.SetFormatter(formatter)

        return nil
}
