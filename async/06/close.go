package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			messages <- "処理中"
		}
		close(messages)
	}()
	// 無限ループ
	for {
		msg, ok := <-messages
		if ok {
			fmt.Println(msg)
			continue
		}
		fmt.Println("処理終了")
		break
	}
}
