package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func runServer(conn net.Conn) {

	for {

		//lets read from conn and write to user (logs first)
		data, err := bufio.NewReader(conn).ReadString(byte('\n'))
		if err != nil {
			log.Printf("Server problem with read form connection %v", err)
		}
		fmt.Printf("Lets see the data from client %v \n", data)

		//lets read some data from the console
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Printf("Server problem with read form console %v", err)
		}

		fmt.Printf("I'm sending data to the client %v \n", text)
		conn.Write([]byte(text))
	}

}

func main() {
	log.Println("Hi server is here")
	nc, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}
	conn, err := nc.Accept()
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}

	runServer(conn)
}
