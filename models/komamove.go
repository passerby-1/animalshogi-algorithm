package models

func SetupMoves() *map[string]interface{} {
	lion := [...][]int{{0, -1}, {1, -1}, {-1, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {0, 1}}
	gilaffe := [...][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	elephant := [...][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	chick := [...][]int{{0, 1}}
	cock := [...][]int{{1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {1, 0}, {0, -1}} // Oh Cock!
	/*
		length := map[string]int{
			"l": 8,
			"g": 4,
			"e": 4,
			"c": 1,
			"h": 6,
		}
	*/
	Moves := map[string]interface{}{
		"l": lion,
		"g": gilaffe,
		"e": elephant,
		"c": chick,
		"h": cock,
		// "length": length,
	}
	return &Moves
}
