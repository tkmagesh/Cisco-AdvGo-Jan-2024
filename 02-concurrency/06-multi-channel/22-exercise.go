/*
Write a program that keeps producing fib series (concurrently) at 500 millisecond intervals until the user hits ENTER key
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		close(done)
	}()
	ch := genFib(done)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

func genFib(done <-chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; {
			select {
			case <-done:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
