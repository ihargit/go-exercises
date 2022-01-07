package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i:= 0;i < 10;i++ {
		squared := z*z
		fmt.Printf("X: %g Squared: %g Z: %g: math.Sqrt: %g\n", x, squared, z, math.Sqrt(x))
		if math.Abs(squared - x) < 0.00000000000001  {
			return z, nil; 
		}
		z -= (squared - x) / (2*z)
	}
	return 0, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g",
		e)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}