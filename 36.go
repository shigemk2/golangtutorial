package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mat := make([][]uint8, dx)
	for i := range mat {
		mat[i] = make([]uint8, dy)
		for j := range mat[i] {
			mat[i][j] = uint8(i ^ j)
		}
	}
	return mat
}

func main() {
	pic.Show(Pic)
}
