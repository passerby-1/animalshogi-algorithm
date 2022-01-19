package search

import (
	"animalshogi/models"
	"animalshogi/tools"

	"github.com/pterm/pterm"
)

func AlphaBetaSearch(pBoards *[]models.Board, playernum int, depth int, alpha int, beta int, reverse int, orgdepth int) (models.Move, int) {

	var bestMove models.Move

	if depth == 0 {
		return YomiBetterMove(pBoards, playernum)
		/*
			bias := math.Sqrt(float64(orgdepth)) / float64(depth+1) // 深ければ深いほど、評価が浅くなるようにしたい
			pterm.Printf("depth: %v, bias: %v\n", depth, bias)
			tmpMove, tmp := YomiBetterMove(pBoards, playernum)
			return tmpMove, int(float64(tmp) * bias)
		*/
	}

	nextmoves := tools.PossibleMove(*pBoards, playernum)

	for i, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := AlphaBetaSearch(nextboard, reversePlayer(playernum), depth-1, beta, alpha, reverse*-1, depth)

		//bias := math.Sqrt(float64(orgdepth)) / float64(depth+1) // 深ければ深いほど、評価が浅くなるようにしたい
		//tmp_alpha = int(float64(tmp_alpha) * bias)

		if tmp_alpha*reverse > alpha {
			alpha = tmp_alpha
			bestMove = move
		}

		if tmp_alpha*reverse >= beta {
			break
		}

		pterm.Printf("BestMove in depth: %v, index: %v is %v\n", depth, i, bestMove)
		tools.PrintBoard(*nextboard)

	}

	return bestMove, alpha

}
