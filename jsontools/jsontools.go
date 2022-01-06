// UnmarshalJSON()
// JSONToBoard()
// komaToPlayerAndType()

package jsontools

import (
	"animalshogi/models"
	"encoding/json"
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

func MasuToXY(masu string) models.XY {
	var result models.XY
	masubytes := []byte(masu)

	switch string(masubytes[0]) {
	case "A":
		result.X = 0
	case "B":
		result.X = 1
	case "C":
		result.X = 2
	case "D":
		result.X = 3 // Player1 の持ち駒
	case "E":
		result.X = 4 // Player2 の持ち駒
	default:
		result.X = -1
	}

	switch string(masubytes[1]) {
	case "1":
		result.Y = 0
	case "2":
		result.Y = 1
	case "3":
		result.Y = 2
	case "4":
		result.Y = 3
	case "5":
		result.Y = 4
	case "6":
		result.Y = 5
	default:
		result.Y = -1
	}

	return result

}
