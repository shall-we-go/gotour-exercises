/*
https://tour.golang.org/methods/22
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (this MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
