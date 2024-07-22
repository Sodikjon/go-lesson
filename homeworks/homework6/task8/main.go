package main

import "fmt"

func main() {
	var loanSum float64 = 500000.00
	fmt.Printf("Сумма Вашего вклада составлет: %.2f\n", loanSum)
	yearlyLoanPercents(&loanSum)
	yearlyLoanPercents(&loanSum)
	yearlyLoanPercents(&loanSum)
	yearlyLoanPercents(&loanSum)
}

// 8. Начисление процентов на депозит с ежегодной проверкой
// • Начальная сумма депозита равна 500000 рублей.
// • Реализуйте функции для начисления процентов каждый год по ставке 6%.
// • Выводите баланс после каждого начисления.
// • Если баланс становится больше 700000 рублей, выведите сообщение "Достигнут лимит начислений".
func yearlyLoanPercents(loanSum *float64) {
	if *loanSum < 700000.00 {
		*loanSum += *loanSum * 0.06
		fmt.Printf("Вам начислен процент на вклад в размере: %.2f\n", *loanSum)
		fmt.Printf("Сумма вашего вклада после начисления составляет: %.2f\n", *loanSum)

	} else {
		fmt.Println("Достигнут лимит начислений")
	}

}
