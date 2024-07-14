package main

import (
	"fmt"
	"strings"
)

func main() {
	DisplaySeparator()
	PrintGreeting()
	NotifyStart()

	DisplaySeparator()

	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	fmt.Println(IsServerActive())

	DisplaySeparator()

	PrintNumber(6)
	GreetUser("Sodiqjon")
	ToggleLight(false)

	DisplaySeparator()

	fmt.Println(Add(3, 5))
	fmt.Println(Concat("salom", "dustam"))
	fmt.Println(IsEven(5))

	DisplaySeparator()

	adder := func(a, b int) int {
		return a + b
	}
	fmt.Println(adder(3, 6))

	concatenator := func(s1, s2 string) string {
		return s1 + s2
	}
	fmt.Println(concatenator("dustam", "kisti?"))

	isPositive := func(x int) bool {
		if x > 0 {
			return true
		} else {
			return false
		}
	}
	fmt.Println(isPositive(-5))

	DisplaySeparator()

	fmt.Println(Calculate(5, 7, calculator))
	Execute(false, printBool)
	fmt.Println(Apply(7, getInt))

	DisplaySeparator()

	factored := Multiplier(5)
	fmt.Println(factored(3))

	repeater := StringRepeater("Salom! ", 3)
	fmt.Println(repeater())

	isCondition := ConditionalPrinter(true)
	isCondition()

	DisplaySeparator()

}

func PrintGreeting() {
	fmt.Println("Hello, World!")
}

func DisplaySeparator() {
	fmt.Println("--------------------")
}

func NotifyStart() {
	fmt.Println("Program started")
}

func GetWelcomeMessage() string {
	return "Welcome!"
}

func GetPiValue() float64 {
	return 3.14
}

func IsServerActive() bool {
	return true
}
func PrintNumber(number int) {
	fmt.Println(number)
}

func GreetUser(name string) {
	fmt.Println(name)
}

func ToggleLight(toggle bool) {
	if toggle {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

func Add(a, b int) int {
	c := a + b
	return c
}

func Concat(s1, s2 string) string {
	return s1 + s2
}

func IsEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func calculator(a, b int) int {
	return a + b
}

func Calculate(a, b int, f func(a, b int) int) int {
	result := f(a, b)
	return result
}

func printBool(bool2 bool) {
	fmt.Println(bool2)
}

func Execute(bool1 bool, f func(bool2 bool)) {
	f(bool1)
}

func getInt(x int) int {
	return x
}

func Apply(a int, f func(x int) int) int {
	return f(a)
}

func Multiplier(factor int) func(a int) int {
	return func(a int) int {
		return a * factor
	}
}

func StringRepeater(s string, n int) func() string {
	return func() string {
		return strings.Repeat(s, n)
	}
}

func ConditionalPrinter(condition bool) func() {
	return func() {
		if condition {
			fmt.Println("Yes")
		}
	}
}
