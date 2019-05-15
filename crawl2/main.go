package main

import (
	"fmt"
	"github.com/yongliu1992/go_learn/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} //acquire a token
	list, err := links.Extract(url)
	<-tokens //release token
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	//links.Extract()
	workList := make(chan []string)
	var n int // number of pending sends to work
	//start with the command-line arguments
	n++
	go func() {
		workList <- os.Args[:1]
	}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}

}
