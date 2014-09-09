// structリテラルは、フィールドの値を列挙することによって、構造体の初期値の割り当てを示しています
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}
	q = &Vertex{1, 2}
	r = Vertex{X: 1}
	s = Vertex{}
)

func main() {
	fmt.Println(p, q, r, s)
}
