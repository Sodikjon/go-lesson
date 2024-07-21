package main

import "fmt"

func main() {
	var balance = 500000
	fmt.Printf("balance is %d\n", balance)
	bankTransfer(90000, &balance)
	bankTransfer(110000, &balance)
}

// 3. Учет операций по банковским переводам с лимитом суммы
// • Начальный баланс счета равен 500000 рублей.
// • Реализуйте функции для выполнения банковских переводов.
// • Если сумма перевода превышает 100000 рублей, выведите сообщение "Сумма перевода превышает лимит".
// • Выводите остаток на счете после каждого перевода.
func bankTransfer(transferSum int, balance *int) {
	if transferSum > 100000 {
		fmt.Println("Сумма перевода превышает лимит")
	} else {
		fmt.Println("Вы перевели: ", transferSum)
		*balance -= transferSum
		fmt.Println("Ваш баланс составляет: ", *balance)
	}
}
