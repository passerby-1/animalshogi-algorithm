package tools

import (
	"fmt"
	"testing"
)

func TestDryrunMove(t *testing.T) {

	testjson := `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}`
	result := JSONToBoard(testjson)
	moves1 := PossibleMove(result, 1)

	for i, move := range moves1 {
		fmt.Printf("\nPlayer1, move:%v\nBefore:", move)
		PrintBoard(result)
		aftermove := DryrunMove(result, move)
		fmt.Printf("After:")
		PrintBoard(aftermove)
		if i <= 2 {
			break
		}
	}

	moves2 := PossibleMove(result, 2)
	for i, move := range moves2 {
		fmt.Printf("\nPlayer2, move:%v\nBefore:\n", move)
		PrintBoard(result)
		aftermove := DryrunMove(result, move)
		fmt.Printf("After:\n")
		PrintBoard(aftermove)
		if i <= 2 {
			break
		}
	}

}
