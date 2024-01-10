package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling f1() to be executed through the scheduler (in future)
	f2()

	// DO NOT DO THIS IN PRODUCTION
	time.Sleep(100 * time.Millisecond) // blocking the main() execution so that the scheduler can look for other goroutines scheduled and execute them

	// fmt.Scanln()
}

func f1() {
	// fmt.Println("f1 invoked")
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
