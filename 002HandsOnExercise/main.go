package main

import "fmt"

func main() {
	s, i, f := "Happiness", 43, 43.2
	fmt.Printf("%v of type %T\n", s, s)
	fmt.Printf("%v of type %T\n", i, i)
	fmt.Printf("%v of type %T\n", f, f)
}
