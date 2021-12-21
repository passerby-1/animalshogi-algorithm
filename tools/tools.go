package tools

import (
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
