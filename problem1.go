package main

import "fmt"

func calculateSum(x int) int {
	var sum int = 0

	for i := 0; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(calculateSum(10001))
}
