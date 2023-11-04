package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) Speak() {
	fmt.Println("My name is: ", p.first, " and age is ", p.age)
}

func main() {
	p1 := person{
		first: "Jitendra",
		age:   30,
	}
	p1.Speak()
}
