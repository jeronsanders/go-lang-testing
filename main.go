package main

import (
	"fmt"
	"runtime"
)

func foo() {
	fmt.Println("This is from foo.")
}

func main() {
	fmt.Println(runtime.NumCPU())
	foo()
}
