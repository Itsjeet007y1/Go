package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 22
	fmt.Println(z)

	a, b := 43, "Happiness"
	fmt.Println(a, b)

	var c float32 = 43.2
	fmt.Printf("%v is of type %T\n", c, c)

	e, f, _ := 34, 56, 44
	fmt.Println(e, f)

}

/*
var fo zero value declaration
short declaration operator
multiple initialization
var when specifity is required
blank identifier
*/
