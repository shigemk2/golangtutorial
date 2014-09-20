package main

import "fmt"

func fibonacci(x int) int {
	if x < 2 {
		return 1
	} else {
		return fibonacci(x-2) + fibonacci(x-1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
