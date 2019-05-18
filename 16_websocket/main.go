package main

import (
	"fmt"
	"log"
	"net/http"

	socket "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Socket payground")

	server, err := socket.NewServer(nil)
	if err != nil {
		log.Fatal("Socket io was not initiated!: ", err)
	}

	server.OnConnect("/", func(s socket.Conn) error {
		s.SetContext("")
		log.Println("We get new root connection with id:", s.ID())
		s.Emit("Hi there")
		return nil
	})
	server.OnEvent("/chat", "reply", func(s socket.Conn, msg string) {
		fmt.Println("Chat message - replay:", msg)
		s.Emit("reply", "have "+msg)
	})

	//first is context (namespace), second is message
	server.OnEvent("/room", "reply", func(s socket.Conn, msg string) {
		fmt.Println("room notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socket.Conn, msg string) string {
		log.Printf("Chat msg from socket event %s", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "msg", func(s socket.Conn, msg string) string {
		log.Printf("root message %s", msg)
		s.Emit("welcome", msg)
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socket.Conn) string {

		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnDisconnect("/", func(s socket.Conn, msg string) {
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
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal("System crashed - ListenAndServe :", err)
	}

}
