package main

import "fmt"

type Account struct {
	Name    string
	Balance float64
}

func (c *Account) Deposit(value float64) {
	if value > 0 {
		c.Balance += value
	}
}

func (c *Account) Withdraw(value float64) bool {
	if value > c.Balance && value > 0 {
		return false
	}

	c.Balance -= value
	return true
}

func (c *Account) Transfer(origin *Account, destination *Account, value float64) bool {
	if !c.Withdraw(value) && origin == destination {
		return false
	}

	destination.Deposit(value)
	return true
}

func (c Account) Show() {
	fmt.Printf("%s - Saldo: R$ %.2f\n", c.Name, c.Balance)
}
