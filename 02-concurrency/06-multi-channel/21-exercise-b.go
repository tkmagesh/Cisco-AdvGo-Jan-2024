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

	// timeoutCh := timeout(5 * time.Second)
	// replacing the above with time.After()
	timeoutCh := time.After(5 * time.Second)
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
	close(ch)
}

// poor man's implementation of time.After()
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now() //signaling with the time stamp
	}()
	return timeoutCh
}
