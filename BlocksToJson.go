package ui

import "encoding/json"

func BlocksToJSON(blocks []Block) (string, error) {
	blocksMap := []blockJsonObject{}

	for _, block := range blocks {
		blocksMap = append(blocksMap, block.ToJsonObject())
	}

	blocksJson, err := json.Marshal(blocksMap)

	return string(blocksJson), err
}
