package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genData(ch)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("All data received")

}

// producer (should produce the data until 5 seconds elapses)
func genData(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- 10 * i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // producer does not own this channel. it was received from the consumer. so the producer has no right to close the channel.
}
