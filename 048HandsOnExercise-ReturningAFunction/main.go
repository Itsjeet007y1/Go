package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	println(y())
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
