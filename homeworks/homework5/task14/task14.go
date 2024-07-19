package main

import (
	"fmt"
)

func main() {
	// Напишите программу, которая выводит таблицу умножения от 1 до n, где n - введенное пользователем число.
	var n int

	fmt.Print("Введите число n для таблицы умножения: ")
	fmt.Scan(&n)

	printMultiplicationTable(n)
}

func printMultiplicationTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d \n", i, j, i*j)
		}
		fmt.Println()
	}
}
