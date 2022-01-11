package tools

import (
	"animalshogi/models"
	"animalshogi/socket"
	"net"
)

// 決着がついているかを判定する関数
// 入力は盤面の []models.Board 返り値は bool と player番号のint
func IsSettle(pBoards *[]models.Board) (bool, int) {
	// キャッチの判定
	// まず、どちらかの王が消えていれば、残っている方の勝ち
	// ありえないが、両方の王が消えるのも例外処理

	// トライの判定
	// 自分の王が相手エリアの端 (player1の場合 Y=0, player2の場合 Y=3) にいるか?
	// その場合、今の盤面の次に考えられる指し手は何か? その中に、自分の王が消えているものはあるか? 無ければトライである

	var kings []models.Board
	for _, board := range *pBoards {
		if board.Type == "l" && board.Coordinate.X <= 2 { // 盤面に出ている王だったら
			kings = append(kings, board)
		}
	}

	switch len(kings) {
	case 0:
		return false, -1 // 例外
	case 1:
		return true, kings[0].Player // 王が1人しかいない = どちらかが勝利しているとき
	}

	for _, king := range kings {
		switch king.Player {
		case 1:
			if king.Coordinate.Y == 0 { // player1 の王が相手エリアの端にいたら
				nextmoves := PossibleMove(*pBoards, 2) // 考えられる次の手をリストアップ
				for _, move := range nextmoves {
					nextboard := DryrunMove(pBoards, move) // 考えられる全ての次の手について
					if isCatch(nextboard, 1) {             // player1 はキャッチされてる?
						return true, 2 // 次の手でキャッチ出来る場合相手の勝ちになる
					}
				}
				return true, 1
			}
		case 2:
			if king.Coordinate.Y == 3 { // player2 の王が相手エリアの端にいたら
				nextmoves := PossibleMove(*pBoards, 1)
				for _, move := range nextmoves {
					nextboard := DryrunMove(pBoards, move)
					// PrintBoard(*nextboard)
					if isCatch(nextboard, 2) { // 次の手でplayer1 はキャッチされてる?
						return true, 1 // 次の手でキャッチ出来る場合相手の勝ちになる
					}
				}
				return true, 2
			}
		}
	}

	return false, -1

}

func isCatch(pBoards *[]models.Board, player int) bool {

	var kings []models.Board
	for _, board := range *pBoards {
		if board.Type == "l" && board.Coordinate.X <= 2 { // 盤面に出ている王だったら
			kings = append(kings, board)
		}
	}

	switch len(kings) {
	case 0:
		return false // 例外

	case 1:
		if kings[0].Player == player { // 残っている王が自分であれば取られていない
			return false
		} else {
			return true // 残っている王が自分ではないため、相手が自分を取っている
		}

	default:
		return false
	}
}

// 並列実行用, turnChan に turn を流す
func TurnCheck(s net.Conn, turnChan chan int) {

	for {
		message := socket.SendRecieve(s, "turn")
		current_turn, _ := Player_num(message)

		switch current_turn {
		case 1:
			turnChan <- 1
			// turnChangeChan <- true
		case 2:
			turnChan <- 2
			// turnChangeChan <- true
		}
	}
}
