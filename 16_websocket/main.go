package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Socket payground")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal("Socket io was not initiated!: ", err)
	} else {
		log.Printf("server %+v\n", server)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("We get new connection with id:", s.ID())
		s.Emit("Hi there")
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	server.OnError("/", func(e error) {
		fmt.Println("Socket error:", e)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there! - just for test")
	})

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("System crashed - ListenAndServe :", err)
	}

}
