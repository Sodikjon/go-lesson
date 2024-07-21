package main

import "fmt"

func main() {
	//var x int = 5
	//fmt.Println(x)
	//b := &x
	//fmt.Println(&x)
	//fmt.Println(b)
	//fmt.Println(*b)

	//var n int = 5
	//fmt.Println("before: ", n)
	//updateValue(&n)
	//fmt.Println("after: ", n)

	//var a, b int = 7, 9
	//fmt.Println("before: ", a, b)
	//swap(&a, &b)
	//fmt.Println("after: ", a, b)

	//var x int
	//var p *int
	//printValue(p)

	//var y int = 7
	//fmt.Println("before", y)
	//increment(&y)
	//fmt.Println("after", y)

	//var x int = -3
	//fmt.Println(isPositive(&x))

	//var s string = "Hello"
	//fmt.Printf("String Before %q \n", s)
	//changeString(&s)
	//fmt.Printf("String After %q\n", s)

	//var n int = 0
	//fmt.Println("Is number even? ", isEven(&n))

	//var ptr *int
	//fmt.Println("Is it nil?", checkNil(ptr))

	//var math, physics, chemistry int
	//mathPtr := &math
	//physicsPtr := &physics
	//chemistryPtr := &chemistry
	//
	//addGrade(mathPtr, 85)
	//addGrade(physicsPtr, 90)
	//addGrade(chemistryPtr, 78)
	//fmt.Printf("Средний бал по предметам: %.2f\n", printAverageGrade(mathPtr, physicsPtr, chemistryPtr))
}

// Напишите функцию updateValue, которая принимает указатель на целое число и увеличивает его значение на 10.
func updateValue(p *int) {
	*p = *p + 10
}

// Напишите функцию swap, которая меняет местами значения двух переменных с использованием указателей.
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Напишите функцию printValue, которая принимает указатель на целое число и выводит его значение.
// Если указатель равен nil, выведите сообщение "Указатель пуст".
func printValue(p *int) {
	if p == nil {
		fmt.Println("Указатель пуст")
	} else {
		fmt.Println("Значение: ", *p)
	}
}

// Напишите функцию increment, которая принимает указатель на целое число и увеличивает его значение на 1.
func increment(p *int) {
	*p = *p + 1
}

// Напишите функцию isPositive, которая принимает указатель на целое число и возвращает true,
// если значение положительное, и false в противном случае.
func isPositive(p *int) bool {
	if *p <= 0 {
		return false
	}
	return true
}

// Напишите функцию changeString, которая принимает указатель на строку и добавляет к ней слово "Go".
func changeString(p *string) {
	*p = *p + " Go"
}

// Напишите функцию double, которая принимает указатель на целое число и удваивает его значение.
func double(p *int) {
	*p *= *p
}

// Напишите функцию isEven, которая принимает указатель на целое число и возвращает true, если число четное,
// и false, если нечетное.
func isEven(p *int) bool {
	if *p%2 == 0 {
		return true
	}
	return false
}

// Напишите функцию checkNil, которая принимает указатель на целое число и проверяет, является ли он нулевым (nil).
func checkNil(p *int) bool {
	if p == nil {
		return true
	}
	return false
}

// Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя.
func vote(candidatePtr *string, votesPtr *int) {
	*votesPtr = *votesPtr + 1
}

func winner(candidate1VotesPtr *int, candidate2VotesPtr *int) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Candidate1"
	} else if *candidate2VotesPtr > *candidate1VotesPtr {
		return "Candidate2"
	} else {
		return "Both"
	}
}

// Напишите программу для учета оценок по предметам студентов. Реализуйте функции для добавления оценки по предмету и вывода среднего балла.
func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}

func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemistryPtr + *chemistryPtr
	return float64(sum) / 3.0
}
