package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// Area
// Метод для вычисления площади
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter
// Метод для вычисления периметра
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Describe
// Метод для вывода всех характеристик прямоугольника
func (r Rectangle) Describe() {
	fmt.Printf("Width: %v, Height: %v, Area: %v, Perimeter: %v\n", r.width, r.height, r.Area(), r.Perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	rect.Describe() // Output: Width: 10, Height: 5, Area: 50, Perimeter: 30

}
