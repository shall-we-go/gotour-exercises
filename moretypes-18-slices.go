/*
https://tour.golang.org/moretypes/18
*/
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		picRow := make([]uint8, dy)
		pic[x] = picRow
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8(x ^ y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
