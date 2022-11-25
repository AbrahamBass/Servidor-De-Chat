package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWriter(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the serve your name %v", clientName)
	messages <- fmt.Sprintf("New Client is here name %v", clientName)

	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%v : %v\n", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%v said goodbye", clientName)
}

func MessageWriter(conn net.Conn, message <-chan string) {
	for message := range message {
		fmt.Fprintln(conn, message)
	}
}

func BroadCast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go BroadCast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}
