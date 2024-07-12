package main

import "fmt"

func main() {
	var number int
	var number8 int8
	var number16 int16
	var number32 int32
	var number64 int64

	var numberFloat32 float32
	var numberFloat64 float64

	var boolean bool
	var runeVar rune
	var byteVar byte
	var stringVar string

	var uNumber uint
	var uNumber8 uint8
	var uNumber16 uint16
	var uNumber32 uint32
	var uNumber64 uint64

	var complex64Var complex64
	var complex128Var complex128

	fmt.Println(number)
	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)

	fmt.Println(numberFloat32)
	fmt.Println(numberFloat64)

	fmt.Println(boolean)
	fmt.Println(runeVar)
	fmt.Println(byteVar)

	fmt.Println(stringVar)

	fmt.Println(uNumber)
	fmt.Println(uNumber8)
	fmt.Println(uNumber16)
	fmt.Println(uNumber32)
	fmt.Println(uNumber64)

	fmt.Println(complex64Var)
	fmt.Println(complex128Var)

	number = 7
	fmt.Println(number)

	number8 = 125
	fmt.Println(number8)

	number16 = -17
	fmt.Println(number16)

	number32 = 1990
	fmt.Println(number32)

	number64 = 202420242024
	fmt.Println(number64)

	numberFloat32 = 3.14159265358979323846
	fmt.Println(numberFloat32)

	numberFloat64 = 3.14159265358979323846264338327950288419716939937510
	fmt.Println(numberFloat64)

	boolean = true
	fmt.Println(boolean)

	runeVar = 'F'
	fmt.Println(runeVar)

	byteVar = 'F'
	fmt.Println(byteVar)

	stringVar = "HumoLab"
	fmt.Println(stringVar)

	uNumber = 123
	fmt.Println(uNumber)

	uNumber8 = 255
	fmt.Println(uNumber8)

	uNumber16 = 301
	fmt.Println(uNumber16)

	uNumber32 = 1111000022
	fmt.Println(uNumber32)

	uNumber64 = 111100002222
	fmt.Println(uNumber64)

	complex64Var = 123456 + 10i
	fmt.Println(complex64Var)
	complex128Var = 1234567890 + 44i
	fmt.Println(complex128Var)

	firstNumber := 3
	secondNumber := 5

	result1 := firstNumber + secondNumber
	result2 := firstNumber - secondNumber
	result3 := firstNumber * secondNumber
	result4 := float64(firstNumber) / float64(secondNumber)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
