package main

import "fmt"

func main() {
	//var sqr UnaryOperation
	//sqr = square
	//fmt.Println(sqr(6))

	var op Operation
	op = add
	fmt.Println(op(3, 6))
	op = sub
	fmt.Println(op(6, 3))
	op = mul
	fmt.Println(op(3, 6))
	op = dev
	fmt.Println(op(6, 2))
}

// 6. Возведение в квадрат
//Определите тип функции UnaryOperation, которая принимает int и возвращает int.
//Напишите функцию для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции.

type UnaryOperation func(int) int

func square(num int) int {
	return num * num
}

// 10. Комбинирование функций.
//Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения,
//вычитания и умножения. Напишите функцию, которая принимает два int и массив функций Operation,
//и возвращает массив результатов применения этих функций.

type Operation func(int, int) int

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func dev(a int, b int) int {
	return a / b
}
