package main

import "fmt"

//// Rectangle
//// Определение структуры
//type Rectangle struct {
//	width, height float64
//}
//
//// Area
//// Метод с value receiver
//func (r Rectangle) Area() float64 {
//	return r.width * r.height
//}
//
//// Scale
//// Метод с pointer receiver
//func (r *Rectangle) Scale(factor float64) {
//	r.width *= factor
//	r.height *= factor
//}

//type Counter struct {
//	value int
//}
//
//func (c Counter) Increment() {
//	c.value++
//}

//type MySlice []int
//
//// Sum
//// Метод для пользовательского типа MySlice
//func (s MySlice) Sum() int {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	return sum
//}

// MyInt
// Определение нового типа на основе встроенного типа int
type MyInt int

// IsEven
// Метод для нового типа MyInt
func (m MyInt) IsEven() bool {
	return m%2 == 0
}

func main() {
	//rect := Rectangle{width: 10, height: 5}
	//fmt.Println("Area:", rect.Area()) // Output: Area: 50
	//// Изменение размера с использованием pointer receiver
	//rect.Scale(2)
	//fmt.Println("Scaled Area:", rect.Area()) // Output: Scaled Area: 200

	//c := Counter{value: 10}
	//c.Increment()
	//fmt.Println("Value after Increment:", c.value) // Output: Value after Increment: 10

	//s := MySlice{1, 2, 3, 4, 5}
	//fmt.Println("Sum:", s.Sum()) // Output: Sum: 15

	num := MyInt(5)
	fmt.Println("Is Even:", num.IsEven()) // Output: Is Even: true

}
