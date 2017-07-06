// Copyright 2017 Ken Miura
package bank_test

import (
	"fmt"
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ex09/ex01/bank"
)

func TestBank1(t *testing.T) {
	bank.Withdraw(bank.Balance())       // テストの前処理
	defer bank.Withdraw(bank.Balance()) // テストの後処理

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

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBank2(t *testing.T) {
	bank.Withdraw(bank.Balance())       // テストの前処理
	defer bank.Withdraw(bank.Balance()) // テストの後処理
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{} // bank.Depositが終わるまでbank.Withdrawを待たせる
		done <- struct{}{}
	}()

	// Bob
	go func() {
		<-done // bank.Depositが終わるまで待つ
		bank.Withdraw(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	bank.Withdraw(bank.Balance()) // テストの前処理
	var tests = []struct {
		input1    int
		input2    int
		expected1 int
		expected2 bool
	}{
		{200, 100, 100, true},
		{200, 200, 0, true},
		{0, 200, 0, false},
	}

	for _, tt := range tests {
		bank.Deposit(tt.input1)

		ok := bank.Withdraw(tt.input2)

		if !(bank.Balance() == tt.expected1 && ok == tt.expected2) {
			t.Fatalf("%d and %t are expected but were %d and %t", tt.expected1, tt.expected2, bank.Balance(), ok)
		}
		bank.Withdraw(bank.Balance()) // テストの後処理
	}
}
