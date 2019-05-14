package main

import (
	"bufio"
	"flag"
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
	log.Println("Hi client is here - approach with flags instead of arguments")

	args := os.Args

	sec := flag.Int("timeout", 60, "time out default 60") //!! Pointer
	port := flag.Int("port", 8989, "port to listen default 8989")
	flag.Parse()

	log.Printf("flag ags: %v, %v and args: %v port %v\n", flag.Args(), *sec, args, *port)

	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Printf("Panic! problem with netcat %v", err)
	}

	runClient(conn)
}
