// DryrunMove()
// AppendMochiKoma()
// Move2String()
// PossibleMove()
// AppendMove()

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

// 考えられるパターンすべてを返す
func PossibleMove(Boards []models.Board, playernum int) []models.Move {
	Moves := *models.SetupMoves()
	var ReturnMoves []models.Move
	var xyChecking models.XY
	for _, board := range Boards {
		if (board.Player == playernum) && (board.Coordinate.X <= 2) { // 相手の駒と持ち駒を除外
			xyNow := board.Coordinate
			komaType := board.Type

			// komaMove := Moves[komaType].([...][]int) とはできない…
			// また、komaMove が [8][]int かもしれないし [1][]int かもしれない という状態にはできない (同じ処理であっても全てについてその処理を書く必要がある, golangにジェネリスクは無い…)
			// golang にジェネリスクが無いと思ってたけど増えたんだっけ?
			// ジェネリスクを使うように書き直す (ただしこれは見た目上の問題)
			// 外の関数に追い出そうとしても、結局 interface{} で受け取ることになるため、改めて[8][]int か [6][]int か...の判定を外の関数でも行わなければならず、二重となる

			switch Moves[komaType].(type) {
			case [8][]int:
				komaMove := Moves[komaType].([8][]int)
				for _, currentMove := range komaMove {
					xdelta := currentMove[0]
					ydelta := currentMove[1]
					xyChecking = xyNow
					xyChecking.X = xyChecking.X + xdelta
					xyChecking.Y = xyChecking.Y + ydelta

					if (xyChecking.X >= 0 && xyChecking.X <= 2) && (xyChecking.Y >= 0 && xyChecking.Y <= 3) {
						komaExists, BoardChecking := QueryBoard(Boards, xyChecking)

						if komaExists { // 動かす先に駒が存在する場合
							if BoardChecking.Player != playernum { //  他人の駒である場合取れる (自プレイヤーである場合は取れないので置けない)
								AppendMove(&ReturnMoves, &xyNow, &xyChecking)
							}
						} else { // 動かす先に駒が存在しない場合 (置ける)
							AppendMove(&ReturnMoves, &xyNow, &xyChecking)
						}
					}
				}
			case [6][]int:
				komaMove := Moves[komaType].([6][]int)
				for _, currentMove := range komaMove {
					xdelta := currentMove[0]
					ydelta := currentMove[1]
					xyChecking = xyNow
					xyChecking.X = xyChecking.X + xdelta
					xyChecking.Y = xyChecking.Y + ydelta

					if (xyChecking.X >= 0 && xyChecking.X <= 2) && (xyChecking.Y >= 0 && xyChecking.Y <= 3) {
						komaExists, BoardChecking := QueryBoard(Boards, xyChecking)

						if komaExists { // 動かす先に駒が存在する場合
							if BoardChecking.Player != playernum { //  他人の駒である場合取れる (自プレイヤーである場合は取れないので置けない)
								AppendMove(&ReturnMoves, &xyNow, &xyChecking)
							}
						} else { // 動かす先に駒が存在しない場合 (置ける)
							AppendMove(&ReturnMoves, &xyNow, &xyChecking)
						}
					}
				}
			case [4][]int:
				komaMove := Moves[komaType].([4][]int)
				for _, currentMove := range komaMove {
					xdelta := currentMove[0]
					ydelta := currentMove[1]
					xyChecking = xyNow
					xyChecking.X = xyChecking.X + xdelta
					xyChecking.Y = xyChecking.Y + ydelta

					if (xyChecking.X >= 0 && xyChecking.X <= 2) && (xyChecking.Y >= 0 && xyChecking.Y <= 3) {
						komaExists, BoardChecking := QueryBoard(Boards, xyChecking)

						if komaExists { // 動かす先に駒が存在する場合
							if BoardChecking.Player != playernum { //  他人の駒である場合取れる (自プレイヤーである場合は取れないので置けない)
								AppendMove(&ReturnMoves, &xyNow, &xyChecking)
							}
						} else { // 動かす先に駒が存在しない場合 (置ける)
							AppendMove(&ReturnMoves, &xyNow, &xyChecking)
						}
					}
				}
			case [1][]int:
				komaMove := Moves[komaType].([1][]int)
				for _, currentMove := range komaMove {
					xdelta := currentMove[0]
					ydelta := currentMove[1]
					xyChecking = xyNow
					xyChecking.X = xyChecking.X + xdelta
					xyChecking.Y = xyChecking.Y + ydelta

					if (xyChecking.X >= 0 && xyChecking.X <= 2) && (xyChecking.Y >= 0 && xyChecking.Y <= 3) {
						komaExists, BoardChecking := QueryBoard(Boards, xyChecking)

						if komaExists { // 動かす先に駒が存在する場合
							if BoardChecking.Player != playernum { //  他人の駒である場合取れる (自プレイヤーである場合は取れないので置けない)
								AppendMove(&ReturnMoves, &xyNow, &xyChecking)
							}
						} else { // 動かす先に駒が存在しない場合 (置ける)
							AppendMove(&ReturnMoves, &xyNow, &xyChecking)
						}
					}
				}
			}
		} else if (board.Player == playernum) && (board.Coordinate.X >= 3) { // 自分の持ち駒について
			// 全てのマスについて駒が置かれているかquery, 無ければ置くことが出来る (二歩等も無いので)
			for i := 0; i < 4; i++ { // 行を動かすループ (Y)
				for j := 0; j < 3; j++ { // 列を動かすループ (X)
					xyChecking.X = j
					xyChecking.Y = i
					komaExists, _ := QueryBoard(Boards, xyChecking)
					if !komaExists {
						AppendMove(&ReturnMoves, &board.Coordinate, &xyChecking)
					}
				}
			}
		}
	}

	return ReturnMoves

}

func AppendMove(pMoves *[]models.Move, pxyNow *models.XY, pxyChecking *models.XY) {
	var Move models.Move
	Move.Src = *pxyNow
	Move.Dst.X = pxyChecking.X
	Move.Dst.Y = pxyChecking.Y
	*pMoves = append(*pMoves, Move)
}
