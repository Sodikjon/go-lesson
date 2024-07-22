package main

import "fmt"

func main() {
	var dailyExpenseLimit float64 = 5000
	makeExpense(2500.00, &dailyExpenseLimit)
	makeExpense(1500.00, &dailyExpenseLimit)
	makeExpense(1500.00, &dailyExpenseLimit)
	makeExpense(500.00, &dailyExpenseLimit)
}

// 7. Учет ежедневных расходов с лимитом
// • Начальный лимит расходов в день равен 5000 рублей.
// • Реализуйте функции для добавления расходов в течение дня.
// • Выводите текущую сумму расходов после каждого добавления.
// • Если сумма расходов превышает лимит, выведите сообщение "Превышен дневной лимит".
func makeExpense(expense float64, dailyExpenseLimit *float64) {
	fmt.Printf("Ваш дневной лимит %.2f\n", *dailyExpenseLimit)
	if *dailyExpenseLimit <= 0 {
		fmt.Printf("Превышен дневной лимит\n")
	} else {
		*dailyExpenseLimit = *dailyExpenseLimit - expense
		fmt.Printf("Вы потратили %.2f\n", expense)
	}
}
