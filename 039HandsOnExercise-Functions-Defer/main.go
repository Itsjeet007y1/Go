package main

import "fmt"

func main() {
	defer foo() // this function will wait untill the completion of the other statments of main function
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
