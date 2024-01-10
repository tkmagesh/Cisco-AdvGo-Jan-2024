/*
	Print the list of prime numbers from 2  to 100 (one after another in main function)
	The process of checking if the given number is prime or not has to be executed concurrently
	DO NOT use the "primes []int" package variable
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	for i := 2; i <= 100; i++ {
		wg.Add(1)
		go process(i, wg, ch)
	}

	printWg := sync.WaitGroup{}
	printWg.Add(1)
	go func() {
		for primeNo := range ch {
			fmt.Println(primeNo)
		}
		fmt.Println("All prime numbers received")
		printWg.Done()
	}()

	wg.Wait()
	close(ch)
	printWg.Wait()

	/*
		done := make(chan struct{})
		go func() {
			for primeNo := range ch {
				fmt.Println(primeNo)
			}
			fmt.Println("All prime numbers received")
			close(done)
		}()

		wg.Wait()
		close(ch)
		<-done
	*/

}

func process(no int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	if isPrime(no) {
		ch <- no
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
