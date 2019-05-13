package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func runClient(conn net.Conn) {

	for {

		//lets read some data from the console
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Printf("Client problem with read form console %v", err)
		}
		fmt.Printf("I'm sending data to the client %v \n", text)
		conn.Write([]byte(text))

		//lets read from conn and write to user (logs first)
		data, err := bufio.NewReader(conn).ReadString(byte('\n'))
		if err != nil {
			log.Printf("Client problem with read form connection %v", err)
		}
		fmt.Printf("Lets see the data from server %v \n", data)
	}

}

func main() {
	log.Println("Hi client is here")

	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}

	runClient(conn)
}
