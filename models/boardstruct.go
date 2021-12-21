package models

type board struct {
	Squares []Square
}

type Square struct {
	Coordinate XY     // 座標
	Type       string // コマの種類
}

type XY struct {
	X int
	Y int
}
