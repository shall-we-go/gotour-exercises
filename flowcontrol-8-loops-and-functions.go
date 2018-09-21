/*
https://tour.golang.org/flowcontrol/8
*/
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x == 0 || x == 1 {
		return x
	}
	start := 0.0
	end := x
	mid := (start + end) / 2
	precision := 1e-10
	for math.Abs(mid*mid-x) > precision {
		if mid*mid > x {
			end = mid
		} else {
			start = mid
		}
		mid = (start + end) / 2
	}
	return mid
}

func main() {
	fmt.Println(Sqrt(2))
}
