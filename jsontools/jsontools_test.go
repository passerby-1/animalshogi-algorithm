package jsontools

import (
	"animalshogi/models"
	"fmt"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	testjson := `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}`
	correctResult := map[string]string{
		"B1": "l2",
		"B2": "g2",
		"B4": "l1",
		"C1": "e2",
		"C3": "g1",
		"C4": "e1",
		"D1": "c1",
		"E1": "c2",
	}

	result, err := UnmarshalJSON([]byte(testjson))
	// fmt.Printf("%v", result)

	if err != nil {
		t.Errorf("Error:%v", err)
	}

	if len(correctResult) != len(result) {
		t.Errorf("Error")
	}

	fmt.Printf("\n")
	for zahyo, koma := range result {
		fmt.Printf("座標:%v, 駒:%v\n", zahyo, koma)

		if koma != correctResult[zahyo] {
			t.Errorf("Wrong pair")
		}
	}

	t.Logf("Unmarshal Success.")

}

func TestJSONToBoard(t *testing.T) {

	testjson := `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}`
	result := JSONToBoard(testjson)
	fmt.Printf("\nResult:\n%v", result)

	// test未完成
}

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
