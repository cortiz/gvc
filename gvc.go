package main

import (
	"fmt"

	"tinygo.org/x/go-llvm"
)

func main() {
	fmt.Printf("GVC 0.0.1 LLVM %s\n", llvm.Version)
}
