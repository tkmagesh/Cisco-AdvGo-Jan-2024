/* WaitGroup */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1) // increment the wg counter by 1
		go f1(wg) // scheduling f1() to be executed through the scheduler (in future)
	}
	f2()

	wg.Wait() //block until the wg counter becomes 0 (counter = 0 by default)
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	// fmt.Println("f1 invoked")
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
	return
}

func f2() {
	fmt.Println("f2 invoked")
}
