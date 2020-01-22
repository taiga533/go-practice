package main

import "fmt"

func main() {
	messages := make(chan string)
	fmt.Print(len(messages))
	go func() {
		messages <- "hello, world"
	}()
	fmt.Print(<-messages)
}
