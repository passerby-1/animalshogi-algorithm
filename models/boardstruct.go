package models

type Board struct {
	Coordinate XY     `json:"coordinate"` // 座標
	Player     int    `json:"player"`     // プレイヤーの番号
	Type       string `json:"type"`       // コマの種類
}

type XY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Move struct {
	Src XY
	Dst XY
}
