package main

import "fmt"

func main() {
	var loanSum float64 = 200000
	fmt.Printf("Сумма Вашего вклада составлет: %.2f\n", loanSum)
	compoundLoanPercents(&loanSum)
	compoundLoanPercents(&loanSum)
	compoundLoanPercents(&loanSum)
	compoundLoanPercents(&loanSum)
}

// 6. Начисление сложных процентов на вклад
// • Начальная сумма вклада равна 200000 рублей.
// • Реализуйте функции для начисления сложных процентов каждый месяц по ставке 5%.
// • Выводите баланс после каждого начисления.
// • Если баланс становится больше 300000 рублей, выведите сообщение "Достигнут лимит начислений".
func compoundLoanPercents(loanSum *float64) {
	if *loanSum < 300000.00 {
		*loanSum += *loanSum * (0.05 / 12.00)
		fmt.Printf("Вам начислен процент на вклад в размере: %.2f\n", *loanSum*(float64(0.05)/12.00))
		fmt.Printf("Сумма вашего вклада составляет: %.2f\n", *loanSum)

	} else {
		fmt.Println("Достигнут лимит начислений")
	}

}
