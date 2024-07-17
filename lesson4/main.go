package main

import "fmt"

func main() {
	//num := 6
	//if num > 9 && num < 100 {
	//	fmt.Println("Двузначное")
	//}
	//
	//m := 78
	//if m >= 0 && m < 10 {
	//	fmt.Println("Однозначное")
	//} else if m >= 10 && m < 100 {
	//	fmt.Println("Двузначное")
	//} else if m >= 100 && m < 1000 {
	//	fmt.Println("трехзначное")
	//} else {
	//	fmt.Println("Слишком большое")
	//}
	//
	//c := 6
	//d := 5
	//
	//if c > d {
	//	fmt.Println(c)
	//} else {
	//	fmt.Println(d)
	//}

	//var n1, n2, n3 int = 5, 7, 8
	//var n1, n2, n3 int = -4, 2, 1
	//var n1, n2, n3 int = 6, 6, 3

	// n1<=n2&&n2<=n3
	// n1<=n3&&n3<=n2
	// n2<=n1&&n1<=n3
	// n2<=n3&&n3<=n1
	// n3<=n1&&n1<=n2
	// n3<=n2&&n2<=n1

	//if n1 <= n2 && n2 <= n3 {
	//	fmt.Println(n3, n2, n1)
	//} else if n1 <= n3 && n3 <= n2 {
	//	fmt.Println(n2, n3, n1)
	//} else if n2 <= n1 && n1 <= n3 {
	//	fmt.Println(n3, n1, n2)
	//} else if n2 <= n3 && n3 <= n1 {
	//	fmt.Println(n1, n3, n2)
	//} else if n3 <= n1 && n1 <= n2 {
	//	fmt.Println(n2, n1, n3)
	//} else {
	//	fmt.Println(n1, n2, n3)
	//}

	//nn := -1
	//switch nn {
	//case 1:
	//	fmt.Println("Positive")
	//	fallthrough
	//case -1:
	//	fmt.Println("Negative")
	//case 0:
	//	fmt.Println("Unsigned")
	//}

	//if n1 <= n2 && n2 <= n3 {
	//	fmt.Println(n3 + n2)
	//} else if n1 <= n3 && n3 <= n2 {
	//	fmt.Println(n2 + n3)
	//} else if n2 <= n1 && n1 <= n3 {
	//	fmt.Println(n3 + n1)
	//} else if n2 <= n3 && n3 <= n1 {
	//	fmt.Println(n1 + n3)
	//} else if n3 <= n1 && n1 <= n2 {
	//	fmt.Println(n2 + n1)
	//} else {
	//	fmt.Println(n1 + n2)
	//}

	//isLeapYear := 1900
	//
	//if isLeapYear%4 == 0 || isLeapYear%100 != 0 {
	//	if isLeapYear%100 != 0 || isLeapYear%400 == 0 {
	//		fmt.Println("Високосный")
	//
	//	} else {
	//		fmt.Println("Обычный")
	//	}
	//} else {
	//	fmt.Println("Обычный")
	//}

	var month int
	fmt.Println("Введите номер месяца: ")
	fmt.Scanln(&month)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Printf("В месяце %d имеется %d дней", month, 31)
	case 2:
		fmt.Printf("В месяце %d имеется %d дней", month, 28)
	case 4, 6, 9, 11:
		fmt.Printf("В месяце %d имеется %d дней", month, 30)
	}

}
