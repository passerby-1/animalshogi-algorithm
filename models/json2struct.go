package models

type board struct {
	Squares []Square
}

type Square struct {
	Coordinate string // 座標
	Type       string // コマの種類
}
