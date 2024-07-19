package main

import "fmt"

func main() {
	// Напишите программу, которая вычисляет произведение всех чисел от 1 до введённого числа n,
	// но прекращает выполнение, если результат превышает 1000.

	var n int
	var result int = 1

	fmt.Println("Введите число: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		result = result * i
		if result > 1000 {
			fmt.Println("Результат больше 1000. Конец... ")
			break
		}
		fmt.Printf("Temp %d -- %d\n", result, i)
	}
	fmt.Printf("Результат %d", result)

}
