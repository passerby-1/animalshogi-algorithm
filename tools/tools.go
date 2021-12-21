package tools

import (
	"golangtest/models"
	"regexp"
	"strconv"
)

func Player_num(msg string) (int, error) {
	regex := `[0-9]`
	reg := regexp.MustCompile(regex)
	result := reg.FindAllStringSubmatch(msg, -1)

	if result != nil {

		result_int, err := strconv.Atoi(result[0][0])

		if err != nil {
			return 0, err
		}

		return result_int, nil

	} else {
		return 0, nil
	}
}

func Ismyturn(turn int, player int) bool {
	if turn == player {
		return true
	} else {
		return false
	}
}

func MasuToXY(masu string) models.XY {
	var result models.XY

	switch masu[0:1] {
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

	switch masu[1:2] {
	case "1":
		result.X = 0
	case "2":
		result.X = 1
	case "3":
		result.X = 2
	case "4":
		result.X = 3
	case "5":
		result.X = 4
	case "6":
		result.X = 5
	default:
		result.X = -1
	}

	return result

}
