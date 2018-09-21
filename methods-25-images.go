/*
https://tour.golang.org/methods/25
*/
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x int
	y int
	w int
	h int
}

func (this Image) Bounds() image.Rectangle {
	return image.Rect(this.x, this.y, this.w, this.h)
}

func (this Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (this Image) At(x, y int) color.Color {
	return color.RGBA{0, 255, 0, 255}
}

func main() {
	m := Image{0, 0, 200, 100}
	pic.ShowImage(m)
}
