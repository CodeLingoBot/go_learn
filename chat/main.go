package main

import (
	"bufio"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err.Error())
	}
	go broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(1, err.Error())
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are" + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Println(conn, msg)
	}
}
