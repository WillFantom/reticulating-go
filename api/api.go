package api

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/willfantom/loading-messages/sims"
)

var (
	pathPrefix string = "/api"
)

func API(router *mux.Router) {
	router.HandleFunc(pathPrefix+"/messages/random", randomMessage)
}

func randomMessage(w http.ResponseWriter, r *http.Request) {
	data := LoadingMessageResponse{
		Message: sims.LoadingMessages[rand.Intn(len(sims.LoadingMessages))],
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Errorln("⚠️	get random message request failed")
		http.Error(w, err.Error(), 500)
		return
	}
}
