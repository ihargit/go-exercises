package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
)

func Pic(dx, dy int) [][]uint8 {
	b := make([][]uint8, dy)
	for i := range b {
		b[i] = make([]uint8, dx)
		for x := range b[i] {
			b[i][x] = uint8(rand.Intn(255))
		}
	}
	return b
}

func main() {
	Pic(2, 2)
	pic.Show(Pic)
}