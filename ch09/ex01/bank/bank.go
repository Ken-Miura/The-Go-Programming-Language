// Copyright 2017 Ken Miura
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan int)
var ok = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraw:
			balance -= amount
			if balance < 0 {
				balance += amount
				ok <- false
			} else {
				ok <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
