// もし初期化子が指定されている場合、型を省略できます
package main

import "fmt"

var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
  fmt.Println(i, j, c, python, java);
}
