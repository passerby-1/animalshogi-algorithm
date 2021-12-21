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

	// var org_json string = `{"B1":"l2","C1":"e2","B2":"g2","C3":"g1","B4":"l1","C4":"e1","D1":"c1","E1":"c2"}`

	for {

		// UnmarshaledJson, _ = tools.UnmarshalJSON([]byte(org_json))

		socket.Send(s, "turn")
		message, _ = socket.Recieve(s)
		current_turn, _ := tools.Player_num(message)
		fmt.Printf("recieved msg: %v", message)
		fmt.Printf("Current turn: %v\n", current_turn)

		time.Sleep(time.Second * 1)
	}
}
