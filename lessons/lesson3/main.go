package main

import "fmt"

func sum(a, b int) int {
	c := a + b
	return c
}

func printHello(name string) {
	fmt.Println("Hello", name, "!")
}

func printHello2() {
	fmt.Println("Hello!")
}

func inc(a int) {
	a++
	fmt.Println(a)
}

func main() {
	fmt.Println(sum(5, 4))
	fmt.Println(sum(6, 7))
	fmt.Println(sum(6, 9))
	//fmt.Println(sum(5.6, 9)) - не работает
	printHello("Sodiqjon")
	printHello2()

	a := 6
	inc(a)

	fmt.Println(a)

	f := sum2(8)
	fmt.Println(f)
	i := f()
	fmt.Println(i)
}

func sum2(a int) func() int {
	b := 6
	return func() int {
		return a + b
	}
}
