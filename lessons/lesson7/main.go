package main

import "fmt"

//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	type Address struct {
//		City  string
//		State string
//	}
//
//	type Person2 struct {
//		Name    string
//		Age     int
//		Address Address
//	}
type Age int
type Number int
type Score int
type Operation func(int, int) int
type Count int
type Temperature float64
type Price float64
type User struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}
type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	//// Создаем и инициализируем структуру
	//p1 := Person{Name: "Alice", Age: 30}
	//fmt.Println(p1)
	//
	//p2 := Person{"Bob", 40}
	//fmt.Println(p2)
	//
	//p3 := new(Person)
	//p3.Name = "Charlie"
	//p3.Age = 50
	//fmt.Println(*p3)
	//
	//p4 := &Person{Name: "David", Age: 60}
	//p4.Age = 61
	//fmt.Println(*p4)
	//
	//p5 := Person2{
	//	Name: "John",
	//	Age:  33,
	//	Address: Address{
	//		City:  "New York",
	//		State: "NY",
	//	},
	//}
	//fmt.Println(p5)
	//fmt.Println(p5.Name)
	//fmt.Println(p5.Age)
	//var age Age
	//age = 18
	//fmt.Println(checkAge(age))

	//var number Number = 19
	//fmt.Println(isEven(number))

	//var score Score
	//score = 11
	//fmt.Println(checkScore(score))

	//var myAdder Operation = add
	//result1 := myAdder(2, 3)
	//fmt.Println(result1)
	//var mySubber Operation = sub
	//result2 := mySubber(5, 3)
	//fmt.Println(result2)
	//var myMuler Operation = mul
	//result3 := myMuler(2, 3)
	//fmt.Println(result3)
	//var myDever Operation = dev
	//result4 := myDever(6, 3)
	//fmt.Println(result4)

	//var counter Count = 10
	//coundown(counter)

	//var temperature Temperature = 0
	//checkTemperature(temperature)

	//var price Price = 100
	//fmt.Println(discounter(price))

	//var user1 User
	//user1.Name = "Alice"
	//user1.Age = 25
	//
	//userInfo(user1)

	//var product1 Product
	//product1.Name = "bananana"
	//product1.Price = 12
	//
	//fmt.Println(showPrice(product1))

	book1 := Book{Author: "Умари Хайём", Title: "Рубоёт", Pages: 177}

	bookInfo(book1)
}

// Abdulhalim Ergashev, [22.07.2024 18:52]
// // Определение возраста совершеннолетия
// // Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// // является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
func checkAge(age Age) string {
	if age < 18 {
		return "Не совершенолетний"
	} else {
		return "Совершенолетний"
	}
}

// Abdulhalim Ergashev, [22.07.2024 18:53]
// // Проверка на четность
// // Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
// // является ли оно четным или нечетным.
func isEven(number Number) string {
	if number%2 == 0 {
		return "Четное"
	} else {
		return "Нечетное"
	}
}

// Abdulhalim Ergashev, [22.07.2024 18:53]
// // Проверка допустимого диапазона
// // Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение,
// // находится ли она в допустимом диапазоне от 0 до 100
func checkScore(score Score) string {
	if score <= 100 && score >= 0 {
		return "В диапазоне"
	} else {
		return "Вне диапазона"
	}
}

// Abdulhalim Ergashev, [22.07.2024 18:54]
// // Арифметические операции
// // Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения,
// // вычитания и умножения. Используйте тип Operation для вызова этих функций
// Реализуем функцию, соответствующую типу Operation
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

//Abdulhalim Ergashev, [22.07.2024 18:54]
//// Сравнение чисел
//// Определите тип функции Comparator, которая принимает два int и возвращает bool. Напишите функции для проверки
//// равенства и сравнения больше/меньше. Используйте тип Comparator для вызова этих функций.
//

//Abdulhalim Ergashev, [22.07.2024 18:54]
//// Применение функции к числу
//// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
//// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.

// Abdulhalim Ergashev, [22.07.2024 19:22]
// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.
func coundown(c Count) {
	for i := c; i > 0; i-- {
		fmt.Println(i)
	}
}

// Abdulhalim Ergashev, [22.07.2024 19:22]
// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).
func checkTemperature(temperature Temperature) {
	if temperature < 0 {
		fmt.Println("ниже нуля")
	} else if temperature > 0 {
		fmt.Println("выше нуля")
	} else {
		fmt.Println("ноль")
	}
}

// Abdulhalim Ergashev, [22.07.2024 19:22]
// Применение скидки
// Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.
func discounter(price Price) Price {
	return price - (price * 0.1)
}

// Abdulhalim Ergashev, [22.07.2024 19:23]
// Информация о пользователе
// Создайте структуру User с полями Name (строка) и Age (целое число). Напишите функцию, которая
// принимает User и выводит информацию о нем.
func userInfo(user User) {
	fmt.Printf("Имя %s и возраст %d лет.\n", user.Name, user.Age)
}

// Abdulhalim Ergashev, [22.07.2024 19:23]
// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price). Напишите функцию, которая
// принимает Product и возвращает сообщение о его стоимости.
func showPrice(product Product) string {
	result := fmt.Sprintf("Данный продукт стоит %.2f денег", product.Price)
	return result
}

// Abdulhalim Ergashev, [22.07.2024 19:23]
// Информация о книге
// Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число). Напишите функцию,
// которая принимает Book и выводит информацию о книге.
func bookInfo(book Book) {
	fmt.Printf("Книга называется %s, её автор %s и в ней есть %d\n", book.Title, book.Author, book.Pages)
}

//Abdulhalim Ergashev, [22.07.2024 19:23]
// Изменение данных через указатель
// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person
// и обновляет его возраст. Выведите информацию до и после обновления.

//Abdulhalim Ergashev, [22.07.2024 19:23]
// Создание и изменение структуры через указатель
// Создайте структуру Rectangle с полями Width и Height. Напишите функцию, которая принимает указатель
// на Rectangle, вычисляет и обновляет его площадь, а затем выводит обновленную площадь.

//Abdulhalim Ergashev, [22.07.2024 19:24]
// Сравнение двух структур через указатели
// Создайте структуру Coordinate с полями X и Y. Напишите функцию, которая принимает указатели на две
// Coordinate и возвращает сообщение о том, равны ли они или нет.
