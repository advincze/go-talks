package main

import (
	"fmt"

	"math"
)

//BEGIN OMIT
func main() {
	//channels
	c := make(chan int, 2)
	//goroutines
	go fetchPrimes(c, 10000)
	fmt.Println("fetch started")
	//receive primes
	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // HL
	}
}

func fetchPrimes(c chan<- int, limit int) {
	for i := 2; i < limit; i++ {
		if isPrime(i) {
			fmt.Println("sending", i)
			c <- i // HL
			fmt.Println("sent", i)
		}
	}
}

//END OMIT

func isPrime(n int) bool {
	if n == 1 || n == 2 {
		return true
	}

	if math.Mod(float64(n), 2) == 0 {
		return false
	}

	for i := 3.0; i <= math.Floor(math.Sqrt(float64(n))); i += 2.0 {
		if math.Mod(float64(n), i) == 0 {
			return false
		}
	}

	return true
}
