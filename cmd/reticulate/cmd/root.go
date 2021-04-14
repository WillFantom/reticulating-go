package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/reticulating-go"
)

var (
	debug bool
	trace bool
)

var rootCmd = &cobra.Command{
	Use:   "reticulate",
	Short: "Get a Sims loading message",
	Long:  `Get loading messages from The Sims in the CLI or server via an API`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if trace {
			logrus.SetLevel(logrus.TraceLevel)
		} else if debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(reticulating.GetLoadingMessage())
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set the log level to debug")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "set the log level to trace")
	apiCmd.PersistentFlags().StringVar(&apiPort, "port", "8080", "port for the api to listen on")
	rootCmd.AddCommand(apiCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalln("🆘	root command failed")
	}
}
