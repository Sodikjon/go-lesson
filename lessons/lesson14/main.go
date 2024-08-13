package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//// Открытие или создание файла для записи
	//file, err := os.Create("output.txt")
	//if err != nil {
	//	fmt.Println("Error creating file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//// Данные для записи
	//data := []byte("Hello, Writer!\n")
	//
	//// Запись данных в файл
	//n, err := file.Write(data)
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}
	//fmt.Printf("Wrote %d bytes to file.\n", n)

	//file1, err2 := os.Open("output.txt")
	//if err2 != nil {
	//	fmt.Println("Error opening file:", err2)
	//	return
	//}
	//defer file1.Close()

	//buf := make([]byte, 1024)
	//n1, err2 := file1.Read(buf)
	//if err2 != nil {
	//	fmt.Println("Error reading file:", err2)
	//	return
	//}
	//fmt.Printf("Read %d bytes: %s\n", n1, buf[:n1])

	//_, err2 = file1.Write([]byte("Hello, Go!"))
	//if err2 != nil {
	//	fmt.Println("Error writing to file:", err2)
	//	return
	//}
	//fmt.Println("Data written successfully")

	// В Go доступны следующие флаги, которые определяют режим работы с файлом:
	//
	//- os.O_RDONLY: Открывает файл только для чтения.
	//- os.O_WRONLY: Открывает файл только для записи.
	//- os.O_RDWR: Открывает файл для чтения и записи.
	//- os.O_APPEND: Открывает файл и добавляет новые данные в конец файла.
	//- os.O_CREATE: Создает файл, если он не существует.
	//- os.O_TRUNC: Очищает файл при открытии, если он существует.
	//- os.O_EXCL: Используется с os.O_CREATE, чтобы файл был создан только если он не
	//существует; если файл существует, возвращает ошибку.
	//
	//Флаги могут быть комбинированы с помощью побитового оператора |

	//// Открытие или создание файла с правами на чтение и запись
	//file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//fmt.Println("File opened successfully")
	//
	//_, err = file.Write([]byte(" Hello, Go Go Go!\n"))
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}
	//fmt.Println("Data written successfully")

	//// Открытие файла как источник потока данных
	//file, err := os.Open("example.txt")
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//// Чтение файла и вывод его содержимого по частям
	//buffer := make([]byte, 4) // Чтение по 1024 байта
	//for {
	//	n, err := file.Read(buffer)
	//	if err == io.EOF {
	//		break // Достигнут конец файла
	//	}
	//	if err != nil {
	//		fmt.Println("Error reading file:", err)
	//		return
	//	}
	//	time.Sleep(1 * time.Second)
	//	// Вывод прочитанной части файла
	//	fmt.Print(string(buffer[:n]))
	//}

	//var input string
	//fmt.Print("Enter your name: ")
	//fmt.Fscan(os.Stdin, &input)
	//fmt.Printf("Hello, %s!\n", input)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() { _ = out.Flush() }()

	var num int
	fmt.Fscan(in, &num)

	fmt.Fprintln(out, num)
}
