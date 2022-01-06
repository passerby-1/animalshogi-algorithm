// PossibleMove()
// AppendMove()

package tools

import (
	"animalshogi/models"
)

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
