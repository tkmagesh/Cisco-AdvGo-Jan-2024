package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genData()
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All data received")
}

// producer
func genData() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		fmt.Println(count)
		for i := 1; i <= count; i++ {
			ch <- 10 * i
		}
		close(ch)
	}()
	return ch
}
