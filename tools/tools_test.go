package tools

import (
	"animalshogi/jsontools"
	"animalshogi/models"
	"fmt"
	"testing"
)

func TestMasuToXY(t *testing.T) {

	result := MasuToXY("C1")

	var correctResult models.XY
	correctResult.X = 2
	correctResult.Y = 0

	if result != correctResult {
		t.Errorf("Error")
	}

	t.Logf("Success")
}

func TestXYToMasu(t *testing.T) {

	var challenge models.XY
	challenge.X = 0
	challenge.Y = 0

	result := XYToMasu(challenge)

	if result != "A1" {
		t.Errorf("Error")
	}

	t.Logf("Success")
}

func TestPrintBoard(t *testing.T) {
	testjson := `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}` // 8駒分が出ればOK
	testboard := jsontools.JSONToBoard(testjson)

	PrintBoard(testboard)

}

func TestMove2string(t *testing.T) {
	var move models.Move
	move.Src.X = 1
	move.Src.Y = 1
	move.Dst.X = 2
	move.Dst.Y = 2

	fmt.Printf("Move:%v\nstring:%v\n", move, Move2string(move))
}
