package main

import "fmt"

func main() {
	x := doMath(33, 2, addition)
	fmt.Println(x)

	y := doMath(33, 2, sustract)
	fmt.Println(y)

	fmt.Printf("%T\n", addition)
	fmt.Printf("%T\n", sustract)
	fmt.Printf("%T\n", doMath)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func addition(a int, b int) int {
	return a + b
}

func sustract(a int, b int) int {
	return a - b
}
