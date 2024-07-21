package main

import "fmt"

func main() {
	var deposit int = 100000
	refillDeposit(100000, &deposit)
	refillDeposit(50000, &deposit)
	refillDeposit(150000, &deposit)
	refillDeposit(90000, &deposit)
	refillDeposit(50000, &deposit)
	refillDeposit(50000, &deposit)
	refillDeposit(50000, &deposit)

}

//1. Учет накопительных счетов с ежемесячным пополнением
//• Начальный баланс накопительного счета равен 100000 рублей.
//• Реализуйте функции для пополнения счета каждый месяц на фиксированную сумму.
//• Выводите баланс после каждого пополнения.
//• Если баланс становится больше 500000 рублей, выведите сообщение "Достигнут лимит накоплений".

func refillDeposit(refillSum int, deposit *int) {
	if *deposit <= 500000 {
		*deposit += refillSum
		fmt.Printf("Ваш баланс составляет: %d рублей\n", *deposit)
	} else {
		fmt.Println("Достигнут лимит накоплений")
	}
}
