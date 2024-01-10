/* WaitGroup */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the wg counter by 1
	go f1()   // scheduling f1() to be executed through the scheduler (in future)
	f2()

	wg.Wait() //block until the wg counter becomes 0 (counter = 0 by default)
}

func f1() {
	// fmt.Println("f1 invoked")
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
