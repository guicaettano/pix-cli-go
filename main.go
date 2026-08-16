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
		fmt.Println("6 - Criar conta")
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

			err := account.Withdraw(value)

			if err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Println("Saque realizado!")
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

			err := accountOrigin.Transfer(accountDestination, value)

			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}

			historic = append(historic, Transaction{
				Origin:      origin,
				Destination: destination,
				Value:       value,
			})

			fmt.Println("Pix realizado com sucesso!")

		case 5:
			if len(historic) == 0 {
				fmt.Println("Nenhuma transação registrada.")
				continue
			}

			for i, t := range historic {
				fmt.Printf("%d. %s -> %s | R$ %.2f\n", i+1, t.Origin, t.Destination, t.Value)
			}

		case 6:
			var key string
			var name string

			fmt.Print("Nome: ")
			fmt.Scan(&name)

			fmt.Print("Chave Pix: ")
			fmt.Scan(&key)

			err := CreateAccount(bank, key, name)

			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}

			fmt.Println("Conta criada com sucesso!")

		case 0:
			fmt.Println("Muito obrigado, até breve!")
			return

		default:
			fmt.Println("Opção inválida.")

		}
	}
}
