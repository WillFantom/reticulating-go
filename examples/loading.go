package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pterm/pterm"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func main() {
	theSims, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("The", pterm.NewStyle(pterm.FgGray)),
		pterm.NewLettersFromStringWithStyle("Sims", pterm.NewStyle(pterm.FgGreen))).
		Srender()
	pterm.DefaultCenter.Print(theSims)
	pterm.Println()
	spinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("20 seconds of garbage messages...")
	time.Sleep(time.Second / 2)
	for i := 39; i > 0; i-- {
		response, err := http.Get("https://sims.willfantom.com/api/messages/random")
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		var messageResponse MessageResponse
		if err := json.NewDecoder(response.Body).Decode(&messageResponse); err != nil {
			panic(err)
		}
		spinner.UpdateText(messageResponse.Message)
		time.Sleep(time.Second / 2)
	}
	spinner.Stop()
}
