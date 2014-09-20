package main

import "fmt"

func Cbrt(x complex128) complex128 {
	return x - ((x*x*x - 7.444404) / 3 * x * x)
}

func main() {
	fmt.Println(Cbrt(2))
}
