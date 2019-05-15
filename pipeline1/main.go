package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarer := make(chan int)
	go counter(naturals)
	go squarer2(squarer, naturals)
	printer(squarer)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer2(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(out <-chan int) {
	for v := range out {
		fmt.Println(v)
	}
}
