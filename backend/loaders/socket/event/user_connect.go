package event

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func UserConnect(s socketio.Conn, username string) {
	println(username)
	var info = UserConnectType{
		Username: username,
		Online:   1,
	}
	s.Emit("user:connected", info)
	log.Println("on connection")
}
