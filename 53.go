// インターフェース型( interface type )は、メソッド群で定義されます。
// Abs メソッドが、 Vertex ではなく *Vertex の定義であり、 Vertex が Abser を満たしていないということになるためエラーとなります
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64
