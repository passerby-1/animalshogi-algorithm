// SIGINTのシグナルハンドリングについて
// https://qiita.com/TubAnri/items/019f8d19b91f32c878cf

package main

import (
	"animalshogi/jsontools"
	"animalshogi/search"
	"animalshogi/socket"
	"animalshogi/tools"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {

	var (
		ip    = flag.String("ip", "localhost", "IP address")
		port  = flag.String("port", "4444", "port number")
		depth = flag.Int("depth", 5, "search depth")
	)

	fmt.Println("Client start.")

	flag.Parse()
	address := *ip + ":" + *port
	s, _ := socket.Connect(address)

	go sub(s, *depth) // 並列実行

	// タイマー
	//	reset := make(chan bool)
	//	go timer.CountUp(time.Second*60, reset)

	// 終了処理等
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit // ここより下はSIGINTを受けて実行

	fmt.Printf("\nSIGINT Signal received, ending client.\n")
	socket.Close(s)
	os.Exit(0)
}

func sub(s net.Conn, depth int) { // goroutine(並列実行, Ctrl+Cキャッチする奴と並列実行)

	message, _ := socket.Recieve(s) // 初回のメッセージ受信
	player, _ := tools.Player_num(message)

	fmt.Printf("Player: %v\n", player)
	fmt.Printf("recieved msg: %v", message)

	// var boardjson string

	for {

		message := socket.SendRecieve(s, "turn")
		current_turn, _ := tools.Player_num(message)

		if current_turn == player { // 自分の番だったら

			message := socket.SendRecieve(s, "boardjson") // 盤面を取得
			time.Sleep(time.Second * 3)                   // GUI 上でまだ駒が動いているため sleep

			currentBoards := jsontools.JSONToBoard(message) // []models.Board に変換
			tools.PrintBoard(currentBoards)

			boolwin, winner := tools.IsSettle(&currentBoards)

			if boolwin {
				fmt.Printf("[FINISHED] The winner is Player %v", winner)
				break
			}

			bestMove, bestScore := search.AlphaBetaSearch(&currentBoards, player, depth, -1000, 1000, 1)
			moveString := tools.Move2string(bestMove)

			fmt.Printf("bestMove:%v, bestScore:%v, sendmsg: %v\n", bestMove, bestScore, moveString)

			message = socket.SendRecieve(s, moveString)
			time.Sleep(time.Second * 3)

		}

		time.Sleep(time.Second * 2)

	}

	socket.Close(s)
	os.Exit(0)

}
