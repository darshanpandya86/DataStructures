package main

import "fmt"

func sumOfNatural(n int) int {
	if n == 0 {
		return 0
	} else {
		fmt.Print(n)
		return sumOfNatural(n-1) + 1
	}
}

func main() {
	fmt.Print(sumOfNatural(5))
}
