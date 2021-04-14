package sims

import (
	"errors"
	"math/rand"
)

type Game struct {
	Name           string   `json:"Name"`
	FandomURL      string   `json:"FandomURL"`
	ExpansionPacks []*Game  `json:"Expansion Packs"`
	StuffPacks     []*Game  `json:"Stuff Packs"`
	Messages       []string `json:"Messages"`
}

func (g *Game) GetRandomExpansionPack() (*Game, error) {
	if len(g.ExpansionPacks) > 0 {
		return g.ExpansionPacks[rand.Intn(len(g.ExpansionPacks))], nil
	}
	return nil, errors.New("game has no expansion packs")
}

func (g *Game) GetRandomStuffPack() (*Game, error) {
	if len(g.StuffPacks) > 0 {
		return g.StuffPacks[rand.Intn(len(g.StuffPacks))], nil
	}
	return nil, errors.New("game has no stuff packs")
}

func (g *Game) GetRandomSubgame() (*Game, error) {
	if rand.Intn(1) == 0 {
		return g.GetRandomExpansionPack()
	}
	return g.GetRandomStuffPack()
}

func (g *Game) GetExpansionPack(name string) (*Game, error) {
	for _, expansion := range g.ExpansionPacks {
		if expansion.Name == name {
			return expansion, nil
		}
	}
	return nil, errors.New("given expansion pack not in game")
}

func (g *Game) GetStuffPack(name string) (*Game, error) {
	for _, pack := range g.StuffPacks {
		if pack.Name == name {
			return pack, nil
		}
	}
	return nil, errors.New("given stuff pack not in game")
}
