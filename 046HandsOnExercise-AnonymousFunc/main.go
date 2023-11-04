package main

import "fmt"

// a named function with identifier
// func (r recieve) identifier(p parameter(s)) (r return(s)) { code }

// a anonymous function
// func identifier(p parameter(s)) (r return(s)) { code }

func main() {
	foo()

	func() {
		fmt.Println("Annonumous func")
	}()

	func(s string) {
		fmt.Println("This is anonymous function showing my name.")
	}("Jeet")

}

func foo() {
	fmt.Println("Foo ran")
}
