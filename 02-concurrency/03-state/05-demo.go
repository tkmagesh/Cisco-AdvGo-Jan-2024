package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Printf("count : %d\n", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}
