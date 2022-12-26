package main

import (
	"Lab6_Parallel_Programming/packages/bankApi"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Lab6. Паралельне програмування")

	bank := *bankApi.BankDefaultConstructor()

	ch := make(chan struct{}, 1)
	go func() {
		ch <- struct{}{}
	}()

	for true {
		var menu int

		fmt.Println("\nОберіть пункт меню:")
		fmt.Println("1 - Створити банк")
		fmt.Println("2 - Додати клієнта для кредитів")
		fmt.Println("3 - Додати клієнта для депозитів")
		fmt.Println("4 - Вивести інформацію про клієнта за прізвищем")
		fmt.Println("5 - Вивести інформацію про клієнта за номером акаунту")
		fmt.Println("6 - Вивести інформацію про банк та усіх клієнтів")
		fmt.Println("Інший символ - Завершити роботу програми")

		fmt.Print("Введіть обраний пункт меню: ")
		fmt.Fscan(os.Stdin, &menu)

		switch menu {
		case 1:
			var name string
			var money float64

			fmt.Print("\nВведіть назву банку: ")
			fmt.Fscan(os.Stdin, &name)

			if name == "" {
				name = "default"
			}

			for true {
				fmt.Print("Введіть кількість грошей на рахунках банку: ")
				fmt.Fscan(os.Stdin, &money)

				if money > 0 {
					break
				} else {
					fmt.Println("Сума не може бути менше або рівною нулю")
				}
			}

			bank = bankApi.CreateNewBank(name, money)
		case 2, 3:
			var name string
			var surname string
			var accountNumber string

			fmt.Print("\nВведіть ім'я користувача: ")
			fmt.Fscan(os.Stdin, &name)

			if name == "" {
				name = "default"
			}

			fmt.Print("Введіть прізвище користувача: ")
			fmt.Fscan(os.Stdin, &surname)

			if surname == "" {
				surname = "default"
			}

			fmt.Print("Введіть номер акаунту користувача: ")
			fmt.Fscan(os.Stdin, &accountNumber)

			if accountNumber == "" {
				accountNumber = "default"
			}

			if menu == 2 {
				client := bankApi.CreateNewClient(name, surname, accountNumber, 1000, -1000)

				go func() {
					<-ch
					bank.AddClient(&client)
					ch <- struct{}{}
					go bankApi.Credit(&client, ch)
				}()
			} else {
				client := bankApi.CreateNewClient(name, surname, accountNumber, 1000, 0)

				go func() {
					<-ch
					bank.AddClient(&client)
					ch <- struct{}{}
					go bankApi.Deposit(&client, ch)
				}()
			}
		case 4:
			var surname string

			fmt.Print("\nВведіть прізвище користувача: ")
			fmt.Fscan(os.Stdin, &surname)

			go func() {
				<-ch
				bank.PrintClientsBySurname(surname)
				ch <- struct{}{}
			}()
		case 5:
			var accountNumber string

			fmt.Print("\nВведіть номер акаунту користувача: ")
			fmt.Fscan(os.Stdin, &accountNumber)

			go func() {
				<-ch
				bank.PrintClientByAccountNumber(accountNumber)
				ch <- struct{}{}
			}()
		case 6:
			go func() {
				<-ch
				bank.PrintInfo()
				ch <- struct{}{}
			}()
		default:
			fmt.Println("\nРоботу програми завершено")
			return
		}
	}
}
