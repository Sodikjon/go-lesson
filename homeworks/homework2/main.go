package main

import (
	"fmt"
	"math"
)

func main() {
	// 1) Дана сторона квадрата a. Найти его периметр P = 4·a.
	var a int
	a = 5
	p := 4 * a
	fmt.Println(p)

	// 2) Дана сторона квадрата a. Найти его площадь S = 2 * a
	s := 2 * a
	fmt.Println(s)

	// 3) Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b).
	b := 4
	s1 := a * b
	fmt.Println(s1)

	p1 := 2 * (a + b)
	fmt.Println(p1)

	// 4) Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14.
	d := 6
	pi := 3.14

	l := pi * float64(d)
	fmt.Println(l)

	// 5) Дана длина ребра куба a. Найти объем куба V = a3 и площадь его поверхности S = 6·a
	a = 4
	v := a * 3
	fmt.Println(v)
	s2 := 6 * a
	fmt.Println(s2)

	// 6) Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и площадь поверхности
	//    S = 2·(a·b + b·c + a·c).
	a = 5
	b = 6
	c := 7

	v1 := a * b * c
	fmt.Println(v1)

	s3 := 2 * (a*b + b*c + a*c)
	fmt.Println(s3)

	// 7) Найти длину окружности L и площадь круга S заданного радиуса R: L = 2·π·R, S = π·R
	//    В качестве значения π использовать 3.14.
	r := 2
	l1 := float64(2) * pi * float64(r)
	fmt.Println(l1)

	// 8) Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.
	a = 3
	b = 4
	avg := (a + b) / 2
	fmt.Println(avg)

	// 9) Даны два неотрицательных числа a и b. Найти их среднее геометрическое, то есть квадратный корень из их
	//    произведения: √a·b.
	var a1 uint = 3
	var b1 uint = 4

	geoAvg := math.Sqrt(float64(a1) * float64(b1))
	fmt.Println(geoAvg)

	// 10) Даны два ненулевых числа. Найти сумму, разность, произведение и частное их квадратов.
	a2 := 5
	b2 := 7
	aSquared := math.Pow(float64(a2), 2)
	bSquared := math.Pow(float64(b2), 2)

	sum := aSquared + bSquared
	fmt.Println(sum)

	diff := aSquared - bSquared
	fmt.Println(diff)

	product := aSquared * bSquared
	fmt.Println(product)

	quotient := aSquared / bSquared
	fmt.Println(quotient)

	// 11) Дано расстояние L в сантиметрах. Используя операцию деления нацело, найти количество полных метров
	//     в нем (1 метр = 100 см).
	lSm := 172
	lM := float64(lSm) / 100
	fmt.Println(lM)

	// 12) Дана масса M в килограммах. Используя операцию деления нацело, найти количество полных тонн в ней
	//     (1 тонна = 1000 кг).
	mKg := 84
	mT := float64(mKg) / 1000
	fmt.Println(mT)

	// 13) Дан размер файла в байтах. Используя операцию деления нацело, найти количество полных килобайтов, которые
	//     занимает данный файл (1 килобайт = 1024 байта).
	b3 := 980
	kB := float64(b3) / 1024
	fmt.Println(kB)

	// 14) Даны целые положительные числа A и B (A > B). На отрезке длины A размещено максимально возможное количество
	//     отрезков длины B (без наложений). Используя операцию деления нацело, найти количество отрезков B,
	//     размещенных на отрезке A.
	a3 := 9
	b4 := 4

	count := a3 / b4
	fmt.Println(count)

	// 15) Даны целые положительные числа A и B (A > B). На отрезке длины A размещено максимально возможное количество
	//     отрезков длины B (без наложений). Используя операцию взятия остатка от деления нацело, найти длину незанятой
	//     части отрезка A.
	a5 := 8
	b5 := 5
	remainder := a5 % b5
	fmt.Println(remainder)
}
