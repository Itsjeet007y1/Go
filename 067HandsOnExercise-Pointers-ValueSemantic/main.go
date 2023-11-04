package main

import "fmt"

// value semantic
func addOne(v int) int {
	return v + 1
}

// pointer semantic
func addOneP(v *int) {
	*v += 1
}

func main() {

	// value semantic
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(a)

	// pointer semantic
	b := 1
	fmt.Println("pointer semantic")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

}
