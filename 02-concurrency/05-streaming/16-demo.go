package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
	fmt.Println("All data received")

}

func genData(ch chan<- int) {
	count := rand.Intn(20)
	fmt.Println(count)
	for i := 1; i <= count; i++ {
		ch <- 10 * i
	}
	close(ch)
}
