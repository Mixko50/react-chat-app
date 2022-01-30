package socket

import (
	"backend/loaders/socket/connect"
	"backend/loaders/socket/event"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

var Users []string

func Init() {

	Server2 := socketio.NewServer(nil)
	Server2.OnConnect("connection", connect.Connect)

	Server2.OnEvent("/", "user:connected", event.UserConnect)
	go Server2.Serve()
	println("I'm here")
	defer Server2.Close()

	http.Handle("/socket.io/", Server2)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
