// このファイルに含まれる関数
// 雑多になっているため要整理
//
// Player_num()
// Ismyturn()
// MasuToXY()
// XYToMasu()
// PrintBoard()
// TypeToKanji()
// Remove()
// player2arrow()
// QueryBoard()
// Move2string()

package tools

import (
	"animalshogi/models"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func Player_num(msg string) (int, error) {
	regex := `[0-9]`
	reg := regexp.MustCompile(regex)
	result := reg.FindAllStringSubmatch(msg, -1)

	if result != nil {

		result_int, err := strconv.Atoi(result[0][0])

		if err != nil {
			return 0, err
		}

		return result_int, nil

	} else {
		return 0, nil
	}
}

func Ismyturn(turn int, player int) bool {
	if turn == player {
		return true
	} else {
		return false
	}
}

func MasuToXY(masu string) models.XY {
	var result models.XY
	masubytes := []byte(masu)

	switch string(masubytes[0]) {
	case "A":
		result.X = 0
	case "B":
		result.X = 1
	case "C":
		result.X = 2
	case "D":
		result.X = 3 // Player1 の持ち駒
	case "E":
		result.X = 4 // Player2 の持ち駒
	default:
		result.X = -1
	}

	switch string(masubytes[1]) {
	case "1":
		result.Y = 0
	case "2":
		result.Y = 1
	case "3":
		result.Y = 2
	case "4":
		result.Y = 3
	case "5":
		result.Y = 4
	case "6":
		result.Y = 5
	default:
		result.Y = -1
	}

	return result

}

func XYToMasu(xy models.XY) string {
	var resultbytes []byte = []byte("XY") // 文字列のスライスによって場所を指定しての文字列更新はできないのでbyte列で扱う

	switch xy.X {
	case 0:
		resultbytes[0] = 'A'
	case 1:
		resultbytes[0] = 'B'
	case 2:
		resultbytes[0] = 'C'
	case 3:
		resultbytes[0] = 'D'
	case 4:
		resultbytes[0] = 'E'
	default:
		resultbytes[0] = 'X'
	}

	switch xy.Y {
	case 0:
		resultbytes[1] = '1'
	case 1:
		resultbytes[1] = '2'
	case 2:
		resultbytes[1] = '3'
	case 3:
		resultbytes[1] = '4'
	case 4:
		resultbytes[1] = '5'
	case 5:
		resultbytes[1] = '6'
	default:
		resultbytes[1] = 'Y'
	}

	return string(resultbytes)
}

func PrintBoard(boards []models.Board) {

	// 持ち駒の表示のため、無駄なループを回している気がするので後で修正する
	// TODO: 持ち駒がない場合, 盤面上に駒が無い場合を追加

	// 持ち駒かそれ以外かで処理が変わるので、先にこれらを別の配列に仕分けしてしまう (一瞬しか持たないデータなので少しくらいメモリ食っても良い)

	var MochiKomaTmp []models.Board
	var KomaTmp []models.Board

	for _, board := range boards {

		if (board.Coordinate.X == 4) || (board.Coordinate.X == 3) { // 持ち駒であれば (E or D)
			MochiKomaTmp = append(MochiKomaTmp, board)
		} else { // 盤面上の駒であれば ((E or D) 以外)
			KomaTmp = append(KomaTmp, board)
		}
	}

	// 持ち駒だけでのソート (E, D の順, 1-6 の順)
	// 先に 1-6 で並べた後、E, D に安定ソートすれば良いかな?
	fmt.Printf("\n[現在の盤面]\nPlayer2の持ち駒:\n")

	count := 0
	if MochiKomaTmp != nil { // 持ち駒が空でなかった場合にプリント

		sort.Slice(MochiKomaTmp, func(i, j int) bool {
			return MochiKomaTmp[i].Coordinate.Y < MochiKomaTmp[j].Coordinate.Y // 1-6 の順番
		})

		sort.SliceStable(MochiKomaTmp, func(i, j int) bool {
			return MochiKomaTmp[i].Coordinate.X > MochiKomaTmp[j].Coordinate.X // E(4), D(3) の順番
		})

		// まず Player2 の持ち駒を表示する (E)
		// fmt.Printf("\n[現在の盤面]\nPlayer2の持ち駒:\n")
		for i, board := range MochiKomaTmp {
			if board.Coordinate.X == 3 { // D に入ったら break, E が空だった場合も即座にこっちで break するので例外追加は必要なし
				count = i
				break
			}

			fmt.Printf("%s ", TypeToKanji(board.Type))

		}
	}

	fmt.Printf("\n----------\n")

	// 通常の盤面部のソート

	sort.Slice(KomaTmp, func(i, j int) bool { // X でソート
		return KomaTmp[i].Coordinate.X < KomaTmp[j].Coordinate.X
	})

	sort.SliceStable(KomaTmp, func(i, j int) bool { // Y でソート, 先の X でのソートの順番を保つため安定ソート
		return KomaTmp[i].Coordinate.Y < KomaTmp[j].Coordinate.Y
	})

	// 通常の盤面の表示
	var xyNow models.XY

	for i := 0; i < 4; i++ { // 行を動かすループ (Y)
		if i != 0 {
			fmt.Printf("\n")
		}

		for j := 0; j < 3; j++ { // 列を動かすループ (X)

			xyNow.X = j
			xyNow.Y = i

			if len(KomaTmp) != 0 { // KomaTmp[0] への nil アクセス防止, KomaTmp != nil では空の構造体の判定ができない

				tmp := KomaTmp[0] // 先頭を取り出し

				if tmp.Coordinate == xyNow {

					fmt.Printf("%s%s", TypeToKanji(tmp.Type), player2arrow(tmp))
					KomaTmp = Remove(KomaTmp, 0) // 今の先頭(0)を削除, 1番目が次の0番目となる

				} else {
					fmt.Printf("□  ")
				}

			} else {
				if xyNow.X == 2 && xyNow.Y == 3 { // 右下の角に駒がなかった場合 KomaTmp の長さが0になり、表示されないのでその例外
					fmt.Printf("□  ")
				}
			}

		}
	}

	// 最後に Player1 の持ち駒を表示する
	fmt.Printf("\n----------\nPlayer1の持ち駒:\n")

	if len(MochiKomaTmp) != count {
		for j := count; j < len(MochiKomaTmp); j++ {
			fmt.Printf("%s ", TypeToKanji(MochiKomaTmp[j].Type))
		}
	}

	fmt.Printf("\n\n")

}

func TypeToKanji(komatype string) string {

	switch komatype {
	case "l":
		return "王"
	case "g":
		return "飛"
	case "c":
		return "歩"
	case "h":
		return "と"
	case "e":
		return "角"
	default:
		return "？"
	}
}

// 渡した配列から n 番目の要素を remove した配列を返す関数
func Remove(slice []models.Board, s int) []models.Board {
	return append(slice[:s], slice[s+1:]...)
}

// Player 番号を見て print 用の矢印を返す
func player2arrow(tmp models.Board) string {
	if tmp.Player == 1 {
		return "↑"
	} else if tmp.Player == 2 {
		return "↓"
	} else {
		return "?"
	}
}

// []models.Board からとある models.XY 座標を持つ models.Board を探してリターンする関数
// 愚直だが多分仕方がない…(reflect使って汎用性の高いcontains関数を作ると逆に遅いハズ)
func QueryBoard(boards []models.Board, xy models.XY) (bool, models.Board) {
	for _, board := range boards {
		if board.Coordinate == xy {
			return true, board
		}
	}

	var nilBoard models.Board
	nilBoard.Coordinate = xy
	nilBoard.Player = -1
	nilBoard.Type = "not found"

	return false, nilBoard
}

// models.Move からどうぶつしょうぎサーバへ送るコマンド形式への変換
func Move2string(move models.Move) string {
	var result string
	result = "mv " + XYToMasu(move.Src) + " " + XYToMasu(move.Dst)
	return result
}
