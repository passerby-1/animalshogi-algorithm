// Player_num()
// Ismyturn()
// MasuToXY()
// XYToMasu()

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

func XYToMasu(xy models.XY) string {
	var resultbytes []byte = []byte("XY") // 文字列のスライスによって場所を指定しての文字列更新はできないのでbyte列で扱う

	switch xy.X {
	case 0:
		resultbytes[0] = 'A'
	case 1:
		resultbytes[0] = 'B'
	case 2:
		resultbytes[0] = 'C'
	case 3:
		resultbytes[0] = 'D'
	case 4:
		resultbytes[0] = 'E'
	default:
		resultbytes[0] = 'X'
	}

	switch xy.Y {
	case 0:
		resultbytes[1] = '1'
	case 1:
		resultbytes[1] = '2'
	case 2:
		resultbytes[1] = '3'
	case 3:
		resultbytes[1] = '4'
	case 4:
		resultbytes[1] = '5'
	case 5:
		resultbytes[1] = '6'
	default:
		resultbytes[1] = 'Y'
	}

	return string(resultbytes)
}
