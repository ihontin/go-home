package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

// SavingsAccount текущий счет
type SavingsAccount struct {
	money float64
	mutex sync.RWMutex
}

func (s *SavingsAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("ошибка, не верная сумма пополнения счета")
		return
	}
	s.mutex.Lock()
	s.money += amount
	s.mutex.Unlock()
}
func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount > s.money || amount <= 0 || s.money < 1000 {
		return errors.New("ошибка, не верная сумма списания")
	}
	s.mutex.Lock()
	s.money -= amount
	s.mutex.Unlock()
	return nil
}
func (s *SavingsAccount) Balance() float64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.money
}

type CustOptions func(*Customer)

type Customer struct {
	Id      int
	Name    string
	Account Account
}

func NewCustomer(id int, options ...CustOptions) *Customer {
	outStruct := &Customer{
		Id: id, Name: "",
	}
	for _, option := range options {
		option(outStruct)
	}
	return outStruct
}

func WithName(s string) CustOptions {
	return func(c *Customer) {
		c.Name = s
	}
}
func WithAccount(a Account) CustOptions {
	return func(c *Customer) {
		c.Account = a
	}
}

//func main() {
//	savings := &SavingsAccount{}
//	savings.Deposit(1000)
//
//	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))
//
//	err := customer.Account.Withdraw(100)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
//}
