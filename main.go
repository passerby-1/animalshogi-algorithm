// SIGINTのシグナルハンドリングについて
// https://qiita.com/TubAnri/items/019f8d19b91f32c878cf

package main

import (
	"flag"
	"fmt"
	"golangtest/socket"
	"golangtest/tools"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {

	fmt.Println("Client start.")

	flag.Parse()
	args := flag.Args()
	address := args[0] + ":" + args[1]
	s, _ := socket.Connect(address)

	// s, _ := socket.Connect("localhost:4444")

	// s, err := socket.Connect("10.128.219.201:4444")

	/*
		if err != nil {
			fmt.Errorf("%s", err)
		}
	*/

	go sub(s) // 並列実行

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit // ここより下はSIGINTを受けて実行

	fmt.Printf("\nSIGINT Signal received, ending client.\n")
	socket.Close(s)
	os.Exit(0)
}

func sub(s net.Conn) { // goroutine(並列実行, Ctrl+Cキャッチする奴と並列実行)

	message, _ := socket.Recieve(s) // 初回のメッセージ受信
	player, _ := tools.Player_num(message)

	fmt.Printf("Player: %v\n", player)
	fmt.Printf("recieved msg: %v", message)

	time.Sleep(time.Second * 1)

	for {

		// UnmarshaledJson, _ = tools.UnmarshalJSON([]byte(org_json))

		socket.Send(s, "turn")
		message, _ = socket.Recieve(s)
		current_turn, _ := tools.Player_num(message)

		if current_turn == player { // 自分の番だったら
			socket.Send(s, "boardjson") // 盤面を取得
			message, _ = socket.Recieve(s)
			time.Sleep(time.Second * 3)

			currentBoards := tools.JSONToBoard(message) // []models.Board に変換
			tools.PrintBoard(currentBoards)

			boolwin, winner := tools.IsSettle(&currentBoards)

			if boolwin {
				fmt.Printf("[FINISHED] The winner is Player %v", winner)
				break
			}

			bestMove, bestScore := tools.MiniMax(&currentBoards, player, 5, 1)
			moveString := tools.Move2string(bestMove)

			fmt.Printf("bestMove:%v, bestScore:%v, sendmsg: %v", bestMove, bestScore, moveString)

			socket.Send(s, moveString)
			time.Sleep(time.Second * 3)
		}
		// fmt.Printf("recieved msg: %v", message)
		// fmt.Printf("Current turn: %v\n", current_turn)

		time.Sleep(time.Second * 2)
	}

	socket.Close(s)
	os.Exit(0)

}
