package tools

import (
	"animalshogi/jsontools"
	"animalshogi/models"
	"fmt"
	"testing"
)

func TestPossibleMove(t *testing.T) {

	testjson := `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}`
	result := jsontools.JSONToBoard(testjson)
	PrintBoard(result)
	moves1 := PossibleMove(result, 1)
	fmt.Printf("Player1 が出すことが出来る手\n%v\n", moves1)
	moves2 := PossibleMove(result, 2)
	fmt.Printf("Player2 が出すことが出来る手\n%v\n", moves2)
}

func TestAppendMove(t *testing.T) {
	var ReturnMoves []models.Move
	var data1 models.Move
	var xyNow models.XY
	xyNow.X = 0
	xyNow.Y = 0
	data1.Src = xyNow
	data1.Dst.X = 1
	data1.Dst.Y = 1
	ReturnMoves = append(ReturnMoves, data1)
	fmt.Printf("Before: %v\n", ReturnMoves)

	var xyChecking models.XY
	xyChecking.X = -1
	xyChecking.Y = -1

	AppendMove(&ReturnMoves, &xyNow, &xyChecking)

	fmt.Printf("After: %v\n", ReturnMoves)
}
