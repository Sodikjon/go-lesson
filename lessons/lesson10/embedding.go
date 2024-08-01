package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

type Employee struct {
	Person
	position string
}

// Introduce
// Переопределение метода Introduce для Employee
//func (e Employee) Introduce() {
//	fmt.Printf("Hi, I'm %s, a %s, and I'm %d years old.\n", e.name, e.position, e.age)
//}

func main() {
	p := Person{name: "Alice", age: 30}
	e := Employee{Person: Person{name: "Bob", age: 40}, position: "Manager"}
	p.Introduce() // Output: Hi, I'm Alice, and I'm 30 years old.
	e.Introduce() // Output: Hi, I'm Bob, a Manager, and I'm 40 years old.

}

