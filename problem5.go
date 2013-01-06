package main

import "fmt"

func divisible120() int {
	for i := 2520; ; i += 2520 {
		for j := 11; j <= 20; j++ {
			if i%j == 0 && j < 20 {
				continue
			} else if i%j == 0 && j == 20 {
				return i
				break
			}
			break
		}
	}

	return 0
}

func main() {
	fmt.Println(divisible120())
}
