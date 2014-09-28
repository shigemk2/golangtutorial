// エラーは、エラー文字列としてエラーそのものを説明できるものです。 このしくみは、Goの組込みのインターフェース型 error に、文字列を返すひとつのメソッド Error をあらかじめ定義したことで実現されています
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
