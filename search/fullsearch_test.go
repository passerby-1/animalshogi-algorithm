package search_test

import (
	"animalshogi/jsontools"
	"animalshogi/search"
	"testing"

	"github.com/pterm/pterm"
)

func TestFullSearchJSON(t *testing.T) {
	search.FullSearchJSON()
}

func TestHashingBoards(t *testing.T) {

	initialBoard := `{"A1":"g2","B1":"l2","C1":"e2","B2":"c2","B3":"c1","A4":"e1","B4":"l1","C4":"g1"}`
	Boards := jsontools.JSONToBoard(initialBoard)
	hash := search.HashingBoards(Boards)
	pterm.Printf("Hash:%v\n", hash)

}
