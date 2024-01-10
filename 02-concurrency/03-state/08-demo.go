/* concurrent safe data manipulation using sync/atomic */

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

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
	atomic.AddInt64(&count, 1)
}
