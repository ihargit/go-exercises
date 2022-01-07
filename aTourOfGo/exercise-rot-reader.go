package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	for {
		n, err := r.r.Read(b)
		for i := 0; i < n; i++ {
			v := b[i]
			if v >= 65 && v <= 90 {
				if v - 13 < 65 {
					b[i] = 90 - (65 - (v - 12))
				} else {
					b[i] = v - 13
				}
			}
			if v >= 97 && v <= 122 {
				if v - 13 < 97 {
					b[i] = 122 - (97 - (v - 12))
				} else {
					b[i] = v - 13
				}
			}
		}
		return n, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}