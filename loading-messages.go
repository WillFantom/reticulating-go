package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/willfantom/loading-messages/api"
	"github.com/willfantom/loading-messages/sims"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	logrus.Infoln("🤨	starting sims server ")

	sims.LogStats()

	router := mux.NewRouter()
	api.API(router)
	logrus.Fatal(http.ListenAndServe(":8080", router))
}
