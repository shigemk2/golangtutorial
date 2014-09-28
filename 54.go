//  ある型に、メソッドを実装することで、インターフェースを実装します。
// インターフェースを実装することを明示的に宣言する必要はありません。
// この暗黙的なインターフェースは、インターフェースを定義するパッケージと、実装されているパッケージとを分離することができます。
package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout.implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
