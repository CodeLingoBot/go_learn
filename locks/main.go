package main

import (
	"fmt"
	"sync"
)

var l sync.RWMutex
var a string

func f() {
	a = "hell world"
	l.Unlock()
}
func main() {
	l.Lock()
	go f()
	l.Lock()
	fmt.Println(a)

}
