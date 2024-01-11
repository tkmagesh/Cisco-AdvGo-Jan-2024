package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go genData(ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All data received")

}

// producer
func genData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := rand.Intn(20)
	fmt.Println(count)
	for i := 1; i <= count; i++ {
		ch <- 10 * i
	}
}
