package main

import "fmt"

// 3. Работа с указателями на числа
//    • Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement().
//   	Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
//    • Методы:
//        ◦ Increment() для увеличения значения числа на 1.
//        ◦ Decrement() для уменьшения значения числа на 1.

type PointerOperations interface {
	Increment()
	Decrement()
}
type IntPointer struct {
	num *int
}

func (ip *IntPointer) Increment() {
	num := *ip.num + 1
	fmt.Printf("num = %d\n", num)
}
func (ip *IntPointer) Decrement() {
	num := *ip.num - 1
	fmt.Printf("num = %d\n", num)
}

func Operate(p PointerOperations) {
	p.Increment()
	p.Decrement()
}

func main() {
	n := 10
	intPointer := IntPointer{num: &n}
	Operate(&intPointer)

}
