package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/willfantom/reticulating-go"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get some stats about reticulating",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("|---- Reticulating Stats")
		fmt.Printf("\n🕹\tGames: %d\n", reticulating.GameCount())
		fmt.Printf("📦\tExpansions: %d\n", reticulating.PackCount())
		fmt.Printf("💬\tMessages: %d\n\n", reticulating.MessageCount())
		fmt.Println("|-----------------------")
	},
}
