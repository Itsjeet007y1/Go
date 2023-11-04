package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6} // unfurling a slice
	x := sum(xi...)               // unfurling a slic

	// x := sum() // will return zero
	fmt.Println("The sum is", x)

	greeting("nobody", "asdf", "asdfsd", "d")
}

func sum(i ...int) int {
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	n := 0
	for _, v := range i {
		n += v
	}
	return n
}

func greeting(prefix string, who ...string) {
	fmt.Println(prefix, who)
}
