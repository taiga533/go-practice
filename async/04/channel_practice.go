package main

import (
	"fmt"
)

/*
# 課題2
channelを使って非同期処理の完了を待ってみよう！

# 条件
channelを使う

# 期待する結果
標準出力に
処理中
実行完了
と表示される
*/
func main() {
	go func() {
		fmt.Println("処理中")
	}()
	fmt.Println("実行完了")

}
