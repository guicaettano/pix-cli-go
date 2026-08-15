package main

import "fmt"

func main() {
	bank := make(map[string]*Account)
	historic := []Transaction{}

	bank["gui"] = &Account{Name: "Guilherme", Balance: 500}
	bank["ni"] = &Account{Name: "Nicole", Balance: 300}

	for {
		fmt.Println("\n=== PIX CLI GO ===")
		fmt.Println("1 - Listar contas")
		fmt.Println("2 - Depositar")
		fmt.Println("3 - Sacar")
		fmt.Println("4 - Transferir (Pix)")
		fmt.Println("5 - Histórico")
		fmt.Println("0 - Sair")

		var option int
		fmt.Print("Escolha: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			for key, account := range bank {
				fmt.Printf("%s - %s: R$ %.2f\n", key, account.Name, account.Balance)
			}

		case 2:
			var key string
			var value float64

			fmt.Print("Conta: ")
			fmt.Scan(&key)

			fmt.Print("Valor: ")
			fmt.Scan(&value)

			account, exist := bank[key]
			if !exist {
				fmt.Println("Conta não encontrada")
				continue
			}

			account.Deposit(value)
			fmt.Println("Depósito realizado!")

		case 3:
			var key string
			var value float64

			fmt.Print("Conta: ")
			fmt.Scan(&key)

			fmt.Print("Valor: ")
			fmt.Scan(&value)

			account, exist := bank[key]
			if !exist {
				fmt.Println("Conta não encontrada")
				continue
			}

			if account.Withdraw(value) {
				fmt.Println("Saque realizado!")
			} else {
				fmt.Println("Saldo insuficiente")
			}

		case 4:
			var origin, destination string
			var value float64

			fmt.Print("Origem: ")
			fmt.Scan(&origin)

			fmt.Print("Destino: ")
			fmt.Scan(&destination)

			fmt.Print("Valor: ")
			fmt.Scan(&value)

			accountOrigin, ok1 := bank[origin]
			accountDestination, ok2 := bank[destination]

			if !ok1 || !ok2 {
				fmt.Println("Conta não encontrada")
				continue
			}

			if accountOrigin.Transfer(accountDestination, value) {
				historic = append(historic, Transaction{
					Origin: origin, Destination: destination, Value: value,
				})

				fmt.Println("Pix realizado com sucesso!")
			} else {
				fmt.Println("Saldo insuficiente.")
			}

		case 5:
			if len(historic) == 0 {
				fmt.Println("Nenhuma transação registrada.")
				continue
			}

			for i, t := range historic {
				fmt.Printf("%d. %s -> %s | R$ %.2f\n", i+1, t.Origin, t.Destination, t.Value)
			}

		case 0:
			fmt.Println("Muito obrigado, até breve!")
			return

		default:
			fmt.Println("Opção inválida.")

		}
	}
}
