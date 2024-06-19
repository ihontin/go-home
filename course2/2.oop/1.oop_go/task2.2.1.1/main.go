package main

import (
	"errors"
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

// Deposit - вносить деньги на счет
// Withdraw - снимать деньги со счета
// Balance - получать текущий баланс

// CurrentAccount текущий счет
type CurrentAccount struct {
	money float64
}

func newCurrentAccount(sum ...float64) *CurrentAccount {
	account := &CurrentAccount{}
	if len(sum) > 0 {
		account.money = sum[0]
	}
	return account
}
func (s *CurrentAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("ошибка, не верная сумма пополнения счета")
	}
	s.money += amount
	return nil
}
func (s *CurrentAccount) Withdraw(amount float64) error {
	if amount > s.money || amount <= 0 {
		return errors.New("ошибка, не верная сумма списания")
	}
	s.money -= amount
	return nil
}
func (s *CurrentAccount) Balance() float64 {
	return s.money
}

// SavingsAccount сберегательный счет
type SavingsAccount struct {
	money float64
}

func newSavingsAccount() *CurrentAccount {
	return &CurrentAccount{}
}
func (s *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("ошибка, не верная сумма пополнения счета")
	}
	s.money += amount
	return nil
}
func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount > s.money || s.money < 500 || amount <= 0 {
		return errors.New("ошибка, не верная сумма списания")
	}
	s.money -= amount
	return nil
}

func (s *SavingsAccount) Balance() float64 {
	return s.money
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}
	ProcessAccount(c)
	ProcessAccount(s)
}
