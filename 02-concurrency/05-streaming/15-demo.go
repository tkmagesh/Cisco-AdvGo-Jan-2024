package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}

}

func genData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- 10 * i
		time.Sleep(500 * time.Millisecond)
	}

}
