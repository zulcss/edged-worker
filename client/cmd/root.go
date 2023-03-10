package cmd

import (
	"os"
	"github.com/spf13/cobra"
        "github.com/sirupsen/logrus"
)

var (
	Verbose 	bool
	Server		string
	Port		int
)

var rootCmd = &cobra.Command{
	Use: 	"edgectl",
	Short:	"StarlingX edge cli",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := InitLog(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&Server, "server", "s", "localhost", "Server to connect to")
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
