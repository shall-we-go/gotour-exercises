/*
https://tour.golang.org/methods/20
*/
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (this ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(this))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	if x == 0 || x == 1 {
		return x, nil
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
	return mid, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
