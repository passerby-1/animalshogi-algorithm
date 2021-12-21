package tools

import (
	"encoding/json"
	"golangtest/models"
)

func UnmarshalJSON(data []byte) (map[string]string, error) {

	var Raw map[string]string
	// fmt.Printf("data ([]byte): %v\n", data)

	err := json.Unmarshal(data, &Raw)
	// fmt.Printf("raw (unmarshaled): %v\n", Raw)

	if err != nil {
		return nil, err
	}

	return Raw, nil
}

func JSONToBoard(json string) []models.Board {

	UnmarshaledJSON, _ := UnmarshalJSON([]byte(json))
	result := []models.Board{}

	for zahyo, koma := range UnmarshaledJSON {
		var tmp models.Board
		zahyoXY := MasuToXY(zahyo) // マス目からXYに変換
		// fmt.Printf("マス目:%v 座標:%v\n", zahyo, zahyoXY)
		tmp.Coordinate = zahyoXY // 座標を代入

		playernum, komatype := komaToPlayerAndType(koma)
		tmp.Player = playernum
		tmp.Type = komatype

		result = append(result, tmp)
	}

	return result

}

func komaToPlayerAndType(koma string) (int, string) {
	komabytes := []byte(koma)
	resultType := string(komabytes[0])
	resultPlayer := int(komabytes[1]) - 48

	return resultPlayer, resultType
}
