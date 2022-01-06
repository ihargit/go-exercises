package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:= 0;i < 10;i++ {
		squared := z*z
		fmt.Printf("X: %g Squared: %g Z: %g: math.Sqrt: %g\n", x, squared, z, math.Sqrt(x))
		if math.Abs(squared - x) < 0.00000000000001  {
			return z; 
		}
		z -= (squared - x) / (2*z)
	}
	return 0;
}

func main() {
	fmt.Println(Sqrt(6))
}
