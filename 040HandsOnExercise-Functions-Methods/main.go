package main

import "fmt"

type person struct {
	first string
}

// Method with receiver
func (p person) speak() {
	fmt.Println("I am ", p.first)
}

// Function
func speakagain(p person) {
	fmt.Println("I speak again ", p.first)
}

func main() {
	p1 := person{
		first: "Jitendra",
	}
	p2 := person{
		first: "Jitu",
	}

	p1.speak()
	p2.speak()

	speakagain(p1)
	speakagain(p2)
}
