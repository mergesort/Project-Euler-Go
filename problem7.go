package main

import "fmt"

func isPrime(num int) bool {
	for i := 3; i < num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func numberedPrime(max int) int {
	counter := 1
	i := 3
	for {
		if isPrime(i) {
			counter++
		}
		if counter == max {
			return i
		}
		i++
	}

	panic("How did we get here??")
}

func main() {
	fmt.Println(numberedPrime(10001))
}
