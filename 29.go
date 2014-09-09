// new(T) という表現は、ゼロ初期化した( zeroed ) T の値をメモリに割り当て、そのポインタを返します。
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}
