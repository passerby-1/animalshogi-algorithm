package tools

import (
	"fmt"
	"golangtest/models"
)

func PossibleMove(slice []models.Board, playernum int) []models.Move {
	Moves := models.SetupMoves() // 毎回やるのアレなのでポインタ渡しとかに

	var xyNow models.XY

	for i := 0; i < 4; i++ { // 行を動かすループ (Y)
		for j := 0; j < 3; j++ { // 列を動かすループ (X)
			xyNow.X = j
			xyNow.Y = i
			// TODO: 駒があるかどうかのif文
			// TODO: 駒があれば駒の種類を取り出し
			// TODO: その駒の種類の可能な動きをMovesから取り出し
			// TODO: その駒の動き方についてループを回し、それぞれについて、
			//		  移動先は枠の範囲内かを確認、OKであれば次へ
			//			1. 自分の駒があればダメ
			//			2. 他人の駒 or 空きマスであれば追加
			//		追加するのは、models.Moveのsrcに今見ている場所のXYを追加
			//		models.Moveのdstにその行き先のマスのXYを追加
			//		これを[]Move に追加して返す
			
			}

		}
	}
}
