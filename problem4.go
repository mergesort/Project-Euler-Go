package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkIntForPalimdrome(num int) bool {
	numStr := strconv.Itoa(num)
	numArray := strings.Split(numStr, "")
	numLength := len(numArray)

	for i := 0; i < numLength/2; i++ {
		if i >= numLength/2 {
			break
		} else if numArray[i] == numArray[numLength-i-1] {
			continue
		} else {
			return false
		}
	}

	return true
}

func palindrome(max int) int {
	highest := 0

	for i := max; i > 100; i-- {
		for j := max; j > 100; j-- {
			if checkIntForPalimdrome(i * j) {
				if i < j {
					break
				}

				if i*j > highest {
					highest = i * j
					break
				}
			}
		}
	}
	return highest
}

func main() {
	maxNum := 999
	fmt.Println(palindrome(maxNum))
}
