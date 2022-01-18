package search

import (
	"animalshogi/jsontools"
	"animalshogi/models"
	"animalshogi/tools"
	"encoding/json"
	"io/ioutil"
	"reflect"
	"sort"

	"github.com/mitchellh/hashstructure/v2"
	"github.com/pterm/pterm"
)

func HashingBoards(Boards []models.Board) uint64 {

	// ハッシュ化前にソートして、駒の順番の影響を受けないようにする
	sort.Slice(Boards, func(i, j int) bool { // X でソート
		return Boards[i].Coordinate.X < Boards[j].Coordinate.X
	})

	sort.SliceStable(Boards, func(i, j int) bool { // Y でソート, 先の X でのソートの順番を保つため安定ソート
		return Boards[i].Coordinate.Y < Boards[j].Coordinate.Y
	})

	Hash, _ := hashstructure.Hash(Boards, hashstructure.FormatV2, nil)

	// pterm.Printf("Hash:%v\n", Hash)

	return Hash
}

func FullSearchJSON() {

	initialBoard := `{"A1":"g2","B1":"l2","C1":"e2","B2":"c2","B3":"c1","A4":"e1","B4":"l1","C4":"g1"}`
	currentBoards := jsontools.JSONToBoard(initialBoard)
	var Hashes []HashTree

	FullSearchRecursive(&Hashes, &currentBoards, 1, 0)

	for _, hash := range Hashes {
		for _, prev := range hash.Prev {
			_, pHashExists := QueryHashTree(&Hashes, prev)
			pHashExists.Next = append(pHashExists.Next, prev)

		}
	}

	file, _ := json.MarshalIndent(Hashes, "", " ")
	_ = ioutil.WriteFile("FullSearch.json", file, 0644)

}

func FullSearchRecursive(Hashes *[]HashTree, pBoards *[]models.Board, playernum int, currentdepth int) {

	currentHash := HashingBoards(*pBoards) // 現在の盤面状態のハッシュ値
	nextmoves := tools.PossibleMove(*pBoards, playernum)

	for i, move := range nextmoves { // 考えられる全ての手について

		nextboard := tools.DryrunMove(pBoards, move)
		nextHash := HashingBoards(*nextboard)
		winbool, winner := tools.IsSettle(nextboard)
		var tmpHashTree HashTree

		if winbool {
			tmpHashTree.WinLose = winner
		} else {
			tmpHashTree.WinLose = 0
		}

		tmpHashTree.Hash = nextHash
		tmpHashTree.Boards = *nextboard
		tmpHashTree.Prev = append(tmpHashTree.Prev, currentHash)
		tmpHashTree.Next = nil // Next は後から Prev を見て埋める

		// pterm.Printf("tmpHashTree1:%v\nMove:%v\n\n", tmpHashTree, move)
		continuebool := AppendHash(Hashes, tmpHashTree, *nextboard, currentHash)

		if continuebool || winbool {
			pterm.Printf("Skipped:%v\n", tmpHashTree.Hash)
			continue
		}

		pterm.Printf("CurrentDepth:%v, i:%v\n", currentdepth, i)
		pterm.Printf("tmpHashTree:%v\n\n", tmpHashTree)

		/*
			if currentdepth == 100 {
				tools.PrintBoard(tmpHashTree.Boards)
				panic("stop")
			}
		*/

		FullSearchRecursive(Hashes, nextboard, ReversePlayer(playernum), currentdepth+1)

	}

	if currentdepth%10000 == 0 {
		pterm.Printf("Current Depth is %v\n", currentdepth)
	}

}

func AppendHash(Hashes *[]HashTree, newHash HashTree, board []models.Board, prev uint64) bool {

	exists, pHashExists := QueryHashTree(Hashes, newHash.Hash)

	if exists {
		equalbool := reflect.DeepEqual(pHashExists.Boards, newHash.Boards)
		pterm.Printf("pHashExists.Boards: %v\n", pHashExists.Boards)
		pterm.Printf("newHash.Boards    : %v\n", newHash.Boards)
		pterm.Printf("equalbool:%v\n", equalbool)

		if equalbool {
			pHashExists.Prev = append(pHashExists.Prev, prev)
			pterm.Printf("[DEBUG] 既出の盤面のハッシュ: %v\n", newHash.Hash)
		} else {
			pterm.Printf("\n[ERROR] ハッシュ衝突!!!!!!!!!\n") // 一旦ハッシュ衝突は考えない, どうぶつしょうぎの盤面で衝突が出たら処理を考える
			pterm.Printf("衝突したハッシュ値: %v\n衝突した盤面: %v\n同じハッシュ値の盤面: %v\n同じハッシュ値の盤面のハッシュ値: %v\n", newHash.Hash, newHash.Boards, pHashExists.Boards, pHashExists.Hash)
			panic("ハッシュ衝突")
		}

		return true // 既存のものだった

	} else {
		pterm.Printf("[DEBUG] 新しい盤面のハッシュ: %v\n", newHash.Hash)
		*Hashes = append(*Hashes, newHash)
		return false // 新規のものだった
	}

	// 盤面状態を逆さにしたら、キャッシュの量を半分に減らせるが prev の扱いが面倒になる
	// rboard := ReverseBoard(board)
	// rnewHash := HashingBoards(*rboard)

}

func QueryHashTree(Hashes *[]HashTree, queryhash uint64) (bool, *HashTree) {

	for _, hash := range *Hashes {
		if hash.Hash == queryhash {
			return true, &hash
		}
	}

	var nilHash HashTree
	nilHash.Hash = 0
	nilHash.Boards = nil
	nilHash.WinLose = 0
	nilHash.Next = nil
	nilHash.Prev = nil

	return false, &nilHash
}

/*
func ReverseBoard(boards []models.Board) *[]models.Board {

	var KomaTmp []models.Board

	for _, board := range boards {

		if (board.Coordinate.X == 4) || (board.Coordinate.X == 3) { // 持ち駒であれば (E or D)
			board.Coordinate.X = reverseDE(board.Coordinate.X) // player をひっくり返す
			KomaTmp = append(KomaTmp, board)

		} else { // 盤面上の駒であれば ((E or D) 以外)
			board.Coordinate.X = reverseX(board.Coordinate.X)
			board.Coordinate.Y = reverseY(board.Coordinate.Y)
			KomaTmp = append(KomaTmp, board)
		}
	}

	// 盤面のソート

	sort.Slice(KomaTmp, func(i, j int) bool { // X でソート
		return KomaTmp[i].Coordinate.X < KomaTmp[j].Coordinate.X
	})

	sort.SliceStable(KomaTmp, func(i, j int) bool { // Y でソート, 先の X でのソートの順番を保つため安定ソート
		return KomaTmp[i].Coordinate.Y < KomaTmp[j].Coordinate.Y
	})

	return &KomaTmp
}
*/

func reverseDE(x int) int {
	if x == 3 {
		return 4
	} else {
		return 3
	}
}

func reverseX(x int) int {
	switch x {
	case 0:
		return 2
	case 1:
		return 1
	case 2:
		return 0
	default:
		return -1
	}
}

func reverseY(y int) int {
	switch y {
	case 0:
		return 3
	case 1:
		return 2
	case 2:
		return 1
	case 3:
		return 0
	default:
		return -1
	}
}
