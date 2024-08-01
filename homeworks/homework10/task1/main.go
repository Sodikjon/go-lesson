package main

import "fmt"

// 1. Книга и Автор
//    • Описание: Реализуйте структуру Book с полями title и author, а также метод GetDetails, который возвращает
//   	строку с деталями книги.
//  	Реализуйте структуру Library с массивом книг и метод AddBook, добавляющий книгу в библиотеку.
//    • Методы:
//        ◦ GetDetails() string для структуры Book
//        ◦ AddBook(book Book) для структуры Library

type Book struct {
	title  string
	author string
}

func (r *Book) GetDetails() string {
	return fmt.Sprintf("Название книги: %s, Автор: %s\n", r.title, r.author)
}

type Library struct {
	books []Book
}

func (r *Library) AddBook(book Book) {
	r.books = append(r.books, book)
}

func main() {
	b1 := Book{"Happy Nation", "Henry Ost"}
	b2 := Book{"Nation Z", "Charles Bobby"}

	fmt.Println(b1.GetDetails(), b2.GetDetails())

	lib := Library{}
	lib.AddBook(b1)
	lib.AddBook(b2)

	fmt.Printf("Our library has %v", lib)

}
