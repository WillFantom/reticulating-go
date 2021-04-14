package reticulating

import (
	"testing"
)

func BenchmarkFlatten(b *testing.B) {
	flattenGameMessages(data)
}

func TestFlatten(t *testing.T) {
	fakeGames := []*Game{
		{
			Messages: []string{
				"a",
				"b",
			},
			Packs: []*Game{
				{
					Messages: []string{
						"c",
						"d",
					},
				},
			},
		},
	}
	messages := flattenGameMessages(fakeGames)
	t.Logf("Messages: %s\n", messages)
	if len(messages) != 4 {
		t.Fail()
	}
}

func TestGetRandomGame(t *testing.T) {
	var games []string
	for i := 0; i < 3; i++ {
		games = append(games, GetRandomGame(data).Name)
	}
	t.Logf("Random Games: %v\n", games)
}

func TestGetRandomPack(t *testing.T) {
	var packs []string
	for i := 0; i < 5; i++ {
		game := GetRandomGame(data)
		packs = append(packs, game.GetRandomPack().Name)
	}
	t.Logf("Random Packs: %v\n", packs)
}

func TestGetRandomMessage(t *testing.T) {
	var messages []string
	for i := 0; i < 5; i++ {
		messages = append(messages, GetLoadingMessage())
	}
	t.Logf("Random Messages (nonflattened): %v\n", messages)
}
