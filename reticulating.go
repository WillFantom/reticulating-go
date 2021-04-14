package reticulating

var (
	gameNameList    []string
	flatMessageList []string
)

func init() {
	gameNameList = make([]string, 0, len(data))
	flatMessageList = flattenGameMessages(data)
}

func GetLoadingMessage() string {
	return flatMessageList[simsRand.Intn(len(flatMessageList))]
}

func GameNameList() []string {
	if gameNameList == nil {
		for idx, game := range data {
			gameNameList[idx] = game.Name
		}
	}
	return gameNameList
}

func GetPacksNameList(game *Game) []string {
	if game.Packs == nil {
		return []string{}
	}
	packNames := make([]string, 0, len(game.Packs))
	for idx, game := range game.Packs {
		packNames[idx] = game.Name
	}
	return packNames
}
