package main

import (
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(wg)

	}
	wg.Wait()
}
