package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 1)
	messages <- "wait a moment"
	go func() {
		fmt.Println("キューが開くのを待っています。")
		messages <- "キューが空いたのでデータを入れました。"
	}()
	time.Sleep(2 * time.Second)
	// 最初にキューに入れたデータを取り出す。
	<-messages
	// キューにデータが入るまで待つ。
	fmt.Println(<-messages)
}
