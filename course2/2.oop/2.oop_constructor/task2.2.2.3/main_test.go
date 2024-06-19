package main

import (
	"testing"
)

type depListTest struct {
	expected float64
	received float64
}

func TestDeposit(t *testing.T) {
	saveAcc := &SavingsAccount{}
	depTest := []depListTest{
		{0, 0},
		{1000.1, 1000.1},
		{1600.1, 600},
	}
	for _, val := range depTest {
		saveAcc.Deposit(val.received)
		if saveAcc.money != val.expected {
			t.Errorf("expected = %v, got = %v", val.expected, saveAcc.money)
		}
	}
}
func TestWithdraw(t *testing.T) {
	saveAcc := &SavingsAccount{}
	depTest := []depListTest{
		{1100, -9},
		{1100, 21100.9},
		{500, 600},
	}
	saveAcc.Deposit(1100)
	for _, val := range depTest {
		err := saveAcc.Withdraw(val.received)
		if (val.received > 1100 && err == nil) || (val.received <= 0 && err == nil) {
			t.Errorf("Error expected: %v, but got = nil. Value received = %v, on deposit = %v", err, val.received, saveAcc.money)
		} else if saveAcc.money != val.expected {
			t.Errorf("expected = %v, got = %v", val.expected, saveAcc.money)
		}
	}
}
func TestBalance(t *testing.T) {
	saveAcc := &SavingsAccount{}
	depTest := []depListTest{
		{0, -400},
		{1000.1, 1000.1},
		{2600.2, 1600.1},
	}
	for _, val := range depTest {
		saveAcc.Deposit(val.received)
		if saveAcc.Balance() != val.expected {
			t.Errorf("expected = %v, got = %v", val.expected, saveAcc.Balance())
		}
	}
}

func TestWithName(t *testing.T) {
	expecteds := []string{"admin", "item2", ""}
	testOrder := &Customer{}
	for _, expected := range expecteds {
		WithName(expected)(testOrder)
		if expected != testOrder.Name {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Name)
		}
	}
}

func TestWithAccount(t *testing.T) {
	savings := &SavingsAccount{}
	expecteds := []Account{savings, &SavingsAccount{}}
	testOrder := &Customer{}
	for _, expected := range expecteds {
		WithAccount(expected)(testOrder)
		if expected != testOrder.Account {
			t.Errorf("expected = %v, got = %v", expected, testOrder.Name)
		}
	}
}

func TestNewCustomer(t *testing.T) {
	idTest := 21
	nameTest := "testuser"
	AccountTest := &SavingsAccount{}
	expected := NewCustomer(idTest,
		WithName(nameTest),
		WithAccount(AccountTest))
	testO := &Customer{idTest, nameTest, AccountTest}
	if expected.Id != testO.Id {
		t.Errorf("expected = %v, got = %v", expected.Id, testO.Id)
	}
	if expected.Name != testO.Name {
		t.Errorf("expected = %v, got = %v", expected.Name, testO.Name)
	}
	if expected.Account != testO.Account {
		t.Errorf("expected = %v, got = %v", expected.Account, testO.Account)
	}

}
