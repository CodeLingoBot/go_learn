package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var h bool
var newYork string
var london string

func init() {
	flag.StringVar(&newYork, "newYork", "", "美国")
	flag.StringVar(&london, "london", "", "美国")
	flag.BoolVar(&h, "help", false, "help")
}

//var newYork = flag.String("newYork","","请输入纽约的服务器newyork=xxx:port")

func main() {

	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	//time.Sleep(5*time.Second)
	//for {
	go doWall(london, "london")
	go doWall(newYork, "newYork")
	//}
	time.Sleep(10 * time.Second)

}

func doWall(hostName, subName string) {
	conn, err := net.Dial("tcp", hostName)
	if err != nil {
		fmt.Print(hostName)
		log.Fatal(hostName, err)
	}
	defer conn.Close()
	mustCopy(subName, os.Stdout, conn)
}

func mustCopy(subName string, dst io.Writer, src io.Reader) {
	_, err := fmt.Fprint(dst, subName)
	if err != nil {
		fmt.Print("err", subName)
	}

	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
