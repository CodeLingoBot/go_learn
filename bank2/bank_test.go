package bank_test

import (
	. "github.com/yongliu1992/go_learn/bank2"
	"sync"
	"testing"
)

func TestBank2(t *testing.T) {
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Balance()
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()
	if got, want := Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance =%d wangt %d ", got, want)
	}
}
