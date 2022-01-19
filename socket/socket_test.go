package socket_test

import (
	"animalshogi/socket"
	"animalshogi/tools"
	"fmt"
	"testing"
	"time"
)

func TestSocket(t *testing.T) {
	fmt.Println("Client start.")
	s, _ := socket.Connect("localhost:4444")

	message, _ := socket.Recieve(s) // 初回のメッセージ受信
	player, _ := tools.Player_num(message)

	fmt.Printf("player:%v", player)

	socket.Send(s, "turn")
	message, _ = socket.Recieve(s)
	current_turn, _ := tools.Player_num(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("current turn:%v\n", current_turn)

	socket.Send(s, "boardjson")
	message, _ = socket.Recieve(s)
	fmt.Printf("message after send boardjson: %v", message)

	socket.Close(s)
}
