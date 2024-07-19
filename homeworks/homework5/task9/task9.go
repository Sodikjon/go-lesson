package main

import "fmt"

func main() {
	// Напишите программу, которая находит сумму цифр числа.

	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := sumOfDigits(number)
	fmt.Printf("Сумма цифр числа %d равна %d\n", number, result)
}

func sumOfDigits(n int) int {
	if n < 0 {
		n = -n // Работаем с модулем числа
	}

	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum
}
