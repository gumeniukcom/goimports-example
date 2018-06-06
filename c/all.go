package c

import "fmt"

var Foo string

func init() {
	Foo = "bar"
	fmt.Printf("%#v\n", Foo)
}
