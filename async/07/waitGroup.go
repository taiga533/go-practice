package main

import (
	"fmt"
	"sync"
)

func main() {
	// waitGroup構造体のポインタ型
	waitGroup := &sync.WaitGroup{}

	// goroutineが実行されるときにWaitGroupのカウントを1増やす
	waitGroup.Add(1)
	go balthasar(waitGroup)
	waitGroup.Add(1)
	go casper(waitGroup)
	waitGroup.Add(1)
	go merchior(waitGroup)

	// waitGroup内のカウントが0になるまで待つ
	waitGroup.Wait()
	fmt.Println("マギシステム終了")

}

func balthasar(waitGroup *sync.WaitGroup) {
	fmt.Println("バルタザール承認")
	waitGroup.Done()
}

func casper(waitGroup *sync.WaitGroup) {
	fmt.Println("キャスパー承認")
	waitGroup.Done()
}

func merchior(waitGroup *sync.WaitGroup) {
	fmt.Println("メルキオール承認")
	waitGroup.Done()
}
