// Copyright 2017 Ken Miura
package bank

import "errors"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawArg = make(chan WithdrawArg)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

type WithdrawArg struct {
	Amount int
	Ok     chan bool
}

func Withdraw(arg WithdrawArg) error {
	if arg.Ok == nil {
		return errors.New("WithdrawArg must not be nil")
	}
	withdrawArg <- arg
	return nil
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case arg := <-withdrawArg:
			balance -= arg.Amount
			if balance < 0 {
				balance += arg.Amount
				arg.Ok <- false
			} else {
				arg.Ok <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
