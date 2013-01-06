package main

import "fmt"

func fibonacci(max int) int {
	earliest := 1
	next := 1
	current := 1
	total := 0

	for next < max {
		current = earliest + next
		earliest = next
		next = current
		if next%2 == 0 {
			total += next
		}
	}

	return total
}

func main() {
	fmt.Println(fibonacci(4000000))
}
