package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v of type %T\n", &x, &x)

	s := "Jitendra"
	fmt.Printf("%v %T\n", &s, &s)
}
