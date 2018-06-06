package main

import (
	"fmt"

	"github.com/vigo5190/goimports-example/a"
	foo "github.com/vigo5190/goimports-example/b"
	bar "github.com/vigo5190/goimports-example/a"
)

func main() {
	fmt.Printf("%#v\n", a.Foo)
	fmt.Printf("%#v\n", foo.Foo)
	fmt.Printf("%#v\n", bar.Foo)
}
