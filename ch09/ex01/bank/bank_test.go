// Copyright 2017 Ken Miura
package bank_test

import (
	"fmt"
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch09/ex01/bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	// Alice
	aliceOk := make(chan bool)
	go func() {
		bank.Withdraw(bank.WithdrawArg{50, aliceOk})
		done <- struct{}{}
	}()
	if !<-aliceOk {
		t.Error("Alice: Withdraw.Ok = false")
	}

	// Bob
	bobOk := make(chan bool)
	go func() {
		bank.Withdraw(bank.WithdrawArg{300, bobOk})
		done <- struct{}{}
	}()
	if <-bobOk {
		t.Error("Bob: Withdraw.Ok = true")
	}

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 250; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
