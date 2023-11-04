package main

import "fmt"

// a named function with identifier
// func (r recieve) identifier(p parameter(s)) (r return(s)) { code }

// a anonymous function
// func identifier(p parameter(s)) (r return(s)) { code }

// Function expression example

func main() {
	foo()

	x := func() {
		fmt.Println("Annonumous func")
	}

	y := func(s string) {
		fmt.Println("This is anonymous function showing my name.", s)
	}

	x()
	y("Jeet")
}

func foo() {
	fmt.Println("Foo ran")
}
