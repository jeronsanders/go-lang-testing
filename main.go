package main

import (
	"fmt"
	"runtime"
)

func foo() {
	defer fmt.Println("This line is deferred.")
	fmt.Println("This is from foo.")

}

func main() {
	fmt.Println(runtime.NumCPU())
	foo()
}
