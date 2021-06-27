package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return factorial(n-1) * n
}

func power(n int, m int) int {
	if m == 0 {
		return 1
	} else {
		return (n * power(n, m-1))
	}
}

func main() {
	fmt.Print(factorial(5))
	fmt.Print(power(2, 3))
}
