package search

import (
	"animalshogi/models"
	"animalshogi/tools"
	"fmt"
)

// 一手読み, 勝つ手があればそれを指す
func Yomi1Move(pBoards *[]models.Board, playernum int) models.Move {
	var result models.Move
	nextmoves := tools.PossibleMove(*pBoards, playernum) // 考えられる次の手をリストアップ

	for _, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)
		result = move

		boolwin, winner := tools.IsSettle(nextboard)
		if boolwin && winner == playernum { // 自分の勝ちだったら
			break
		}
	}

	return result

}

// 一手読み, より良い手を返すようにするための枠組みとなる, 最善手の評価値も返す
func YomiBetterMove(pBoards *[]models.Board, playernum int) (models.Move, int) {

	bestScore := -100000
	var result models.Move

	nextmoves := tools.PossibleMove(*pBoards, playernum)
	for _, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)
		score := staticScoring(nextboard, playernum)

		if score > bestScore {
			bestScore = score
			result = move
			// fmt.Printf("Score:%v\n", bestScore)
		}

	}

	return result, bestScore

}

func staticScoring(pBoards *[]models.Board, playernum int) int {

	boolwin, winner := tools.IsSettle(pBoards)

	if boolwin && (winner == playernum) { // 自分の勝ちだったら
		return 100000 // 勝ちなので最高点

	} else {

		// 仮なので、勝ちではなかった場合 持ち駒の数×100 を返す

		count := 0

		for _, board := range *pBoards {
			if board.Player == playernum {
				count++
			}
		}

		return count * 100

	}
}

func MiniMax(pBoards *[]models.Board, playernum int, depth int, orgDepth int, reverse int) (models.Move, int) {

	var bestMove models.Move
	var alpha int

	if depth == 0 {
		return YomiBetterMove(pBoards, playernum)
	}

	// 評価値の最大値を記録するための変数
	alpha = -1000 * reverse

	nextmoves := tools.PossibleMove(*pBoards, playernum)
	for _, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := MiniMax(nextboard, reversePlayer(playernum), depth-1, orgDepth, reverse*-1)

		if tmp_alpha*reverse > alpha {
			alpha = tmp_alpha
			bestMove = move
			fmt.Printf("current alpha:%v depth:%v\n", alpha, depth)
		}

		if tmp_alpha == 100000 {
			break
		}
	}

	/*
		fmt.Printf("PLAYER %v (reverse:%v) depth: %v\nbestMove:%v, alpha:%v\nBEFORE:\n", playernum, reverse, depth, bestMove, alpha)
		tools.PrintBoard(*pBoards)
		fmt.Printf("AFTER:\n")
		tools.PrintBoard(*tools.DryrunMove(pBoards, bestMove))
		fmt.Printf("\n")
	*/

	return bestMove, alpha

}

func reversePlayer(playernum int) int {
	if playernum == 1 {
		return 2
	} else {
		return 1
	}
}
