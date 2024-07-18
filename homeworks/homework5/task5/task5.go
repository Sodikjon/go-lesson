package main

import "fmt"

func main() {
	// Напишите программу, которая находит наибольший общий делитель (НОД) двух чисел, используя алгоритм Евклида.

	var a, b, gcd int

	fmt.Println("введите первое цисло")
	fmt.Scan(&a)

	fmt.Println("введите второе цисло")
	fmt.Scan(&b)

	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	gcd = a + b
	fmt.Printf("наибольший общий делитель равен %d", gcd)
}
