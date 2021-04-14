package cmd

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/willfantom/reticulating-go"
)

type LoadingMessageResponse struct {
	BaseGame string `json:"base_game,omitempty"`
	SubGame  string `json:"sub_game,omitempty"`
	Message  string `json:"message"`
}

const (
	pathPrefix string = "/api"
)

var (
	apiPort string
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Serve loading messages over http",
	Run: func(cmd *cobra.Command, args []string) {
		apiRouter := mux.NewRouter()
		apiRouter.HandleFunc(pathPrefix+"/messages/random", randomMessage)
		apiRouter.HandleFunc(pathPrefix+"/messages/{game}/random", gameRandomMessage)
		apiRouter.HandleFunc(pathPrefix+"/messages/{game}/{pack}/random", packRandomMessage)
		logrus.Fatalln(http.ListenAndServe(":"+apiPort, apiRouter))
	},
}

func randomMessage(w http.ResponseWriter, r *http.Request) {
	logrus.Traceln("📩	request for loading message")
	data := LoadingMessageResponse{
		Message: reticulating.GetLoadingMessage(),
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Errorln("⚠️	get random message request failed")
		http.Error(w, err.Error(), 500)
		return
	}
}

func gameRandomMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logrus.Traceln("📩	request for loading message from ", vars["game"])
	gameName, _ := url.QueryUnescape(vars["game"])
	game := reticulating.GameByName(gameName)
	if game == nil {
		logrus.Errorln("⚠️	get game's random message request failed")
		http.Error(w, "game does not exist", 500)
		return
	}
	data := LoadingMessageResponse{
		Message: game.GetRandomMessage(),
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Errorln("⚠️	get game's random message request failed")
		http.Error(w, err.Error(), 500)
		return
	}
}

func packRandomMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logrus.Traceln("📩	request for loading message from ", vars["game"])
	game := reticulating.GameByName(vars["game"])
	if game == nil {
		logrus.Errorln("⚠️	get game's random message request failed")
		http.Error(w, "game does not exist", 500)
		return
	}
	pack := game.PackByName(vars["pack"])
	if pack == nil {
		logrus.Errorln("⚠️	get pack's random message request failed")
		http.Error(w, "pack does not exist", 500)
		return
	}
	data := LoadingMessageResponse{
		Message: pack.GetRandomMessage(),
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Errorln("⚠️	get pack's random message request failed")
		http.Error(w, err.Error(), 500)
		return
	}
}
