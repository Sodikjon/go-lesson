package main

import "fmt"

// 2. Оценки студента
//    • Описание: Реализуйте структуру Student с полем grades (срез оценок).
//   	Реализуйте метод AddGrade, добавляющий оценку, и метод AverageGrade, возвращающий среднее значение оценок.
//    • Методы:
//        ◦ AddGrade(grade int)
//        ◦ AverageGrade() float64

type Student struct {
	grades []int
}

func (r *Student) AddGrade(grade int) {
	r.grades = append(r.grades, grade)
}
func (r *Student) AverageGrade() float64 {
	var sum float64
	for _, grade := range r.grades {
		sum += float64(grade)
	}
	return sum / float64(len(r.grades))
}
func main() {
	s1 := Student{grades: make([]int, 0)}
	s2 := Student{grades: make([]int, 0)}

	s1.AddGrade(70)
	s1.AddGrade(86)
	s1.AddGrade(65)

	s2.AddGrade(85)
	s2.AddGrade(87)

	fmt.Printf("Average grade for student 1 is: %.2f\n", s1.AverageGrade())
	fmt.Printf("Average grade for student 2 is: %.2f\n", s2.AverageGrade())
}
