// SIGINTのシグナルハンドリングについて
// https://qiita.com/TubAnri/items/019f8d19b91f32c878cf

package main

import (
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
	s, err := socket.Connect("localhost:4444")
	// s, err := socket.Connect("10.128.219.201:4444")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	go sub(s) // 並列実行

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit // ここより下はSIGINTを受けて実行

	fmt.Printf("\nSIGINT Signal received, ending client.\n")
	socket.Close(s)
	os.Exit(0)
}

func sub(s net.Conn) { // goroutine(並列実行, Ctrl+Cキャッチする奴と並列実行)

	message, _ := socket.Recieve(s)
	player, _ := tools.Player_num(message)

	fmt.Printf("Player: %v\n", player)
	fmt.Printf("recieved msg: %v", message)

	time.Sleep(time.Second * 1)

	for {
		// counter := 0

		socket.Send(s, "turn")
		message, _ = socket.Recieve(s)
		current_turn, _ := tools.Player_num(message)
		fmt.Printf("recieved msg: %v", message)
		fmt.Printf("Current turn: %v\n", current_turn)

		/*
			if tools.Ismyturn(current_turn, player) {
				fmt.Printf("My turn!")
				sendmsg := "mv " + move[counter]
				socket.Send(s, sendmsg)
				counter++

				if counter == 1 {
					break
				}
			}
		*/

		time.Sleep(time.Second * 1)
	}
}
