package models

func SetupMoves() map[string]interface{} {
	lion := [...][]int{{0, -1}, {1, -1}, {-1, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {0, 1}}
	gilaffe := [...][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	elephant := [...][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	Chick := [...][]int{{0, 1}}
	Cock := [...][]int{{1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {1, 0}, {0, -1}} // Oh Cock!

	Moves := map[string]interface{}{
		"l": lion,
		"g": gilaffe,
		"e": elephant,
		"c": Chick,
		"h": Cock,
	}
	return Moves
}
