package search

import (
	"animalshogi/models"
	"animalshogi/tools"
)

func AlphaBetaSearch(pBoards *[]models.Board, playernum int, depth int, alpha int, beta int, reverse int) (models.Move, int) {

	var bestMove models.Move

	if depth == 0 {
		return YomiBetterMove(pBoards, playernum)
	}

	nextmoves := tools.PossibleMove(*pBoards, playernum)

	for _, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := AlphaBetaSearch(nextboard, reversePlayer(playernum), depth-1, beta, alpha, reverse*-1)

		if tmp_alpha*reverse > alpha {
			alpha = tmp_alpha
			bestMove = move
		}

		if tmp_alpha*reverse >= beta {
			break
		}

	}

	return bestMove, alpha

}
