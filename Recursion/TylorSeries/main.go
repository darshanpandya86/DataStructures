package main

import "fmt"

var ss float32 = 0.00

func TylorSeries(x int, n int) float32 {
	if n == 0 {
		return ss
	}
	ss = 1 + float32(x)*(ss/float32(n))
	return TylorSeries(x, n-1)

}

func main() {
	fmt.Print(TylorSeries(1, 10))
}
