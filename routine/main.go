package main

import "fmt"

func main() {
	var done = make(chan bool)
	var data = make(chan int)
	go product(data)
	go consumer(data, done)
	<-done
}

func product(data chan<- int) {
	for i := 0; i < 8; i++ {
		data <- i
	}
	close(data)
}

func consumer(nums <-chan int, done chan<- bool) {
	for x := range nums {
		fmt.Println(x)
	}
	done <- true
}
