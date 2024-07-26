package main

import "fmt"

func main() {
	//var age Age = 19
	//fmt.Println(isTeenager(age))

	//var speed Speed = 120.01
	//fmt.Println(isAllowedSpeed(speed))

	//var score Score = -10
	//fmt.Println(scoreInfo(score))

	var bmi BMI = 23.5
	fmt.Println(whichBMICategory(bmi))
}

// 2. Проверка возраста
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек подростком (возраст от 13 до 19 лет включительно) или нет.
type Age int

func isTeenager(age Age) string {
	if age <= 19 && age >= 13 {
		return "Является подростком!"
	} else {
		return "Не является подростком!"
	}
}

// 3. Проверка скорости
// Определите тип Speed на основе float64. Напишите функцию, которая принимает скорость и возвращает сообщение о том,
// является ли она допустимой (от 0 до 120 км/ч включительно) или нет.
type Speed float64

func isAllowedSpeed(speed Speed) string {
	if speed >= 0 && speed <= 120 {
		return "Скорость допустимая!"
	}
	return "Скорость не допустима!"
}

// 4. Проверка счета
// Определите тип Score на основе int. Напишите функцию, которая принимает счет и возвращает сообщение о том,
// является ли он положительным, отрицательным или нулевым.
type Score int

func scoreInfo(score Score) string {
	if score > 0 {
		return "Счет положительный."
	} else if score < 0 {
		return "Счет отрицательный."
	} else {
		return "Счет нулевой"
	}
}

// 5. Классификация BMI
// Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и возвращает сообщение о том,
// в какой категории оно находится: "Underweight", "Normal weight", "Overweight", или "Obesity".
// Underweight 0.0 – 18.4 Normal range 18.5 – 24.9 Overweight 25.0 – 29.9 Obese 30.0 +
type BMI float64

func whichBMICategory(bmi BMI) string {
	if bmi > 0 && bmi < 18.4 {
		return "Underweight"
	} else if bmi > 18.5 && bmi < 24.9 {
		return "Normal weight"
	} else if bmi > 25.0 && bmi < 29.9 {
		return "Overweight"
	} else if bmi > 30.0 {
		return "Obesity"
	} else {
		return "BMI not identified"
	}
}
