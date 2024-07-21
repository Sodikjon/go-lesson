package main

import "fmt"

func main() {
	var loanSum = 3000000
	fmt.Println("Cумма кредита составляет: ", loanSum)
	loanRepayment(1500000, &loanSum)
	loanRepayment(1000000, &loanSum)
	loanRepayment(150000, &loanSum)
	loanRepayment(150000, &loanSum)

}

// 2. Рассчет выплат по кредиту с фиксированной ставкой
// • Начальная сумма кредита равна 3000000 рублей.
// • Реализуйте функции для ежемесячного расчета выплат по кредиту с фиксированной процентной ставкой 10%.
// • Выводите остаток по кредиту после каждой выплаты.
// • Если остаток по кредиту становится меньше 500000 рублей, выведите сообщение "Почти погашен кредит".
func loanRepayment(repaymentSum int, loanSum *int) {
	if *loanSum < 500000 {
		fmt.Println("Почти погашен кредит")
	} else {
		fmt.Println("Оплачено: ", repaymentSum)
		*loanSum -= repaymentSum
		fmt.Println("Остаток по кредиту: ", *loanSum)
	}
}
