package main

import "fmt"

var x = 41

const z = 42

func main() {

	y := 43
	fmt.Printf("Value %v of having have type of %T\n", x, x)
	fmt.Printf("Value %v of having type of %T\n", y, y)
	fmt.Printf("Value %v of having type of %T\n", z, z)

}
