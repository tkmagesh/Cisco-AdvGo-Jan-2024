package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	done := genData(ch)

	go func() {
		<-done
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("All data received")

}

// producer
func genData(ch chan<- int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		count := rand.Intn(20)
		fmt.Println(count)
		for i := 1; i <= count; i++ {
			ch <- 10 * i
		}
		close(done)
	}()
	return done
}
