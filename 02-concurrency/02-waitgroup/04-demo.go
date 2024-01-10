/* WaitGroup */
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "Number of goroutines to spin")
	flag.Parse()

	fmt.Printf("Spinning %d goroutines.. Hit ENTER to start!\n", count)
	fmt.Scanln()

	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go fn(wg, i)
	}
	wg.Wait()
	fmt.Println("All goroutines completed.. Hit ENTER to shutdown")
	fmt.Scanln()
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	return
}
