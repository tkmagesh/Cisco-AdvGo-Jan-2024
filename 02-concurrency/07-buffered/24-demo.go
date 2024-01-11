package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panicked :", err)
			return
		}
		fmt.Println("Thank you")
	}()

	// divisor := 7
	divisor := 0

	// sequential
	/*
		result := divideSync(100, divisor)
		fmt.Println(result)
	*/

	// concurrent
	/*
		ch := divideAsync1(100, divisor)
		fmt.Println(<-ch)
	*/

	/*
		ch, err := divideAsync2(100, divisor)
		select {
		case result := <-ch:
			fmt.Println("result :", result)
		case e := <-err:
			fmt.Println("error :", e)
		}
	*/

	ch, _ := divideAsync2(100, divisor)
	result := <-ch
	fmt.Println("result :", result)

}

// sequential
func divideSync(x, y int) int {
	time.Sleep(2 * time.Second)
	result := x / y
	return result
}

// concurrent (not handling the panic)
func divideAsync1(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		result := x / y
		ch <- result
	}()
	return ch
}

// handling panic at the goroutine
func divideAsync2(x, y int) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error, 1)
	// errCh := make(chan error)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(error); ok {
					errCh <- e // will result in a deadlock if the errCh is NOT a buffered channel and the user has not initiated the receive operation
					return
				}
			}
		}()
		time.Sleep(2 * time.Second)
		result := x / y
		ch <- result
	}()
	return ch, errCh
}
