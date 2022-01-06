package tools

import (
	"animalshogi/json"
	"fmt"
	"testing"
)

func TestIsSettle(t *testing.T) {

	// player2が王を取って勝ちになるtestjson
	// testjson := `{"B1":"l2","B4":"e2","B2":"g2","C3":"g1","C4":"e1","D1":"c1","E1":"c2"}`

	// player2 がトライしているtestjson
	// testjson := `{"A4":"l2","B4":"e2","B2":"g2","C3":"g1","C4":"l1","D1":"c1","E1":"c2"}`

	// player2 がトライしようとしているが、次で取られるため負けとなる (player1の勝ちとなる) testjson
	testjson := `{"B4":"l2","A4":"e2","B2":"g2","C3":"g1","C4":"l1","D1":"c1","E1":"c2"}`

	result := json.JSONToBoard(testjson)
	PrintBoard(result)

	boolwin, winner := IsSettle(&result)
	fmt.Printf("IsSettle:%v for player%v\n\n", boolwin, winner)

	/*
		nextmoves := PossibleMove(result, 2) // 考えられる次の手をリストアップ
		for _, move := range nextmoves {
			nextboard := DryrunMove(&result, move)
			boolwin, winner := IsSettle(nextboard)
			fmt.Printf("IsSettle:%v for player%v\n", boolwin, winner)
		}
	*/
}
