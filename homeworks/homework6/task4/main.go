package main

import "fmt"

func main() {
	var loanSum = 100000.00
	fmt.Printf("Сумма Вашего вклада составлет: %.2f\n", loanSum)
	loanPercents(&loanSum)
	loanPercents(&loanSum)
	loanPercents(&loanSum)
	loanPercents(&loanSum)

}

// 4. Учет процентов по вкладам с ежемесячной капитализацией
//   - Начальный вклад равен 100000 рублей.
//   - Реализуйте функции для ежемесячного начисления процентов по ставке 5% годовых.
//   - Капитализируйте проценты ежемесячно и выводите сумму вклада после каждого начисления.
//   - Если сумма вклада превышает 150000 рублей, выведите сообщение "Достигнут лимит вклада".
func loanPercents(loanSum *float64) {
	if *loanSum < 150000.00 {
		*loanSum += *loanSum * (0.05 / 12.00)
		fmt.Printf("Вам начислен процент на вклад в размере: %.2f\n", *loanSum*(float64(0.05)/12.00))
		fmt.Printf("Сумма вашего вклада составляет: %.2f\n", *loanSum)

	} else {
		fmt.Println("Достигнут лимит вклада")
	}

}
