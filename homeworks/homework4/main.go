package main

import (
	"fmt"
)

func main() {
	//GetGrade()
	//GetDiscount()
	//GetTemperatureDescription()
	//fmt.Println(CheckNumber(2))
	//fmt.Println(CheckAge(18))
	//fmt.Println(CheckPassword("123"))
	//fmt.Printf("Сумма модулей: %d\n", Add(5, 8))
	//fmt.Printf("Равны?: %s\n", CompareStrings("hello", "world"))
	//fmt.Printf("Кто больше?: %d\n", Max(6, 9))

	//var n1, n2 int = 10, 5
	//
	//operation := func(a, b int) int {
	//	var result int
	//	if a >= b {
	//		result = a - b
	//	} else {
	//		result = b - a
	//	}
	//	return result
	//}
	//
	//diff := operation(n1, n2)
	//fmt.Printf("Разность между %d и %d: %d\n", n1, n2, diff)

	//concat := func(s1, s2 string) string {
	//	if s1 == "" && s2 == "" {
	//		return ""
	//	} else if s1 == "" {
	//		return s2
	//	} else if s2 == "" {
	//		return s1
	//	} else {
	//		return s1 + " " + s2
	//	}
	//}
	//
	//result := concat("Hello", "")
	//fmt.Printf("Результат: %s\n", result)

	//multiply := func(a, b int) int {
	//	var result int
	//
	//	if a < 0 {
	//		a = -a
	//	}
	//
	//	if b < 0 {
	//		b = -b
	//	}
	//
	//	result = a * b
	//	return result
	//}
	//
	//product := multiply(5, -3)
	//fmt.Printf("Произведение модулей 5 и -3: %d\n", product)
	//
	//product = multiply(-4, -2)
	//fmt.Printf("Произведение модулей -4 и -2: %d\n", product)

	//add := func(x, y int) int {
	//	return x + y
	//}
	//
	//multiply := func(x, y int) int {
	//	return x * y
	//}
	//
	//sum := ApplyOperation(-5, 3, add)
	//product := ApplyOperation(-4, -2, multiply)
	//
	//fmt.Printf("Сумма модулей -5 и 3: %d\n", sum)
	//fmt.Printf("Произведение модулей -4 и -2: %d\n", product)

	//check := func(a int) bool {
	//	if a != 0 {
	//		return true
	//	}
	//	return false
	//}
	//checked := CheckCondition(0, check)
	//fmt.Printf("Проверяем состояние: %s\n", checked)

	//formatString := func(s string) string {
	//	return strings.ToUpper(s)
	//}
	//
	//formated := FormatAndPrint("", formatString)
	//fmt.Println(formated)

	//multiplier3 := CreateMultiplier(-3)
	//fmt.Println(multiplier3(-2))

	//greeterHello := CreateGreeter("Welcome")
	//fmt.Println(greeterHello(""))

	validatorQwerty := CreateValidator("Qwerty")
	fmt.Println(validatorQwerty("qwerty"))
}

func GetGrade() {
	var score int
	fmt.Println("Введите оценку ")
	fmt.Scanf("%d", &score)
	switch score {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Printf("B")
	case 3:
		fmt.Printf("C")
	}
}

func GetDiscount() {
	var amount int

	fmt.Println("Введите сумму покупки ")
	fmt.Scanf("%d", &amount)

	if amount > 1000 {
		fmt.Println("Скидка 10%")
	} else if amount > 500 && amount <= 1000 {
		fmt.Println("Скидка 5%")
	} else if amount <= 500 {
		fmt.Println("Скидка 0%")
	}
}

func GetTemperatureDescription() {
	var temperature int

	fmt.Println("Введите температуру ")
	fmt.Scanf("%d", &temperature)

	if temperature < 10 {
		fmt.Println("Холодно")
	} else if temperature >= 10 && temperature < 25 {
		fmt.Println("Тепло")
	} else if temperature >= 25 {
		fmt.Println("Жарко")
	}
}

func CheckNumber(number int) string {

	if number < 0 {
		return "Отрицательное"
	} else if number == 0 {
		return "Ноль"
	} else {
		return "Положительное"
	}
}

func CheckAge(age int) string {

	if age < 18 {
		return "Несовершеннолетний"
	} else {
		return "Совершеннолетний"
	}
}

func CheckPassword(password string) string {

	if password == "1234" {
		return "Пароль верный"
	} else {
		return "Пароль неверный"
	}
}

func Add(a, b int) int {
	var sumOfAbs int

	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	sumOfAbs = a + b

	return sumOfAbs
}

func CompareStrings(s1, s2 string) string {

	if s1 == s2 {
		return "равны"
	} else {
		return "не равны"
	}
}

func Max(n1, n2 int) int {

	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func ApplyOperation(a, b int, operation func(int, int) int) int {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	result := operation(a, b)

	return result
}

func CheckCondition(a int, check func(int) bool) string {

	result := check(a)

	if result == true {
		return "Условие выполнено"
	} else {
		return "Условие не выполнено"
	}
}

func FormatAndPrint(s string, formatS func(s1 string) string) string {
	result := formatS(s)

	if result == "" {
		return "Пустая строка"
	} else {
		return result
	}
}

func CreateMultiplier(a int) func(int) int {

	multiply := func(a1 int) int {
		var aAbs int
		if a1 < 0 {
			aAbs = -a1
		}
		return aAbs * a
	}

	return multiply
}

func CreateGreeter(s string) func(name string) string {

	return func(name string) string {
		if name == "" {
			return s + " Гость"
		}
		return s + " " + name
	}
}

func CreateValidator(s string) func(password string) bool {

	return func(password string) bool {
		if password == s {
			return true
		}
		return false
	}
}
