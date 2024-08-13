package main

import (
	"fmt"
	"strings"
)

// 	2. Форматирование строки
//    • Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower(). Реализуйте структуру MyFormatter,
//   	которая будет работать со строкой и реализует этот интерфейс.
//    • Методы:
//        ◦ ToUpper() string для преобразования строки в верхний регистр.
//        ◦ ToLower() string для преобразования строки в нижний регистр.

type Formatter interface {
	ToUpper() string
	ToLower() string
}

type MyFormatter struct {
	text string
}

func (mf MyFormatter) ToUpper() string {
	mf.text = strings.ToUpper(mf.text)
	return mf.text
}
func (mf MyFormatter) ToLower() string {
	mf.text = strings.ToLower(mf.text)
	return mf.text
}

func main() {
	myFormatter := new(MyFormatter)
	myFormatter.text = "hello world"
	fmt.Printf("%s\n", myFormatter.ToLower())
	fmt.Printf("%s\n", myFormatter.ToUpper())
}
