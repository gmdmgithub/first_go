package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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
	args := os.Args
	port := "8989"
	if len(args) > 0 {
		_, err := strconv.Atoi(args[1])
		if err == nil {
			port = args[1]
		}
	}
	nc, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}
	log.Printf("Server port is %s", port)
	conn, err := nc.Accept()
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}

	runServer(conn)
}
