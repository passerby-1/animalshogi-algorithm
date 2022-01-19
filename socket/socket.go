// ソケット通信について (receive部分は変更)
// https://ren.nosuke.me/blog/202006/20200615/
//
// Connect()
// Close()
// Send()
// Recieve()
// SendRecieve()

package socket

import (
	"bufio"
	"net"
	"time"

	"github.com/pterm/pterm"
)

func Connect(addr string) (net.Conn, error) {

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Close(conn net.Conn) error {
	err := conn.Close()
	return err
}

func Send(conn net.Conn, msg string) error {
	msg = msg + "\n"

	// _, err := fmt.Fprintf(conn, msg)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.Write([]byte(msg))

	return err
}

func Recieve(conn net.Conn) (string, error) {
	status, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		return "", err
	}

	return status, nil
}

func SendRecieve(conn net.Conn, msg string) string {
	Send(conn, msg)
	recieved, _ := Recieve(conn)
	pterm.Printf("recieved msg (send:%v): %v", msg, recieved)
	return recieved
}
