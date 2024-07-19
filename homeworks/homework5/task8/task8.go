package main

import "fmt"

func main() {
	// Напишите программу, которая выводит все делители числа.

	var number int
	fmt.Print("Введите число для поиска его делителей: ")
	fmt.Scan(&number)

	if number <= 0 {
		fmt.Println("Пожалуйста, введите положительное число.")
		return
	}

	printDivisors(number)
}

func printDivisors(n int) {
	fmt.Printf("Делители числа %d:\n", n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
