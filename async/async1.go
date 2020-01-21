package main

import (
	"fmt"
	"time"
)

/*
# 課題1
goroutineを使ってみよう
プログラムはF5を押せばデバックできるよ！

# 期待する結果
標準出力に
hello
kawasaki
と表示される
*/
func main() {
	go sayHello()
	fmt.Println("kawasaki")
	time.Sleep(2 * time.Second)
}

func sayHello() {
	fmt.Println("hello")
}
