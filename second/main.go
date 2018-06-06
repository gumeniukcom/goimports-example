package main

import (
	"fmt"

	"github.com/vigo5190/goimports-example/a"
	_ "github.com/vigo5190/goimports-example/c"
)

func main() {
	fmt.Printf("%#v\n", a.Foo)
}
