package main

import "sync"

import "fmt"

func main() {
	// mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		// mutex.Lock()
		wg.Add(1)
		go func(i int) {
			fmt.Printf("%v番目に作られたGoルーチン", i)
			// mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
