package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance = amount + balance
}
func Balance() int {
	return balance
}
func main() {
	fmt.Println("last", goDeposit())
}

func goDeposit() int {
	//var n = sync.WaitGroup{}
	//n.Add(2)
	go func() {

		Deposit(200)
		fmt.Println("= 加200", Balance())
		//n.Done()
	}()
	go func() {
		Deposit(100)
		//	n.Done()
	}()
	//n.Wait()
	time.Sleep(1 * time.Second)
	return balance
}
