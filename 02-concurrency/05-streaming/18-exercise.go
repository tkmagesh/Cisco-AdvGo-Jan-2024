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

var primes []int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 2; i <= 100; i++ {
		wg.Add(1)
		go process(i, wg)
	}
	wg.Wait()
	fmt.Println(primes)
}

func process(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	if isPrime(no) {
		mutex.Lock()
		{
			primes = append(primes, no)
		}
		mutex.Unlock()
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
