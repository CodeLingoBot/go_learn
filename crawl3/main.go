package main

import (
	"fmt"
	"github.com/yongliu1992/go_learn/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	workList := make(chan []string)
	unSeekLinks := make(chan string)
	go func() {
		workList <- os.Args[1:]
	}()
	for i := 0; i < 20; i++ {
		for link := range unseekLinks {
			foundLinks := crawl(link)
			go func() {
				workList <- foundLinks
			}()
		}
	}
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unSeekLinks <- link
			}
		}
	}
}
