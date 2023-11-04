package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var c, python, java bool
	const j int = 44
	fmt.Println(add(10, 30))
	fmt.Println(c, python, java, j)
	fmt.Printf("%T \t %T \t %T \t %T \t", c, python, java, j)
	sayHello()
}

func sayHello() {
	// fmt.Println("hello world", math.Pi)
	// a, b := swap("Hello", "World !")
	// fmt.Println(a, b)
}
