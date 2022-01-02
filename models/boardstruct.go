package models

type Board struct {
	Coordinate XY     // 座標
	Player     int    // プレイヤーの番号
	Type       string // コマの種類
}

type XY struct {
	X int
	Y int
}

type Move struct {
	Src XY
	Dst XY
}
