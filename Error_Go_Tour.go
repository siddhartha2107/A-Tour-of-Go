package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	f float64
}
func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: ")
}

func run(x float64) (error) {
	return &ErrNegativeSqrt{x}
}

func Sqrt(x float64) float64 {
	if x<float64(0) {
			if err := run(x); err != nil {
				fmt.Print(err)
				return x
		}
	}
	z:=1.0
	for i:=0;i<10; 
	{	
		z -= (z*z - x) / (2*z)
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(-2))
}
