package main

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/willfantom/reticulating-go"
)

func main() {
	theSims, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("The", pterm.NewStyle(pterm.FgGray)),
		pterm.NewLettersFromStringWithStyle("Sims", pterm.NewStyle(pterm.FgGreen))).
		Srender()
	pterm.Print(theSims)
	pterm.Println()
	spinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("20 seconds of garbage messages...")
	time.Sleep(time.Second)
	for i := 19; i > 0; i-- {
		spinner.UpdateText(reticulating.GetLoadingMessage())

		time.Sleep(time.Second)
	}
	spinner.Success()
}
