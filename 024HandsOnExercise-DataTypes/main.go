package main

import "fmt"

func main() {
	// create a slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Printf("%v of type %T at index %v\n", v, v, i)
	}

	fmt.Printf("%v of type %T\n", x, x)
	fmt.Printf("%#v of type %T", x, x)
}
