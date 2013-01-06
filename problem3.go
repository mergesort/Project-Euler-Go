package main

import "fmt"

func isPrime(num int64) bool {
	var i int64 = 3
	for ; i < num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primeFactor(max int64) int64 {

	primeFactors := []int64{}
	var i int64 = 1

	for ; i < max; i += 2 {
		if max%i == 0 {
			if isPrime(i) {
				primeFactors = append(primeFactors, i)
				fmt.Println(i, primeFactors)
			}
		}
		fmt.Println(i)
	}

	return primeFactors[len(primeFactors)]
}

func main() {
	fmt.Println(primeFactor(600851475143))
}
