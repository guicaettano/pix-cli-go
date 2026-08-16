package main

import "testing"

func TestCreateAccount(t *testing.T) {
	bank := make(map[string]*Account)

	err := CreateAccount(bank, "gui", "Guilherme")

	if err != nil {
		t.Errorf("não deveria ter ocorrido erro: %v", err)
	}

	account, exists := bank["gui"]

	if !exists {
		t.Error("a conta deveria existir")
	}

	if account.Name != "Guilherme" {
		t.Errorf("nome esperado Guilherme, recebeu %s", account.Name)
	}
}

func TestTransferInsufficientBalance(t *testing.T) {
	guilherme := Account{
		Name:    "Guilherme",
		Balance: 100,
	}

	nicole := Account{
		Name:    "Nicole",
		Balance: 300,
	}

	err := guilherme.Transfer(&nicole, 200)

	if err == nil {
		t.Error("a transferência deveria ter falhado")
	}

	if guilherme.Balance != 100 {
		t.Errorf(
			"saldo do Guilherme deveria continuar em 100, recebeu %s",
			formatMoney(guilherme.Balance),
		)
	}

	if nicole.Balance != 300 {
		t.Errorf(
			"saldo da Nicole deveria continuar em 300, recebeu %s",
			formatMoney(nicole.Balance),
		)
	}
}
