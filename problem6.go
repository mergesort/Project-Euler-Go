package main

import "fmt"

func sumSquare(max int) int {
	squares := 0
	summed := 0
	for i := 1; i <= max; i++ {
		summed += i * i
		squares += i
	}

	return (squares * squares) - summed
}

func main() {
	fmt.Println(sumSquare(100))
}
