package main

import (
	"sync"
	"time"
)

var a string
var once sync.Once

func main() {
	go doprint()
	go dotwoprint()
	time.Sleep(3 * time.Second)
}

func setup() {
	a = "hello,world"
	print(a)
}

func doprint() {
	once.Do(setup)
}

func dotwoprint() {
	once.Do(setup)
}
