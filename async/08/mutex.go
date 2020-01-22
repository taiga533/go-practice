package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	value string
	mutex *sync.Mutex
}

// 1度そのまま実行したら、コメントアウトを外して再度実行してみよう！
func main() {
	sampleItem := &item{
		value: "",
		mutex: &sync.Mutex{},
	}
	messages := []string{
		"hello",
		"world",
		"kawasaki",
		"taiga",
	}
	for _, message := range messages {
		// sampleItem.mutex.Lock()
		go sampleItem.update(message)
	}
	time.Sleep(3 * time.Second)
}

func (i *item) update(value string) {
	i.value = value
	fmt.Printf("itemの値は%vに更新されました。\n", i.value)
	// i.mutex.Unlock()
}
