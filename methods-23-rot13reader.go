/*
https://tour.golang.org/methods/23
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (this rot13Reader) Read(b []byte) (int, error) {
	b = b[0:0]
	for {
		temp := make([]byte, 1)
		_, err := this.r.Read(temp)
		if err == io.EOF {
			break
		}
		A := byte('A')
		if (temp[0] >= A && temp[0] <= A+13) || (temp[0] >= A+26 && temp[0] <= A+26+13) {
			temp[0] += 13
		} else if (temp[0] > A+13 && temp[0] <= A+26) || (temp[0] > A+26+13 && temp[0] <= A+26+13+26) {
			temp[0] -= 13
		}
		b = append(b, temp[0])
	}
	return len(b), io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
