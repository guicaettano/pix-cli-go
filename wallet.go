package main

import "fmt"

type Account struct {
	Name    string
	Balance int64
}

func formatMoney(value int64) string {
	return fmt.Sprintf("R$ %d,%02d", value/100, value%100)
}

func CreateAccount(bank map[string]*Account, key string, name string) error {
	if _, exists := bank[key]; exists {
		return fmt.Errorf("conta já existe")
	}

	bank[key] = &Account{
		Name:    name,
		Balance: 0,
	}

	return nil
}

func (c *Account) Deposit(value int64) {
	if value > 0 {
		c.Balance += value
	}
}

func (c *Account) Withdraw(value int64) error {
	if value <= 0 {
		return fmt.Errorf("o valor deve ser maior que zero")
	}

	if value > c.Balance {
		return fmt.Errorf("saldo insuficiente")
	}

	c.Balance -= value
	return nil
}

func (c *Account) Transfer(destination *Account, value int64) error {
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
	fmt.Printf("%s - Saldo: R$ %s\n", c.Name, formatMoney(c.Balance))
}
