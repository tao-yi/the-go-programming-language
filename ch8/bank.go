package main

import "sync"

func main() {

}

var mu sync.RWMutex
var balance int = 0

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.RLock() // locks for reading
	defer mu.RUnlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}
