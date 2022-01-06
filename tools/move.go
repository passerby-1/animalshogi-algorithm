package tools

import (
	"animalshogi/models"
)

// 手を指す
// func ExecuteMove()

// その手を指したときの盤面状況を返す
func DryrunMove(Boards *[]models.Board, move models.Move) *[]models.Board {
	// TODO: 1. 駒の行き先をquery
	//       2. Boards からソースを削除、行き先を追加
	//		 3. queryが存在した場合、それを適切なplayerの持ち駒に追加

	// var newBoards []models.Board
	newBoards := make([]models.Board, len(*Boards))
	copy(newBoards, *Boards)

	dstkomaExists, _ := QueryBoard(newBoards, move.Dst)

	if dstkomaExists { // 宛先に相手の駒があった場合 (自分の持ち駒になる場合)

		for i, board := range newBoards {
			if board.Coordinate == move.Dst {
				// player1,2 -> D列かE列かの判定, D, E列の一番早い奴を見つけてその次に追加
				// このときの player は board.Player==move.Dst となる player ではない方 (取る方)
				var player int

				if board.Player == 1 {
					player = 2
				} else {
					player = 1
				}

				AppendMochiKoma(&newBoards, player, &newBoards[i])

			}
		}
	}

	// 今の駒の座標を、新しい座標に書き換える
	for i, board := range newBoards {
		if board.Coordinate == move.Src {
			newBoards[i].Coordinate = move.Dst
		}
	}

	// fmt.Printf("INTERNAL:\n")
	// PrintBoard(newBoards)

	return &newBoards

}

// 持ち駒を追加する, koma は playernum の player が取った駒についての Board
func AppendMochiKoma(pBoards *[]models.Board, playernum int, koma *models.Board) {
	maxY := -1
	for _, board := range *pBoards {
		if board.Coordinate.X == playernum+2 { // D か E の適切なそれについて
			if board.Coordinate.Y > maxY {
				maxY = board.Coordinate.Y
			}
		}
	}
	// これで今の player の持ち駒の数が求まった, これの次のインデックスに追加することになる
	// 今の相手の駒の、player を自分に書き換え、座標を持ち駒に書き換える
	koma.Player = playernum
	koma.Coordinate.X = playernum + 2
	koma.Coordinate.Y = maxY + 1

	if koma.Type == "h" { // ニワトリ (と金) を取ったときひよこ (歩) に戻す
		koma.Type = "c"
	}

	// *pBoards = append(*pBoards, newMochiKoma)
}

/*
func RemoveSpecifiedBoard(slice []models.Board, xy models.XY) []models.Board {
	for i, board := range slice {
		if board.Coordinate == xy {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return nil
}
*/

// models.Move からどうぶつしょうぎサーバへ送るコマンド形式への変換
func Move2string(move models.Move) string {
	var result string
	result = "mv " + XYToMasu(move.Src) + " " + XYToMasu(move.Dst)
	return result
}
