package main

import "fmt"

func main() {
	var ch = make(chan string, 3)
	go func() {
		ch <- "hello"
	}()

	go func() {
		ch <- "world"
	}()

	go func() {
		ch <- "天下无双"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
