package main

import "fmt"

type Logger struct{}

// Log
// Метод с вариативными параметрами
func (Logger) Log(messages ...string) {
	fmt.Printf("%T\n", messages)
	for _, message := range messages {
		fmt.Println("Log:", message)
	}
}

func main() {
	logger := Logger{}
	logger.Log("Message 1", "Message 2", "Message 3") // Output: Log: Message 1
	//         Log: Message 2
	//         Log: Message 3
}
