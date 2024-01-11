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
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(5 * time.Second)
		timeoutCh <- time.Now() //signaling with the time stamp
	}()

LOOP:
	for i := 1; ; i++ {
		select {
		case <-timeoutCh:
			fmt.Println("timeout triggered")
			break LOOP
		case ch <- 10 * i:
			time.Sleep(500 * time.Millisecond)
		}
	}
	close(ch) // producer does not own this channel. it was received from the consumer. so the producer has no right to close the channel.
}
