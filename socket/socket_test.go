package socket

import (
	"animalshogi/tools"
	"fmt"
	"testing"
	"time"
)

func TestSocket(t *testing.T) {
	fmt.Println("Client start.")
	s, _ := Connect("localhost:4444")

	message, _ := Recieve(s) // 初回のメッセージ受信
	player, _ := tools.Player_num(message)

	fmt.Printf("player:%v", player)

	Send(s, "turn")
	message, _ = Recieve(s)
	current_turn, _ := tools.Player_num(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("current turn:%v\n", current_turn)

	Send(s, "boardjson")
	message, _ = Recieve(s)
	fmt.Printf("message after send boardjson: %v", message)

	Close(s)
}
