package main

import (
	"fmt"
	"math"
)

func Sqrt(z float64) float64 {
	v := float64(0)
	fmt.Println(math.Sqrt(z))
	fmt.Println("-----------------------")
	for i := float64(0); i < 10; i++ {
		v = z - ((z*z - i) / (2 * z))
		fmt.Println(v)
	}
	fmt.Println("-----------------------")
	return v
}

func main() {
	fmt.Println(Sqrt(2))
}
