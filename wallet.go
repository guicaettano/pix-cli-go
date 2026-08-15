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

func (c *Account) Withdraw(value float64) error {
	if value <= 0 {
		return fmt.Errorf("o valor deve ser maior que zero")
	}

	if value > c.Balance {
		return fmt.Errorf("saldo insuficiente")
	}

	c.Balance -= value
	return nil
}

func (c *Account) Transfer(destination *Account, value float64) error {
	if c == destination {
		return fmt.Errorf("a conta de origem e destino são iguais")
	}

	if err := c.Withdraw(value); err != nil {
		return err
	}

	destination.Deposit(value)

	return nil
}

func (c Account) Show() {
	fmt.Printf("%s - Saldo: R$ %.2f\n", c.Name, c.Balance)
}
