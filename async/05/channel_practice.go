package main

import "fmt"

/*
# 課題3
channelの挿入時に起こるブロッキングを体感しよう

# 条件
messagesチャンネルから取り出した「さようなら」を使う

# 期待する結果
標準出力に
さようなら
と表示される
*/
func main() {
	messages := make(chan string, 1)
	messages <- "こんにちは"
	go func() {
		messages <- "さようなら"
	}()
	fmt.Println(<-messages)
}
