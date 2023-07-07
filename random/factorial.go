package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	result := factorial(number)
	fmt.Printf("The factorial of %d is %d\n", number, result)
}
