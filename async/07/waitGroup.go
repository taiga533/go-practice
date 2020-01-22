package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(1)
	go balthasar(waitGroup)
	waitGroup.Add(1)
	go casper(waitGroup)
	waitGroup.Add(1)
	go merchior(waitGroup)

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
