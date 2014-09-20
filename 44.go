package main

import "fmt"

func fibonacci() func() int {
	sum := 0
	x := 0
	return func() int {
		if x < 2 {
			x++
			return sum
		} else {
			x++
			sum += 2
			return sum
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
