package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		//handle err
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//handle err

		}
		go handleCon(conn)
	}

}

func handleCon(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	//NOTE: ignoring potential errors from input.Err() ------关闭资源啊啊啊
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
