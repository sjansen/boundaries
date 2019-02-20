package main

import (
	"fmt"

	"example.com/foo/pkg/a"
	"example.com/foo/pkg/b"
	"example.com/foo/pkg/c"
)

func main() {
	fmt.Println(a.Value())
	fmt.Println(b.Value())
	fmt.Println(c.Value())
}
