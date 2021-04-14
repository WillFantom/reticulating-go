package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/reticulating-go"
)

var (
	debug   bool
	trace   bool
	instant bool
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
		if !instant {
			s := spinner.New(spinner.CharSets[34], 50*time.Millisecond)
			s.Start()
			wait := time.Millisecond * 100
			for wait <= time.Second {
				s.Suffix = fmt.Sprintf("\t%s", reticulating.GetLoadingMessage())
				time.Sleep(wait)
				wait = wait + (time.Millisecond * 100)
			}
			s.Stop()
		}
		fmt.Printf("\033[2K\r%s\n", reticulating.GetLoadingMessage())
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set the log level to debug")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "set the log level to trace")
	rootCmd.Flags().BoolVarP(&instant, "instant", "i", false, "skip the animation")
	apiCmd.PersistentFlags().StringVar(&apiPort, "port", "8080", "port for the api to listen on")
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(statsCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalln("🆘	root command failed")
	}
}
