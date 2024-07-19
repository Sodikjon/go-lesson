package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Напишите программу, которая запрашивает у пользователя ввод числа и проверяет, является ли оно палиндромом.

	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if isPalindrome(number) {
		fmt.Printf("%d является палиндромом.\n", number)
	} else {
		fmt.Printf("%d не является палиндромом.\n", number)
	}

}

func isPalindrome(n int) bool {
	// Преобразуем число в строку
	str := strconv.Itoa(n)

	// Проверяем, является ли строка палиндромом
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
