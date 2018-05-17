package main

import (
	"fmt"
)

func main() {
	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for num := range numberList {
		if (num % 2) == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
