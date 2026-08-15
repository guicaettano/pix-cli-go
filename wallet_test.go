package main

import "testing"

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
			"saldo do Guilherme deveria continuar em 100, recebeu %.2f",
			guilherme.Balance,
		)
	}

	if nicole.Balance != 300 {
		t.Errorf(
			"saldo da Nicole deveria continuar em 300, recebeu %.2f",
			nicole.Balance,
		)
	}
}
