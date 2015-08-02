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

	//receive primes
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func fetchPrimes(c chan<- int, limit int) {
	for i := 0; i < limit; i++ {
		if isPrime(i) {
			c <- i
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
