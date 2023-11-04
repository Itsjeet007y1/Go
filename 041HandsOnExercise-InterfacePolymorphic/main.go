package main

import "fmt"

type person struct {
	first string
}
type secretAgent struct {
	person
	ltk bool
}

// Method with receiver
func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent ", sa.first)
}

// Function
func speakagain(p person) {
	fmt.Println("I speak again ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Jitendra",
	}
	p2 := person{
		first: "Jitu",
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p1.speak()
	p2.speak()

	sa1.speak()
	saySomething(sa1)
	saySomething(p1)

	speakagain(p1)
	speakagain(p2)
}
