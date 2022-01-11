// SIGINTのシグナルハンドリングについて
// https://qiita.com/TubAnri/items/019f8d19b91f32c878cf

package main

import (
	"animalshogi/jsontools"
	"animalshogi/search"
	"animalshogi/socket"
	"animalshogi/timer"
	"animalshogi/tools"
	"flag"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/pterm/pterm"
)

func main() {

	pterm.Println("Client start.")

	// 実行時の引数の処理等, ソケット通信接続

	var (
		ip    = flag.String("ip", "localhost", "IP address")
		port  = flag.String("port", "4444", "port number")
		depth = flag.Int("depth", 5, "search depth")
	)

	flag.Parse()
	address := *ip + ":" + *port
	s, _ := socket.Connect(address)

	// ターンのチェック用
	turnChan := make(chan int)

	// タイマー用
	timeChan := time.NewTimer(time.Second * 55)
	tickChan := time.NewTimer(time.Second * 1)
	resetChan := make(chan bool)
	resetCompreteChan := make(chan bool)

	// プレイヤー番号
	// sub() に入れるとTurnCheckとメッセージを送るタイミングが被るので仮に外に出している

	message, _ := socket.Recieve(s) // 初回のメッセージ受信 (You are PlayerN)
	player, _ := tools.Player_num(message)

	pterm.Printf("Player: %v\n", player)
	pterm.Printf("recieved msg: %v", message)

	// 並列実行
	go sub(s, player, *depth, turnChan, resetChan, resetCompreteChan)
	go timer.Timer(timeChan, tickChan, resetChan, resetCompreteChan, turnChan)
	go tools.TurnCheck(s, turnChan)
	go func() { // リセットが完了する度に、resetChan を false へ戻すため
		for {
			select {
			case <-resetCompreteChan:
				resetChan <- false
			default:
			}
		}
	}()

	// 終了処理等
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit // ここより下はSIGINTを受けて実行

	pterm.Printf("\nSIGINT Signal received, ending client.\n")
	socket.Close(s)
	os.Exit(0)
}

func sub(s net.Conn, player int, depth int, turnChan chan int, resetChan chan bool, resetCompreteChan chan bool) {
	/*
		go func() { // リセットが完了する度に、resetChan を false へ戻すため
			for {
				select {
				case <-resetCompreteChan:
					resetChan <- false
				default:
				}
			}
		}()
	*/

	for {
		select {
		case currentTurn := <-turnChan:

			pterm.Printf("[DEBUG] Current turn: %v\n", currentTurn)

			if currentTurn == player {

				// resetChan <- true // タイマーリセット
				// resetChan が何か悪さをしている様子, コメントを外すとデッドロックがどこかに発生する

				message := socket.SendRecieve(s, "boardjson")   // 盤面を取得
				currentBoards := jsontools.JSONToBoard(message) // []models.Board に変換
				tools.PrintBoard(currentBoards)

				boolwin, winner := tools.IsSettle(&currentBoards)

				if boolwin {
					pterm.Printf("[FINISHED] The winner is Player %v\n", winner)
					break
				}

				bestMove, bestScore := search.AlphaBetaSearch(&currentBoards, player, depth, -1000, 1000, 1)
				moveString := tools.Move2string(bestMove)

				pterm.Printf("bestMove:%v, bestScore:%v, sendmsg: %v\n", bestMove, bestScore, moveString)

				message = socket.SendRecieve(s, moveString) // moved

				// time.Sleep(time.Second * 2)

				// resetChan <- true // 自分が打ち終わると共に相手の番になるのでタイマーリセット
			}

		}

	}

}
