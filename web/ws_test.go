package web

import (
	"io"
)

func EchoServer(ws *websocket.Conn) {
	fmt.Print("ws debug output")
	fmt.Println(ws)
	io.Copy(ws, ws)
}
