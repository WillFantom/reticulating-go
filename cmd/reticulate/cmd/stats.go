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
		fmt.Printf("\nš¹\tGames: %d\n", reticulating.GameCount())
		fmt.Printf("š¦\tExpansions: %d\n", reticulating.PackCount())
		fmt.Printf("š¬\tMessages: %d\n\n", reticulating.MessageCount())
		fmt.Println("|-----------------------")
	},
}
