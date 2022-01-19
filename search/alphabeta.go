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
	}

	nextmoves := tools.PossibleMove(*pBoards, playernum)

	for i, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := AlphaBetaSearch(nextboard, reversePlayer(playernum), depth-1, beta, alpha, reverse*-1, depth)

		if tmp_alpha*reverse > alpha {
			alpha = tmp_alpha
			bestMove = move
		}

		if alpha*reverse >= beta {
			break
		}

		pterm.Printf("BestMove in depth: %v, index: %v is %v\n", depth, i, bestMove)
		tools.PrintBoard(*nextboard)

	}

	return bestMove, alpha

}

/*
func AlphaBetaSearch1(pBoards *[]models.Board, depth int, alpha int, beta int, orgdepth int) (models.Move, int) {

	playernum := 1
	var bestMove models.Move

	if depth == 0 {
		return YomiBetterMove(pBoards, playernum)
	}

	nextmoves := tools.PossibleMove(*pBoards, playernum)

	for i, move := range nextmoves {
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := AlphaBetaSearch2(nextboard, depth-1, beta, alpha, depth)

		if tmp_alpha > alpha {
			alpha = tmp_alpha
			bestMove = move
			// pterm.Printf("if 1 bestmove: %v", bestMove)
		}

		if alpha >= beta {
			break
		}

		pterm.Printf("BestMove in depth: %v, index: %v is %v\n", depth, i, bestMove)
		tools.PrintBoard(*nextboard)

	}

	return bestMove, alpha

}

func AlphaBetaSearch2(pBoards *[]models.Board, depth int, alpha int, beta int, orgdepth int) (models.Move, int) {

	playernum := 2
	var bestMove models.Move

	if depth == 0 {
		return YomiBetterMove(pBoards, playernum)
	}

	nextmoves := tools.PossibleMove(*pBoards, playernum)
	// pterm.Printf("nextmoves: %v", nextmoves)

	for i, move := range nextmoves {
		pterm.Printf("Checking: %v", move)
		nextboard := tools.DryrunMove(pBoards, move)

		_, tmp_alpha := AlphaBetaSearch1(nextboard, depth-1, beta, alpha, depth)

		//bias := math.Sqrt(float64(orgdepth)) / float64(depth+1) // 深ければ深いほど、評価が浅くなるようにしたい
		//tmp_alpha = int(float64(tmp_alpha) * bias)

		if tmp_alpha < alpha {
			alpha = tmp_alpha
			bestMove = move
			// pterm.Printf("if 2 bestmove: %v", bestMove)
		}

		if alpha <= beta {
			break
		}

		pterm.Printf("BestMove in depth: %v, index: %v is %v\n", depth, i, bestMove)
		// tools.PrintBoard(*nextboard)
	}

	return bestMove, alpha

}
*/
