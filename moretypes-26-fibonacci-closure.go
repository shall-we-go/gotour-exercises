/*
https://tour.golang.org/moretypes/26
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	fibonacciFunc := func(num int) int {
		a := 1
		b := 0
		temp := 0
		for num >= 0 {
			temp = a
			a = a + b
			b = temp
			num--
		}
		return b
	}
	return fibonacciFunc
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
