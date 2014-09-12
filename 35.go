// インデックスや値を " _ "（アンダーバー） へ代入することで破棄することができます。
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	for key, _ := range pow {
		fmt.Printf("%d\n", key)
	}
}
