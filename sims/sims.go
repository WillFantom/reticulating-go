package sims

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"

	"github.com/sirupsen/logrus"
)

const (
	dataPath        string = "./gamedata.json"
	minimumMessages int    = 1
	minimimGames    int    = 1
)

var (
	totalGames      int
	Games           []*Game
	LoadingMessages []string
)

func init() {

	logrus.Infoln("👀	parsing possible platitudes")

	dataFile, err := ioutil.ReadFile(dataPath)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err.Error(),
		}).Fatalln("🆘	could not read in messages file")
	}

	if err := json.Unmarshal([]byte(dataFile), &Games); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err.Error(),
		}).Fatalln("🆘	could not read messages file as a json")
	}

	for _, game := range Games {
		flattenData(game)
	}

	if totalGames < minimimGames {
		logrus.WithField("games", totalGames).Fatalln("🆘	not enough games found")
	}
	if len(LoadingMessages) < minimumMessages {
		logrus.WithField("messages", len(LoadingMessages)).Fatalln("🆘	not enough loading messages found")
	}

}

func LogStats() {
	logrus.WithFields(logrus.Fields{
		"total games":                len(Games),
		"total games (inc subgames)": totalGames,
		"total messages":             len(LoadingMessages),
	}).Infoln("📖	sending you sims stats ")
}

func flattenData(g *Game) {
	totalGames++
	LoadingMessages = append(LoadingMessages, g.Messages...)
	for _, expansion := range g.ExpansionPacks {
		flattenData(expansion)
	}
	for _, stuff := range g.StuffPacks {
		flattenData(stuff)
	}
}

func GetRandomGame() (*Game, error) {
	if len(Games) > 0 {
		return Games[rand.Intn(len(Games))], nil
	}
	return nil, errors.New("no games")
}
