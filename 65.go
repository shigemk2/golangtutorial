// チャネルは、 バッファ できます。 バッファchannel の初期化のために、 make の２つ目の引数にバッファの長さを与えます
package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
