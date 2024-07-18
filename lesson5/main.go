package main

import "fmt"

func main() {
	//var i int
	//for i = 0; i <= 9; i++ {
	//	fmt.Print(i)
	//}
	//fmt.Println(i)

	//var num = 546
	//for num > 0 {
	//	a := num % 10
	//	fmt.Print(a)
	//	num = num / 10
	//}

	//for n := 546; n > 0; n /= 10 {
	//	a := n % 10
	//	fmt.Print(a)
	//}

	//for {
	//	time.Sleep(1 * time.Second)
	//	fmt.Print("Wuahahaha ")
	//}

	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		//continue
	//		break
	//	}
	//	fmt.Println(i)
	//
	//}

	//for i := 1; i < 10; i++ {
	//	for j := 1; j < 10; j++ {
	//		fmt.Printf("| %d * %d = %d |\n", i, j, i*j)
	//	}
	//	fmt.Println("|____________|")
	//}

	//var n = 100
	//var result float64 = 1
	//for i := 2; i < n; i++ {
	//	result = result + (float64(1) / float64(i))
	//	//fmt.Println(result)
	//}
	//fmt.Println(result)

	//For11. Дано целое число N (>0). Найти сумму
	//N2 + (N + 1)2 + (N + 2)2 + . . . + (2·N)2
	//(целое число).

	//var n = 5
	//var result int = 1
	//for i := 0; i <= n; i++ {
	//	result += (result + i) * (result + i)
	//	fmt.Println(result)
	//}
	//fmt.Println(result)

	//series2()
	series6()

}

func series2() {
	var p float64 = 1
	var n float64

	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		p = p * n

	}
	fmt.Println(p)
}

func series6() {
	//Дано целое число N и набор из N положительных вещественных
	//чисел. Вывести в том же порядке дробные части всех чисел из данного
	//набора (как вещественные числа с нулевой целой частью),
	//а также произведение всех дробных частей

	var n int
	var a, b float64
	var p float64 = 1
	fmt.Println("Enter N")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter A")
		fmt.Scan(&a)
		b = a - float64(int(a))
		fmt.Println(b)
		p = p * b
		fmt.Println(p)
	}
	fmt.Printf("Result %.2f", p)
}
