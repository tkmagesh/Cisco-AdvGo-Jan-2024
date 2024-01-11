package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	ch <- 10
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	ch <- 20
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
}
