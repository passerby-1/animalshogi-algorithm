package tools

import (
	"golangtest/models"
	"testing"
)

func TestMasuToXY(t *testing.T) {

	result := MasuToXY("A1")

	var correctResult models.XY
	correctResult.X = 0
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
