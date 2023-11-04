package main

import "fmt"

func main() {

	c := make(chan int, 2)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("--------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)

	// specific to general doesn't assign
	// c = cr
	// c = cs

	// general to specific can be assign
	cr = c
	cs = c
}
